Go + HTMX + Templ + Tailwind

```cli
go install github.com/cosmtrek/air@latest
go install github.com/a-h/templ/cmd/templ@latest

npm install -g pnpm
cd styles
pnpm i

cd ..
air
```

get tailwind vscode working
```json
// settings.json
"tailwindCSS.includeLanguages": {
    "templ": "html"
}
```