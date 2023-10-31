#!/bin/sh

set -e

version=$1

if [ x"$version" = "x" ]; then
    echo "Usage: release.sh <version>"
    exit 1
fi

if ! [[ "$version" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "Invalid version ${version}"
    exit 2
fi

gh release create ${version} --draft --title="${version}" --generate-notes --target=main