Go + HTMX + Templ + Tailwind

```cli
go install github.com/cosmtrek/air@latest
go get github.com/a-h/templ/cmd/templ@latest
go get github.com/go-chi/chi/v5

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