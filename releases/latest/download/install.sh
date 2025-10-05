#!/bin/bash
set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

OS="$(uname -s)"
ARCH="$(uname -m)"

echo -e "${BLUE} Installing Run - Universal Script Runner${NC}"
echo "=========================================="

# Detect OS
case "$OS" in
  Linux*)   OS=linux;;
  Darwin*)  OS=darwin;;
  *)
    echo -e "${RED}✗ Unsupported OS: $OS${NC}"
    echo "Please download manually from: https://github.com/Khaliiloo/run/releases"
    exit 1
    ;;
esac

# Detect Architecture
case "$ARCH" in
  x86_64)   ARCH=amd64;;
  arm64)    ARCH=arm64;;
  aarch64)  ARCH=arm64;;
  *)
    echo -e "${RED}✗ Unsupported architecture: $ARCH${NC}"
    echo "Please download manually from: https://github.com/Khaliiloo/run/releases"
    exit 1
    ;;
esac

BINARY="run-${OS}-${ARCH}"
URL="https://github.com/Khaliiloo/run/releases/latest/download/${BINARY}"
INSTALL_DIR="/usr/local/bin"

echo -e "${YELLOW}Detected: ${OS}-${ARCH}${NC}"
echo "Downloading from: $URL"
echo ""

# Download binary
if command -v curl &> /dev/null; then
    if ! curl -fL "$URL" -o /tmp/run 2>/dev/null; then
        echo -e "${RED}✗ Download failed. Please check your internet connection.${NC}"
        echo "Or download manually from: https://github.com/Khaliiloo/run/releases"
        exit 1
    fi
elif command -v wget &> /dev/null; then
    if ! wget -q "$URL" -O /tmp/run 2>/dev/null; then
        echo -e "${RED}✗ Download failed. Please check your internet connection.${NC}"
        echo "Or download manually from: https://github.com/Khaliiloo/run/releases"
        exit 1
    fi
else
    echo -e "${RED}✗ Neither curl nor wget found. Please install one of them.${NC}"
    exit 1
fi

# Check if download was successful
if [ ! -f /tmp/run ]; then
    echo -e "${RED}✗ Download failed${NC}"
    exit 1
fi

echo -e "${GREEN}✓ Downloaded successfully${NC}"

# Make executable
chmod +x /tmp/run

# Try to move to /usr/local/bin
echo "Installing to $INSTALL_DIR..."
if [ -w "$INSTALL_DIR" ]; then
    mv /tmp/run "$INSTALL_DIR/run"
else
    echo -e "${YELLOW}Installing to $INSTALL_DIR requires sudo...${NC}"
    sudo mv /tmp/run "$INSTALL_DIR/run"
fi

# Verify installation
if command -v run &> /dev/null; then
    VERSION=$(run --version 2>/dev/null || echo "unknown")
    echo ""
    echo -e "${GREEN}✓ Run installed successfully!${NC}"
    echo -e "${GREEN}  Version: $VERSION${NC}"
    echo ""
    echo "Try it out:"
    echo -e "  ${BLUE}run --version${NC}        # Show version"
    echo -e "  ${BLUE}run --list${NC}           # List supported languages"
    echo -e "  ${BLUE}run yourscript.py${NC}    # Run a script"
    echo ""
    echo "Documentation: https://github.com/Khaliiloo/run"
else
    echo -e "${RED}✗ Installation failed. Run is not in PATH.${NC}"
    echo "You may need to add $INSTALL_DIR to your PATH."
    exit 1
fi