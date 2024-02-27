#!/bin/bash

available() {
    # check path for arg 1, don't show stdout output, redirect stderr to stdout
    command -v "$1" >/dev/null 2>&1
}

error() {
    echo "$1" >&2;
    exit 1;
}

if ! available "go"; then
    error "Please install go."
fi

if ! available "npm"; then
    error "Please install node (npm)."
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

commit_count=$(git rev-list --all --count);
version="0.1.$(($commit_count + 1))";