# [helm-charts](https://bakito.github.io/helm-charts/)

## Setup

Install [chart-releaser](https://github.com/helm/chart-releaser)
env:
```ỳaml
export CR_TOKEN=${GITHUB_TOKEN}
export CR_PACKAGE_PATH=/home/xxx/path/to/this/repo/packages
```

~/.cr/cr.yaml

```ỳaml
owner: bakito
git-repo: helm-charts
token: ghp_...
git-base-url: https://api.github.com/
git-upload-url: https://uploads.github.com/
charts-repo: https://bakito.github.io/helm-charts/
index-path: /home/xxx/path/to/this/repo/docs/index.yaml
pages-branch: main
bakito @ ~/git/github
```

## Usage

In the chart repository

```bash
# build the package
cr package <chars-dir>

# upload the chart
cr upload --skip-existing

# update the index
cr index
go run ./cmd/index2md/ > docs/README.md
```

In this repo

```bash
git add .
git commit -m""
git push
```
