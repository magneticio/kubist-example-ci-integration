#!/usr/bin/env bash

set -e

VERSION=$1

echo "Tagging with version: ${VERSION}"

set -x

git tag -a "${VERSION}" -m "Release for version ${VERSION}"

git push origin ${VERSION}

echo "Pushed version: ${VERSION}"
