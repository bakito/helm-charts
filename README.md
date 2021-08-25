# helm-charts

## Setup
Install [chart-releaser](https://github.com/helm/chart-releaser)
~/.cr/cr.yaml

```ỳaml
owner: bakito
git-repo: helm-charts
package-path: /home/xxx/path/to/this/repo/packages
token: ghp_...
git-base-url: https://api.github.com/
git-upload-url: https://uploads.github.com/
charts-repo: https://bakito.github.io/helm-charts/
index-path: /home/xxx/path/to/this/repo/docs/index.yaml

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
```

In this repo
```bash
git add .
git commit -m""
git push
```
