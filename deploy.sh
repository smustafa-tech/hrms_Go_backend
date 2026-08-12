#!/bin/bash
# Auto deploy script for Railway

set -e

echo "🚀 Starting deployment..."

# Install Railway CLI if not present
if ! command -v railway &> /dev/null; then
    echo "📥 Installing Railway CLI..."
    npm install -g @railway/cli
fi

# Build the application
echo "📦 Building application..."
go build -o server ./cmd/server

# Run tests if any
if [ -f "go.test" ] || [ -d "test" ]; then
    echo "🧪 Running tests..."
    go test ./...
fi

# Login to Railway if not already
if ! railway whoami &> /dev/null; then
    echo "🔐 Please login to Railway..."
    railway login
fi

# Link project if not linked
if [ ! -f ".railway" ]; then
    echo "🔗 Linking Railway project..."
    railway link
fi

# Deploy to Railway
echo "☁️ Deploying to Railway..."
railway up

echo "✅ Deployment complete!"