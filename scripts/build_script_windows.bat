@echo off
setlocal

set "BINARY_NAME=bot"
set "PATH_MAIN=.\main"
set "BUILD_DIR=.\build"
set "PLATFORMS=windows/amd64,windows/386,windows/arm,linux/amd64,linux/386,linux/arm,linux/arm64,darwin/amd64,darwin/arm64"

if not exist %BUILD_DIR% (
    mkdir %BUILD_DIR%
)

if "%1" == "dev" (
    go build -o %BUILD_DIR%\%BINARY_NAME%.exe %PATH_MAIN%
) else (
    for %%P in (%PLATFORMS%) do (
        for /f "tokens=1,2 delims=/" %%A in ("%%P") do (
            set "GOOS=%%A"
            set "GOARCH=%%B"
            echo Building for GOOS=%%A and GOARCH=%%B...
            go build -o %BUILD_DIR%\%BINARY_NAME%-%%A-%%B.exe %PATH_MAIN%
        )
    )
)

echo Build Complete!
pause