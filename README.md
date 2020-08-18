[![create-release](https://github.com/kasuboski/resume/workflows/create-release/badge.svg)](https://github.com/kasuboski/resume/actions?query=workflow%3Acreate-release)
## Current generation steps
* Edit `resume.json` or `hack/resume.html.tmpl` for data or template changes
* Create html with `go run hack/template.go`
* Get pdf from browser (or use workflow)

## Workflow
This repo has a GitHub Actions workflow that will generate the html and pdf versions of `resume.json`. It will also create a release and update my personal site when a tag is pushed.