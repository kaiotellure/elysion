#!/bin/bash

if [ ! -e ".github/version" ]; then
    echo ".github/version not found, run tidy.sh before releasing."
    exit 1
fi

version=$(cat .github/version)
echo "Releasing: v$version"

mkdir release 2>/dev/null
go build -ldflags="-X github.com/ikaio/tailmplx/help.Version=$version" -o ./release/

cp -r web/public release
tailwindcss --minify -c web/tailwind.config.js -i web/input.css -o release/public/assets/output.css