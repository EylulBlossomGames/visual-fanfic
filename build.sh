#!/bin/bash

# Exit immediately if a command exits with a non-zero status.
set -e

# Name of your application
APP_NAME="visual-fanfic"

# Output directory
OUTPUT_DIR="dist"

# Clean the output directory
rm -rf "$OUTPUT_DIR"
mkdir -p "$OUTPUT_DIR"

# Define the target platforms
PLATFORMS=(
    "linux/amd64"
    "darwin/amd64"
    "darwin/arm64"
    "windows/amd64"
)

# Iterate over platforms and compile
for platform in "${PLATFORMS[@]}"
do
    # Separate GOOS and GOARCH
    GOOS=${platform%/*}
    GOARCH=${platform#*/}

    PLATFORM_FOLDER="$OUTPUT_DIR/$GOOS-$GOARCH"

    mkdir "$PLATFORM_FOLDER"

    cp -R src/ "$PLATFORM_FOLDER"
    cp -R config/ "$PLATFORM_FOLDER"

    # Define the output file name
    OUTPUT_NAME="$OUTPUT_DIR/$GOOS-$GOARCH/$APP_NAME"

    # Add the .exe extension for Windows
    if [ "$GOOS" == "windows" ]; then
        OUTPUT_NAME+=".exe"
    fi

    echo "Building for $GOOS/$GOARCH..."

    # The actual build command
    CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -o "$OUTPUT_NAME" .
done

echo "Done!"
