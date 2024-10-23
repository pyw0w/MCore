#!/bin/bash

BINARY_NAME="bot"
PATH_MAIN="./main"
BUILD_DIR="./build"
PLATFORMS=("windows/amd64" "windows/386" "windows/arm" "linux/amd64" "linux/386" "linux/arm" "linux/arm64" "darwin/amd64" "darwin/arm64")

# Create build directory if it doesn't exist
if [ ! -d "$BUILD_DIR" ]; then
  mkdir -p "$BUILD_DIR"
fi

# Loop over each platform and build
for PLATFORM in "${PLATFORMS[@]}"; do
    IFS="/" read -r -a SPLIT <<< "$PLATFORM"
    GOOS="${SPLIT[0]}"
    GOARCH="${SPLIT[1]}"
    OUTPUT="$BUILD_DIR/${BINARY_NAME}-${GOOS}-${GOARCH}"
    if [ "$GOOS" == "windows" ]; then
        OUTPUT+=".exe"
    fi
    echo "Building for GOOS=$GOOS and GOARCH=$GOARCH..."
    env GOOS="$GOOS" GOARCH="$GOARCH" go build -o "$OUTPUT" "$PATH_MAIN"
done

echo "Build Complete!"