#!/bin/bash
available() {
    command -v "$1" >/dev/null 2>&1
}

commit_count=$(git rev-list --all --count)
version="0.1.$(($commit_count + 1))"
echo "Creating $version release."

mkdir release 2>/dev/null
go build -ldflags="-X help.Version=$version" -o ./release/

cp -r web/public release
tailwindcss --minify -c web/tailwind.config.js -i web/input.css -o release/public/assets/output.css