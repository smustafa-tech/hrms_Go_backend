# HRMS Backend — Docker Commands Cheatsheet
# ==========================================
# यह file रोज़ काम आएगी। ध्यान से पढ़ो।


# ─────────────────────────────────────────
# 1. DAILY START (Laptop ON करने के बाद)
# ─────────────────────────────────────────

# सही order में चलाओ — पहले DB, फिर Backend, फिर Frontend
docker start postgres
docker start hrms-backend
docker start hrms-frontend

# तीनों एक साथ
docker start postgres hrms-backend hrms-frontend


# ─────────────────────────────────────────
# 2. DAILY STOP (Laptop बंद करने से पहले)
# ─────────────────────────────────────────

# तीनों एक साथ बंद करो
docker stop hrms-frontend hrms-backend postgres


# ─────────────────────────────────────────
# 3. STATUS CHECK (चल रहा है या नहीं)
# ─────────────────────────────────────────

# Running containers देखो
docker ps

# सभी containers देखो (stopped भी)
docker ps -a


# ─────────────────────────────────────────
# 4. LOGS (Error आए तो यहाँ देखो)
# ─────────────────────────────────────────

# Backend के logs
docker logs hrms-backend

# Postgres के logs
docker logs postgres

# Live logs देखो (real-time)
docker logs -f hrms-backend
docker logs -f postgres

# सिर्फ last 50 lines देखो
docker logs --tail 50 hrms-backend


# ─────────────────────────────────────────
# 5. CONTAINER के अंदर जाओ (Debugging)
# ─────────────────────────────────────────

# Backend container के अंदर
docker exec -it hrms-backend sh

# Postgres container के अंदर
docker exec -it postgres sh

# Postgres database में directly जाओ
docker exec -it postgres psql -U postgres -d hrms_db

# Container से बाहर आओ
# exit


# ─────────────────────────────────────────
# 6. FRESH START (सब हटाओ और दोबारा बनाओ)
# ─────────────────────────────────────────

# Containers हटाओ
docker rm -f hrms-backend postgres

# Image दोबारा बनाओ
docker build -t hrms-backend:v1 .

# Postgres चलाओ
docker run -d \
  --name postgres \
  --network hrms-network \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres123 \
  -e POSTGRES_DB=hrms_db \
  -v hrms-postgres-data:/var/lib/postgresql/data \
  -p 5433:5432 \
  postgres:16-alpine

# Backend चलाओ
docker run -d \
  --name hrms-backend \
  --network hrms-network \
  -p 5000:5000 \
  -e APP_NAME=HRMS \
  -e APP_ENV=production \
  -e PORT=5000 \
  -e DB_HOST=postgres \
  -e DB_PORT=5432 \
  -e DB_USER=postgres \
  -e DB_PASSWORD=postgres123 \
  -e DB_NAME=hrms_db \
  -e JWT_SECRET=mysecretkey \
  hrms-backend:v1


# ─────────────────────────────────────────
# 7. CODE CHANGE के बाद (Rebuild)
# ─────────────────────────────────────────

# Step 1 - पुराना container हटाओ
docker rm -f hrms-backend

# Step 2 - नई image बनाओ
docker build -t hrms-backend:v1 .

# Step 3 - नया container चलाओ
docker run -d \
  --name hrms-backend \
  --network hrms-network \
  -p 5000:5000 \
  -e APP_NAME=HRMS \
  -e APP_ENV=production \
  -e PORT=5000 \
  -e DB_HOST=postgres \
  -e DB_PORT=5432 \
  -e DB_USER=postgres \
  -e DB_PASSWORD=postgres123 \
  -e DB_NAME=hrms_db \
  -e JWT_SECRET=mysecretkey \
  hrms-backend:v1


# ─────────────────────────────────────────
# 8. API TEST COMMANDS (curl)
# ─────────────────────────────────────────

# Register
curl -s -X POST http://localhost:5000/api/auth/admin/register \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "Mustafa",
    "lastName": "Ahmed",
    "email": "mustafa@test.com",
    "password": "password123",
    "organizationName": "Test Company",
    "phone": "9999999999",
    "slug": "test-company"
  }' | python3 -m json.tool

# Login
curl -s -X POST http://localhost:5000/api/auth/user/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "mustafa@test.com",
    "password": "password123"
  }' | python3 -m json.tool


# ─────────────────────────────────────────
# 9. IMAGES (बनी हुई images देखो)
# ─────────────────────────────────────────

# सभी images देखो
docker images

# एक image हटाओ
docker rmi hrms-backend:v1

# Unused images हटाओ (space खाली करो)
docker image prune


# ─────────────────────────────────────────
# 10. NETWORK & VOLUME
# ─────────────────────────────────────────

# Networks देखो
docker network ls

# Network बनाओ (सिर्फ एक बार)
docker network create hrms-network

# Volumes देखो (database data)
docker volume ls

# ─────────────────────────────────────────
# IMPORTANT NOTES
# ─────────────────────────────────────────
#
# DB_HOST=postgres   → container का नाम है, localhost नहीं
# Port 5432          → container के अंदर postgres का port
# Port 5433          → तुम्हारे laptop से access करने का port
# Port 5000          → backend का port (laptop और container दोनों)
# hrms-network       → दोनों containers इसी से बात करते हैं
# hrms-postgres-data → database का data यहाँ save होता है
#                      container हटाने पर भी data safe रहता है


# ─────────────────────────────────────────
# 11. DOCKER COMPOSE (Recommended तरीका)
# ─────────────────────────────────────────
# docker-compose.yml यहाँ है:
# /home/mustafa/Desktop/HRMS_Go_Poject/docker-compose.yml
#
# इस folder में जाओ पहले:
# cd /home/mustafa/Desktop/HRMS_Go_Poject

# तीनों containers एक साथ start करो
docker compose up -d

# तीनों containers एक साथ बंद करो
docker compose down

# Status देखो
docker compose ps

# सबके logs एक साथ देखो (live)
docker compose logs -f

# सिर्फ backend के logs
docker compose logs -f hrms-backend

# सिर्फ frontend के logs
docker compose logs -f hrms-frontend

# सिर्फ postgres के logs
docker compose logs -f postgres

# Code change के बाद — rebuild करके restart करो
docker compose up -d --build

# सिर्फ backend rebuild करो
docker compose up -d --build hrms-backend

# सिर्फ frontend rebuild करो
docker compose up -d --build hrms-frontend

# सब बंद करो + containers हटाओ (data safe रहेगा)
docker compose down

# सब बंद करो + containers + volumes हटाओ (data भी जाएगा!)
# WARNING: यह database का data भी delete करेगा
docker compose down -v

# ─────────────────────────────────────────
# URLS — Browser में खोलो
# ─────────────────────────────────────────
# Frontend  → http://localhost:3000
# Backend   → http://localhost:5000
# Database  → localhost:5433 (pgAdmin या DBeaver से)
