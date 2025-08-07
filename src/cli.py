"""Command line interface for the application."""

import asyncio
import sys
from typing import Optional

import typer
import uvicorn
from rich.console import Console
from rich.table import Table

from src.core.config import get_settings
from src.core.database import create_tables, drop_tables
from src.core.logging import configure_logging, get_logger

configure_logging()
logger = get_logger(__name__)
console = Console()

app = typer.Typer(help="Enterprise Python Template CLI")
settings = get_settings()


@app.command()
def serve(
    host: str = typer.Option(settings.host, help="Host to bind to"),
    port: int = typer.Option(settings.port, help="Port to bind to"),
    reload: bool = typer.Option(False, help="Enable auto-reload"),
    workers: int = typer.Option(1, help="Number of workers"),
):
    """Start the web server."""
    console.print(f"🚀 Starting {settings.app_name} server...", style="green")
    console.print(f"📍 Server will be available at http://{host}:{port}")
    console.print(f"📚 API documentation at http://{host}:{port}/docs")
    
    uvicorn.run(
        "src.main:app",
        host=host,
        port=port,
        reload=reload,
        workers=workers if not reload else 1,
        log_config=None,
    )


@app.command()
def migrate():
    """Run database migrations (create tables)."""
    console.print("🔄 Running database migrations...", style="yellow")
    
    async def _migrate():
        try:
            await create_tables()
            console.print("✅ Database migrations completed successfully", style="green")
        except Exception as e:
            console.print(f"❌ Migration failed: {e}", style="red")
            sys.exit(1)
    
    asyncio.run(_migrate())


@app.command()
def reset_db():
    """Reset database (drop and recreate all tables)."""
    if not typer.confirm("Are you sure you want to reset the database? This will delete all data."):
        console.print("❌ Database reset cancelled", style="yellow")
        return
    
    console.print("🔄 Resetting database...", style="yellow")
    
    async def _reset():
        try:
            await drop_tables()
            await create_tables()
            console.print("✅ Database reset completed successfully", style="green")
        except Exception as e:
            console.print(f"❌ Database reset failed: {e}", style="red")
            sys.exit(1)
    
    asyncio.run(_reset())


@app.command()
def test(
    verbose: bool = typer.Option(False, "--verbose", "-v", help="Verbose output"),
    coverage: bool = typer.Option(False, "--coverage", "-c", help="Run with coverage"),
    fast: bool = typer.Option(False, "--fast", "-f", help="Skip slow tests"),
):
    """Run tests."""
    import subprocess
    
    console.print("🧪 Running tests...", style="yellow")
    
    cmd = ["poetry", "run", "pytest"]
    
    if verbose:
        cmd.append("-v")
    
    if coverage:
        cmd.extend(["--cov=src", "--cov-report=html", "--cov-report=term"])
    
    if fast:
        cmd.extend(["-m", "not slow"])
    
    try:
        result = subprocess.run(cmd, check=True)
        console.print("✅ Tests completed successfully", style="green")
    except subprocess.CalledProcessError as e:
        console.print(f"❌ Tests failed with exit code {e.returncode}", style="red")
        sys.exit(e.returncode)


@app.command()
def lint():
    """Run code linting."""
    import subprocess
    
    console.print("🔍 Running linting...", style="yellow")
    
    commands = [
        ["poetry", "run", "black", "--check", "src", "tests"],
        ["poetry", "run", "isort", "--check-only", "src", "tests"],
        ["poetry", "run", "flake8", "src", "tests"],
        ["poetry", "run", "mypy", "src"],
    ]
    
    failed = False
    for cmd in commands:
        try:
            subprocess.run(cmd, check=True)
        except subprocess.CalledProcessError:
            failed = True
    
    if failed:
        console.print("❌ Linting failed", style="red")
        sys.exit(1)
    else:
        console.print("✅ Linting completed successfully", style="green")


@app.command()
def format_code():
    """Format code using black and isort."""
    import subprocess
    
    console.print("🎨 Formatting code...", style="yellow")
    
    commands = [
        ["poetry", "run", "black", "src", "tests"],
        ["poetry", "run", "isort", "src", "tests"],
    ]
    
    for cmd in commands:
        subprocess.run(cmd, check=True)
    
    console.print("✅ Code formatting completed", style="green")


@app.command()
def info():
    """Show application information."""
    table = Table(title=f"{settings.app_name} Information")
    table.add_column("Setting", style="cyan")
    table.add_column("Value", style="green")
    
    table.add_row("App Name", settings.app_name)
    table.add_row("Version", settings.app_version)
    table.add_row("Environment", settings.environment)
    table.add_row("Debug", str(settings.debug))
    table.add_row("Database URL", settings.database_url)
    table.add_row("Redis URL", settings.redis_url)
    table.add_row("Log Level", settings.log_level)
    
    console.print(table)


@app.command()
def create_user(
    email: str = typer.Option(..., help="User email"),
    username: str = typer.Option(..., help="Username"),
    password: str = typer.Option(..., help="Password"),
    first_name: Optional[str] = typer.Option(None, help="First name"),
    last_name: Optional[str] = typer.Option(None, help="Last name"),
    superuser: bool = typer.Option(False, help="Create as superuser"),
):
    """Create a new user."""
    from src.core.database import get_db
    from src.schemas.user import UserCreate
    from src.services.user_service import UserService
    
    console.print(f"👤 Creating user: {username} ({email})", style="yellow")
    
    async def _create_user():
        try:
            async for db in get_db():
                user_service = UserService(db)
                user_data = UserCreate(
                    email=email,
                    username=username,
                    password=password,
                    first_name=first_name,
                    last_name=last_name,
                )
                
                user = await user_service.create_user(user_data)
                
                if superuser:
                    user.is_superuser = True
                    await db.commit()
                
                console.print(f"✅ User created successfully with ID: {user.id}", style="green")
                break
        except Exception as e:
            console.print(f"❌ User creation failed: {e}", style="red")
            sys.exit(1)
    
    asyncio.run(_create_user())


if __name__ == "__main__":
    app()