#!/bin/bash
set -e

IMAGE_NAME=
DOCKER_PLATFORM=

docker build --platform  -f benzhi.Dockerfile -t  .

echo ""
echo "✅ Docker image '' built successfully!"
echo ""
echo "📋 Next steps (for testing):"
echo " • Interactive shell：docker run -it "
echo ""
