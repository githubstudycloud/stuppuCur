"""API route definitions."""

from fastapi import APIRouter

from src.api.v1.users import router as users_router
from src.api.v1.health import router as health_router

api_router = APIRouter()

# Include sub-routers
api_router.include_router(health_router, prefix="/health", tags=["health"])
api_router.include_router(users_router, prefix="/users", tags=["users"])