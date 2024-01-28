
# Running
Make sure you have all requirements installed:
```sh
go install github.com/cosmtrek/air@latest
go install github.com/a-h/templ/cmd/templ@latest
npm i -g tailwindcss
```
Then run the application with code changes watch:
`air`

# Build
```sh
mkdir build
mkdir build/web
mkdir build/tmp
rm -rf web/public/upload
cp -r web/public build/web/public
cp -r tmp/tailmplx build/tailmplx
zip -r build build
rm -rf build
```