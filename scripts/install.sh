#!/bin/bash
available() {
    command -v "$1" >/dev/null 2>&1
}

if ! available "go"; then
    echo "Please install go."
    exit 1
fi

if ! available "npm"; then
    echo "Please install node (npm)."
    exit 1
fi

if ! available "air"; then
    echo "Installing: air"
    go install github.com/cosmtrek/air@latest
fi

if ! available "templ"; then
    echo "Installing: templ"
    go install github.com/a-h/templ/cmd/templ@latest
fi

if ! available "tailwindcss"; then
    echo "Installing: tailwindcss (cli)"
    npm i -g tailwindcss
fi