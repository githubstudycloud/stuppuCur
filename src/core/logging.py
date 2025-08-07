"""Structured logging configuration."""

import logging
import sys
from typing import Any, Dict

import structlog
from rich.console import Console
from rich.logging import RichHandler

from src.core.config import get_settings

settings = get_settings()


def configure_logging() -> None:
    """Configure structured logging."""
    
    # Configure stdlib logging
    logging.basicConfig(
        format="%(message)s",
        stream=sys.stdout,
        level=getattr(logging, settings.log_level.upper()),
    )
    
    # Rich console for development
    if settings.is_development:
        console = Console()
        handler = RichHandler(
            console=console,
            rich_tracebacks=True,
            markup=True,
            show_time=True,
            show_path=True,
        )
        
        root_logger = logging.getLogger()
        root_logger.handlers = [handler]
    
    # Configure structlog
    processors = [
        structlog.contextvars.merge_contextvars,
        structlog.stdlib.filter_by_level,
        structlog.stdlib.add_logger_name,
        structlog.stdlib.add_log_level,
        structlog.stdlib.PositionalArgumentsFormatter(),
        structlog.processors.TimeStamper(fmt="iso"),
        structlog.processors.StackInfoRenderer(),
    ]
    
    if settings.log_format == "json":
        processors.append(structlog.processors.JSONRenderer())
    else:
        processors.extend([
            structlog.processors.CallsiteParameterAdder(
                parameters=[structlog.processors.CallsiteParameter.FUNC_NAME]
            ),
            structlog.dev.ConsoleRenderer(colors=settings.is_development),
        ])
    
    structlog.configure(
        processors=processors,
        wrapper_class=structlog.stdlib.BoundLogger,
        logger_factory=structlog.stdlib.LoggerFactory(),
        cache_logger_on_first_use=True,
    )


def get_logger(name: str = None) -> structlog.BoundLogger:
    """
    Get a structured logger.
    
    Args:
        name: Logger name
        
    Returns:
        BoundLogger: Configured logger instance
    """
    return structlog.get_logger(name)


def log_request_middleware(request_id: str) -> Dict[str, Any]:
    """
    Create logging context for request.
    
    Args:
        request_id: Unique request identifier
        
    Returns:
        Dict: Logging context
    """
    return {
        "request_id": request_id,
        "service": settings.app_name,
        "version": settings.app_version,
        "environment": settings.environment,
    }