
# Running
**1.** Make sure you have all CLI requirements installed: `./scripts/tidy.sh`<br/>
**2.** Then run the application: `air`<br/>

# Deploying
> [!TIP]
> Download lastest build artifact produced by `build.yml`: [Nightly Link](https://nightly.link/ikaio/tailmplx/workflows/build/main/release.zip)

```bash
# Backup current version
cp -r nalvok nalvok-backup
# Fetch latest build by GH Actions (-L = follow redirects)
curl -L -o update.zip https://nightly.link/ikaio/tailmplx/workflows/build/main/release.zip
# Unzip overwriting (-o) at the nalvok/ dir
unzip -o update.zip -d nalvok/

cd nalvok
# Make sure the executable has permission to run
chmod u+x tailmplx
# Kill old running version
pkill tailmplx
# Run new version on background
DATABASE=../main.production.db PUBLIC_FOLDER=./public nohup ./tailmplx &
```
