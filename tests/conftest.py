"""Pytest configuration and shared fixtures."""

import asyncio
import os
from typing import AsyncGenerator, Generator

import pytest
import pytest_asyncio
from httpx import AsyncClient
from sqlalchemy.ext.asyncio import AsyncSession, create_async_engine
from sqlalchemy.orm import sessionmaker

from src.core.config import Settings
from src.core.database import Base, get_db
from src.main import app


@pytest.fixture(scope="session")
def event_loop() -> Generator[asyncio.AbstractEventLoop, None, None]:
    """Create an instance of the default event loop for the test session."""
    loop = asyncio.get_event_loop_policy().new_event_loop()
    yield loop
    loop.close()


@pytest.fixture
def test_settings() -> Settings:
    """Create test settings."""
    return Settings(
        environment="test",
        database_url="sqlite+aiosqlite:///test.db",
        redis_url="redis://localhost:6379/1",
        log_level="DEBUG",
    )


@pytest_asyncio.fixture
async def test_db(test_settings: Settings) -> AsyncGenerator[AsyncSession, None]:
    """Create test database session."""
    engine = create_async_engine(
        test_settings.database_url,
        echo=True,
    )
    
    async with engine.begin() as conn:
        await conn.run_sync(Base.metadata.create_all)
    
    async_session = sessionmaker(
        engine, class_=AsyncSession, expire_on_commit=False
    )
    
    async with async_session() as session:
        yield session
        await session.rollback()
    
    async with engine.begin() as conn:
        await conn.run_sync(Base.metadata.drop_all)
    
    await engine.dispose()


@pytest_asyncio.fixture
async def client(test_db: AsyncSession) -> AsyncGenerator[AsyncClient, None]:
    """Create test HTTP client."""
    app.dependency_overrides[get_db] = lambda: test_db
    
    async with AsyncClient(app=app, base_url="http://test") as ac:
        yield ac
    
    app.dependency_overrides.clear()


@pytest.fixture
def mock_user_data() -> dict:
    """Mock user data for testing."""
    return {
        "id": 1,
        "email": "test@example.com",
        "username": "testuser",
        "is_active": True,
        "is_superuser": False,
    }


@pytest.fixture
def auth_headers(mock_user_data: dict) -> dict:
    """Create authentication headers for testing."""
    # In a real application, you would generate a proper JWT token here
    return {"Authorization": "Bearer test-token"}


class TestDatabase:
    """Test database utilities."""
    
    @staticmethod
    async def create_user(db: AsyncSession, **kwargs) -> dict:
        """Create a test user."""
        from src.models.user import User
        
        user_data = {
            "email": "test@example.com",
            "username": "testuser",
            "hashed_password": "hashed_password",
            **kwargs
        }
        user = User(**user_data)
        db.add(user)
        await db.commit()
        await db.refresh(user)
        return user