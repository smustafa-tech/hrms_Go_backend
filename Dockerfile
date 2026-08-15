# ╔══════════════════════════════════════════════════════════════╗
# ║                  STAGE 1 — builder                          ║
# ║  Purpose: Compile the Go source code into a binary          ║
# ║  This entire stage is THROWN AWAY after build               ║
# ║  Nothing from here reaches the final image except           ║
# ║  the compiled binary we explicitly copy                     ║
# ╚══════════════════════════════════════════════════════════════╝

# FROM tells Docker: "Start with this base image."
# golang:1.23-alpine is the official Go image based on Alpine Linux.
#
# Why Alpine? It is a minimal Linux distro (~5MB vs ~900MB for Ubuntu).
# We use it in the builder too so the toolchain is consistent.
#
# Why 1.23? Your go.mod says "go 1.26.4" but Go 1.26 is not released yet.
# This is likely a typo in go.mod. Go 1.23 is the latest stable version.
# The compiler is fully backward compatible — your code will compile fine.
#
# "AS builder" gives this stage a name so Stage 2 can reference it.
FROM golang:1.25-alpine AS builder

# RUN executes a shell command INSIDE the container during build.
# apk is Alpine's package manager (like apt on Ubuntu).
# We install:
#   git         → some Go packages need git to fetch dependencies
#   ca-certificates → needed for HTTPS calls during go mod download
# --no-cache means: don't store the package index, keeps image smaller.
RUN apk add --no-cache git ca-certificates

# WORKDIR sets the working directory inside the container.
# All following commands (COPY, RUN) happen relative to this path.
# If /app doesn't exist, Docker creates it automatically.
# Think of it as: cd /app
WORKDIR /app

# COPY <source-on-your-machine> <destination-inside-container>
#
# WHY copy go.mod and go.sum BEFORE the source code?
# This is the most important Docker optimization trick called LAYER CACHING.
#
# Docker builds images in layers. Each instruction = one layer.
# Docker caches each layer. If nothing changed, it reuses the cache.
#
# go.mod and go.sum only change when you add/remove dependencies.
# Your source code changes every time you write code.
#
# If we copied everything at once:
#   → Every code change = re-download ALL dependencies = slow (minutes)
#
# By copying go.mod/go.sum first and downloading deps separately:
#   → Code change = skip dep download, use cache = fast (seconds)
#   → Only a new dependency = re-download deps = expected
COPY go.mod go.sum ./

# Download all dependencies listed in go.mod.
# -v flag = verbose, shows what is being downloaded.
# Results are cached in /go/pkg/mod inside the container.
# This layer is cached as long as go.mod and go.sum don't change.
RUN go mod download

# NOW copy the rest of your source code.
# The dot "." means: copy everything from your project folder
# (respecting .dockerignore) into the current WORKDIR (/app).
# This happens AFTER dependency download intentionally (caching).
COPY . .

# Compile the Go application into a single binary called "main".
#
# CGO_ENABLED=0
#   CGO allows Go to call C code. We disable it because:
#   - Alpine uses musl libc, not glibc. CGO binaries built with glibc
#     will CRASH on Alpine.
#   - Disabling CGO produces a fully static binary with zero dependencies.
#   - The binary can run on ANY Linux, even a completely empty container.
#
# GOOS=linux
#   Cross-compile target OS = Linux.
#   Even if you build this on macOS or Windows, output is a Linux binary.
#
# go build
#   -o main         → name the output binary "main"
#   -ldflags="-w -s"
#       -w = strip DWARF debug info (not needed in production)
#       -s = strip symbol table
#       Together they reduce binary size by ~30%.
#   ./cmd/server/main.go → your entry point file
RUN CGO_ENABLED=0 GOOS=linux GOTOOLCHAIN=auto go build -o main -ldflags="-w -s" ./cmd/server/main.go


# ╔══════════════════════════════════════════════════════════════╗
# ║                  STAGE 2 — runner                           ║
# ║  Purpose: Run the compiled binary in a minimal environment  ║
# ║  This is the FINAL image that gets deployed                 ║
# ║  It has NO Go compiler, NO source code, NO dev tools        ║
# ╚══════════════════════════════════════════════════════════════╝

# Start completely fresh with Alpine Linux (~5MB).
# No Go compiler. No source code. Clean slate.
# This is what makes the final image tiny and secure.
FROM alpine:3.19

# Install runtime dependencies only.
#
# ca-certificates
#   SSL/TLS certificates. Required for your app to make HTTPS calls
#   (e.g. to external APIs). Without this, HTTPS requests fail with
#   "certificate signed by unknown authority".
#
# tzdata
#   Timezone data. Required if your app uses time.LoadLocation()
#   or formats timestamps with timezones. Good to always include.
RUN apk add --no-cache ca-certificates tzdata

# Security Best Practice: Never run your app as root inside a container.
#
# By default, Docker containers run as root (uid=0).
# If an attacker exploits your app, they get root inside the container.
# Running as a non-root user limits the blast radius.
#
# addgroup -S hrms   → create a system group named "hrms"
# adduser -S -G hrms hrms → create a system user "hrms" in group "hrms"
# -S means system account (no password, no home directory shell)
RUN addgroup -S hrms && adduser -S -G hrms hrms

# Set working directory in the runner stage.
WORKDIR /app

# COPY --from=builder
#   This is the magic of multi-stage builds.
#   "from=builder" means: reach back into Stage 1 (named "builder")
#   and grab the file at /app/main
#   Copy it into the current stage at /app/main
#
# We copy ONLY the compiled binary. Nothing else.
# No source code. No go.mod. No .env. Just the binary.
COPY --from=builder /app/main .

# Create the uploads directory inside the container.
# Your app may need to write uploaded files here at runtime.
# We create it now so the directory exists when the app starts.
RUN mkdir -p /app/uploads

# Give the "hrms" user ownership of the /app directory.
# Without this, the hrms user cannot read the binary or write to uploads/.
# chown -R = recursive change ownership
# hrms:hrms = user:group
RUN chown -R hrms:hrms /app

# Switch from root to the "hrms" user.
# All commands after this line run as "hrms", not root.
# The container process also runs as "hrms".
USER hrms

# EXPOSE documents which port the container listens on.
# This does NOT actually open the port — it is documentation.
# The actual port mapping happens in the docker run command with -p.
# Your app listens on PORT env var, defaulting to 5000 in main.go.
EXPOSE 5000

# HEALTHCHECK tells Docker how to test if your container is healthy.
# Docker will run this command every 30 seconds.
#   --interval=30s  → check every 30 seconds
#   --timeout=10s   → if check takes more than 10s, mark as failed
#   --start-period=15s → wait 15s after start before first check
#                        (gives your app time to connect to DB)
#   --retries=3     → mark unhealthy after 3 consecutive failures
# The login endpoint accepts POST only, so a GET health check against it always
# returns 404. BusyBox's netcat is available in Alpine and verifies that the
# application is actively listening without invoking an application endpoint.
HEALTHCHECK --interval=30s --timeout=10s --start-period=15s --retries=3 \
    CMD nc -z -w 5 127.0.0.1 5000 || exit 1

# CMD is the command that runs when the container starts.
# ["./main"] runs the compiled binary directly.
# We use JSON array format (exec form) — NOT shell form ("./main").
# Exec form: process runs directly, receives OS signals (SIGTERM for graceful shutdown)
# Shell form: process runs inside /bin/sh -c, signals may not reach your app.
CMD ["./main"]
