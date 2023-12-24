Go + HTMX + Templ + Tailwind

# Setup
```sh
go install github.com/cosmtrek/air@latest
go get

npm install -g pnpm
cd styles
pnpm i

cd ..
air
```

## TailwindCSS IntelliSense
Getting it working by mapping it to HTML.
```js
// settings.json
"tailwindCSS.includeLanguages": {
    "templ": "html"
}
```