"""Health check API endpoints."""

from typing import Dict

from fastapi import APIRouter, Depends
from sqlalchemy.ext.asyncio import AsyncSession

from src.core.database import get_db
from src.core.logging import get_logger

router = APIRouter()
logger = get_logger(__name__)


@router.get("/", response_model=Dict[str, str])
async def health_check():
    """Basic health check."""
    return {"status": "healthy"}


@router.get("/detailed", response_model=Dict[str, str])
async def detailed_health_check(db: AsyncSession = Depends(get_db)):
    """Detailed health check with database connectivity."""
    try:
        # Test database connection
        await db.execute("SELECT 1")
        db_status = "healthy"
    except Exception as e:
        logger.error("Database health check failed", error=str(e))
        db_status = "unhealthy"
    
    return {
        "status": "healthy" if db_status == "healthy" else "degraded",
        "database": db_status,
    }