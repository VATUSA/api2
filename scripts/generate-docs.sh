#!/bin/bash

echo "Generating docs..."
echo "- Downloading swag cli"
curl -Ls https://github.com/swaggo/swag/releases/download/v1.7.6/swag_1.7.6_Linux_x86_64.tar.gz -o /tmp/swag.tar.gz

echo "- Extracting swag cli"
mkdir /tmp/swag
tar -zxf /tmp/swag.tar.gz -C /tmp/swag

echo "- Generating docs"
/tmp/swag/swag init --parseDependency --parseInternal --parseDepth 3

echo "- Cleaning up"
rm -rf /tmp/swag

echo "- Done"