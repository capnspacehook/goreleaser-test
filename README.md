# go-project-template

Template repository for Go projects. Contains an example Go application and a minimal Dockerfile.

## Details

Dependabot is configured to keep Go, Docker and Github Actions dependencies up to date by opening
pull requests when new versions are released.

Multiple workflows are configured that will:

- Lint Go code
- Check that `go.mod` is tidied
- Test Go code and fuzz for 10 minutes
- Lint workflow files with [actionlint](https://github.com/rhysd/actionlint)
- Lint Dockerfile with [hadolint](https://github.com/hadolint/hadolint)
- Build, publish, sign and scan Docker images with [cosign](https://github.com/sigstore/cosign), [grype](https://github.com/anchore/grype) and [trivy](https://github.com/aquasecurity/trivy)
- Build, sign, publish binaries and create release with [goreleaser](https://github.com/goreleaser/goreleaser) and [cosign](https://github.com/sigstore/cosign)

Almost all workflows will trigger when appropriate files are modified from pushes or pull requests. 
Binaries will only be released when a semver compatible tag is pushed however.

## Usage

Workflow files will work without modification, as will releasing Docker images and binaries with goreleaser.
Note that only `linux/amd64` images and binaries are built by default, so you may need to add more target
operating systems and/or architectures based off of your requirements.

The workflow that tests and fuzzes Go code will still pass if no tests or fuzz tests are present, so when you
do add tests and fuzz tests the workflow will run them without needing any changes from you.
