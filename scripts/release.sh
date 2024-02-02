#!/bin/bash
available() {
    command -v "$1" >/dev/null 2>&1
}

mkdir release
go build -o ./release/

cp -r web/public release
tailwindcss --minify -c web/tailwind.config.js -i web/input.css -o release/public/assets/output.css