#!/usr/bin/env bash

set -e

VERSION="$(<VERSION)"

git tag -s "$VERSION" -m "Release $VERSION"
git push origin "$VERSION"

gh release new "$VERSION" --verify-tag
