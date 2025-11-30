@echo off
echo ============================================
echo   Wardrobe Outfit Generator
echo   Starting Backend Server...
echo ============================================
echo.

cd /d "%~dp0"

echo Checking Go installation...
go version
if %errorlevel% neq 0 (
    echo ERROR: Go is not installed or not in PATH
    pause
    exit /b 1
)

echo.
echo Starting server on http://localhost:8080
echo.
echo Press Ctrl+C to stop the server
echo.

go run cmd/main.go

pause

