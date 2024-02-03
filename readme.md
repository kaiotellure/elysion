
# Running
**1.** Make sure you have all CLI requirements installed: `./scripts/tidy.sh`<br/>
**2.** Then run the application: `air`<br/>

# Deploying
> [!WARNING]
> Please run `./scripts/tidy.sh` before commiting, so `.github/version` keeps track of the current version allowing `./scripts/release.sh` to perform the version linking correctly during `workflows/build.yml`.

> [!TIP]
> Download lastest build artifact produced by `workflows/build.yml`: [Nightly Link](https://nightly.link/ikaio/tailmplx/workflows/build/main/release.zip)

> [!IMPORTANT]
> This script is just what I'm currently using for my personal deployment, this is not by any form an enforced action, and it's highly recommended that you devote the time needed to research the best deployment strategy for the size of your application and its scalability plans.

```bash
# Backup your current version
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
