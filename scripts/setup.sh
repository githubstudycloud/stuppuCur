#!/bin/bash
# Project setup script

set -e

echo "🚀 Setting up Enterprise Python Template..."

# Check if Poetry is installed
if ! command -v poetry &> /dev/null; then
    echo "❌ Poetry is not installed. Please install Poetry first."
    echo "Visit: https://python-poetry.org/docs/#installation"
    exit 1
fi

# Install dependencies
echo "📦 Installing dependencies..."
poetry install

# Install pre-commit hooks
echo "🪝 Installing pre-commit hooks..."
poetry run pre-commit install

# Create .env file if it doesn't exist
if [ ! -f .env ]; then
    echo "📄 Creating .env file..."
    cat > .env << EOF
# Environment
ENVIRONMENT=development
DEBUG=true

# Database
DATABASE_URL=sqlite+aiosqlite:///./app.db
DATABASE_ECHO=false

# Redis
REDIS_URL=redis://localhost:6379/0

# Security
SECRET_KEY=your-secret-key-change-in-production
ALGORITHM=HS256
ACCESS_TOKEN_EXPIRE_MINUTES=30

# Logging
LOG_LEVEL=INFO
LOG_FORMAT=pretty

# Monitoring
ENABLE_METRICS=true
ENABLE_TRACING=true
EOF
    echo "✅ Created .env file with default settings"
else
    echo "📄 .env file already exists, skipping..."
fi

# Run database migrations
echo "🗄️ Setting up database..."
poetry run python -m src.cli migrate

echo "✅ Setup completed successfully!"
echo ""
echo "🎯 Next steps:"
echo "   1. Update .env file with your configuration"
echo "   2. Run 'poetry run python -m src.cli serve' to start the server"
echo "   3. Visit http://localhost:8000/docs for API documentation"
echo ""
echo "📚 Available commands:"
echo "   poetry run python -m src.cli --help"