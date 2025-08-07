"""Example test file to demonstrate testing patterns."""

import pytest
from httpx import AsyncClient


class TestHealthCheck:
    """Test health check endpoints."""
    
    async def test_health_check(self, client: AsyncClient):
        """Test health check endpoint."""
        response = await client.get("/health")
        assert response.status_code == 200
        assert response.json() == {"status": "healthy"}


class TestUserAPI:
    """Test user API endpoints."""
    
    async def test_create_user(self, client: AsyncClient):
        """Test user creation."""
        user_data = {
            "email": "newuser@example.com",
            "username": "newuser",
            "password": "testpassword123"
        }
        response = await client.post("/api/v1/users/", json=user_data)
        assert response.status_code == 201
        data = response.json()
        assert data["email"] == user_data["email"]
        assert data["username"] == user_data["username"]
        assert "id" in data
    
    @pytest.mark.slow
    async def test_get_users(self, client: AsyncClient):
        """Test getting users list."""
        response = await client.get("/api/v1/users/")
        assert response.status_code == 200
        data = response.json()
        assert isinstance(data, list)


@pytest.mark.integration
class TestDatabaseIntegration:
    """Test database integration."""
    
    async def test_database_connection(self, test_db):
        """Test database connection."""
        result = await test_db.execute("SELECT 1")
        assert result.scalar() == 1


@pytest.mark.unit
class TestUtilities:
    """Test utility functions."""
    
    def test_example_utility(self):
        """Test example utility function."""
        # Example unit test
        assert True