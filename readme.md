
# Running
Make sure you have all CLI requirements installed:
```sh
go install github.com/cosmtrek/air@latest
go install github.com/a-h/templ/cmd/templ@latest
npm i -g tailwindcss
```
Then run the application with: `air`

# Tips
- Linux: Find process: `ps -A | grep tailmplx`
- Linux: Run process on background: `nohup ./tailmplx &`

Download lastest build: [Nightly Link](https://nightly.link/ikaio/tailmplx/workflows/build/main/release.zip)

# AWS Pull Update
```bash
cp -r nalvok nalvok-backup
curl -L -o update.zip https://nightly.link/ikaio/tailmplx/workflows/build/main/release.zip
unzip -o update.zip -d nalvok/
cd nalvok
chmod u+x tailmplx
nohup ./tailmplx &
```
