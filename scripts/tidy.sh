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

# Create version file based on the commit counts
commit_count=$(git rev-list --all --count)
version="0.1.$(($commit_count + 1))"
echo "Inserting $version into .github/version"
echo -n "$version" > .github/version