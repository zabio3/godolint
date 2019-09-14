godolint
===

[![GitHub release](http://img.shields.io/github/release/zabio3/godolint.svg?style=flat-square)](https://github.com/zabio3/godolint/releases/latest)
[![Build Status](https://travis-ci.org/zabio3/godolint.svg?branch=master)](https://travis-ci.org/zabio3/godolint)
[![Golang CI](https://golangci.com/badges/github.com/zabio3/godolint.svg)](https://golangci.com/r/github.com/zabio3/godolint)
[![codecov](https://codecov.io/gh/zabio3/godolint/branch/master/graph/badge.svg)](https://codecov.io/gh/zabio3/godolint)
[![Go Report Card](https://goreportcard.com/badge/github.com/zabio3/godolint)](https://goreportcard.com/report/github.com/zabio3/godolint)
[![GoDoc](https://godoc.org/github.com/zabio3/godolint?status.svg)](https://godoc.org/github.com/zabio3/godolint)
[![Maintainability](https://api.codeclimate.com/v1/badges/4c1c216781e5592d4194/maintainability)](https://codeclimate.com/github/zabio3/godolint/maintainability)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

A Dockerfile linter that helps you build [best practice](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/) Docker images (inspired by [Haskell Dockerfile Linter](https://github.com/hadolint/hadolint)). 
For static analysis of AST, [moby/buildkit parser](https://github.com/moby/buildkit/tree/master/frontend/dockerfile/parser) is used, and lint check is done.
This tool performs docker rule checks based on the parsed AST.

## Dependency
 
  - [moby/buildkit/frontend/dockerfile/parser](https://github.com/moby/buildkit)
     - The official Dockerfile parser provided by moby. Used when parsing Dockerfile.
          
## Usage

You can run godolint locally to lint your Dockerfile.

```
$ godolint <Dockerfile>
```

##### Example

To check Dockerfile

```
$ godolint testdata/DL3000_Dockerfile
#3 DL3000 Use absolute WORKDIR. 

$ godolint testdata/DL3001_Dockerfile
#6 DL3001 For some bash commands it makes no sense running them in a Docker container like `ssh`, `vim`, `shutdown`, `service`, `ps`, `free`, `top`, `kill`, `mount`, `ifconfig`. 
```

#### Options

You can set some options:

```
Available options:
  --ignore RULECODE     A rule to ignore. If present, the ignore list in the
                        config file is ignored

Other Commands:
  --help        -h      Help about any command
  --version     -v      Print the version information
```

##### Example

To check Dockerfile (exclude specific rules).

```
$ godolint --ignore DL3000 testdata/DL3000_Dockerfile
```

## Install

You can download binary from release page and place it in $PATH directory.

Or you can use go get

```
$ go get github.com/zabio3/godolint
```

## Rules

An implemented rules. (Dockerfile lint rule provided by [hadolint](https://github.com/hadolint/hadolint))

| Rule                                                         | Description                                                                                                                                         |
|:-------------------------------------------------------------|:----------------------------------------------------------------------------------------------------------------------------------------------------|
| [DL3000](https://github.com/hadolint/hadolint/wiki/DL3000)   | Use absolute WORKDIR.                                                                                                                               |
| [DL3001](https://github.com/hadolint/hadolint/wiki/DL3001)   | For some bash commands it makes no sense running them in a Docker container like ssh, vim, shutdown, service, ps, free, top, kill, mount, ifconfig. |
| [DL3002](https://github.com/hadolint/hadolint/wiki/DL3002)   | Last user should not be root.                                                                                                                       |
| [DL3003](https://github.com/hadolint/hadolint/wiki/DL3003)   | Use WORKDIR to switch to a directory.                                                                                                               |
| [DL3004](https://github.com/hadolint/hadolint/wiki/DL3004)   | Do not use sudo as it leads to unpredictable behavior. Use a tool like gosu to enforce root.                                                        |
| [DL3005](https://github.com/hadolint/hadolint/wiki/DL3005)   | Do not use apt-get upgrade or dist-upgrade.                                                                                                         |
| [DL3007](https://github.com/hadolint/hadolint/wiki/DL3007)   | Using latest is prone to errors if the image will ever update. Pin the version explicitly to a release tag.                                         |
| [DL3006](https://github.com/hadolint/hadolint/wiki/DL3006)   | Always tag the version of an image explicitly.                                                                                                      |
| [DL3008](https://github.com/hadolint/hadolint/wiki/DL3008)   | Pin versions in apt-get install.                                                                                                                    |
| [DL3009](https://github.com/hadolint/hadolint/wiki/DL3009)   | Delete the apt-get lists after installing something.                                                                                                |
| [DL3010](https://github.com/hadolint/hadolint/wiki/DL3010)   | Use ADD for extracting archives into an image.                                                                                                      |
| [DL3011](https://github.com/hadolint/hadolint/wiki/DL3011)   | Valid UNIX ports range from 0 to 65535.                                                                                                             |
| [DL3012](https://github.com/hadolint/hadolint/wiki/DL3012)   | Provide an email address or URL as maintainer. (This rule is DEPRECATED and no longer active)                                                       |
| [DL3013](https://github.com/hadolint/hadolint/wiki/DL3013)   | Pin versions in pip.                                                                                                                                |
| [DL3014](https://github.com/hadolint/hadolint/wiki/DL3014)   | Use the `-y` switch.                                                                                                                                |
| [DL3015](https://github.com/hadolint/hadolint/wiki/DL3015)   | Avoid additional packages by specifying --no-install-recommends.                                                                                    |
| [DL3016](https://github.com/hadolint/hadolint/wiki/DL3016)   | Pin versions in `npm`.                                                                                                                              |
| [DL3017](https://github.com/hadolint/hadolint/wiki/DL3017)   | Do not use `apk upgrade`.                                                                                                                           |
| [DL3018](https://github.com/hadolint/hadolint/wiki/DL3018)   | Pin versions in apk add. Instead of `apk add <package>` use `apk add <package>=<version>`.                                                          |
| [DL3019](https://github.com/hadolint/hadolint/wiki/DL3019)   | Use the `--no-cache` switch to avoid the need to use `--update` and remove `/var/cache/apk/*` when done installing packages.                        |
| [DL3020](https://github.com/hadolint/hadolint/wiki/DL3020)   | Use `COPY` instead of `ADD` for files and folders.                                                                                                  |
| [DL3021](https://github.com/hadolint/hadolint/wiki/DL3021)   | `COPY` with more than 2 arguments requires the last argument to end with `/`.                                                                       |
| [DL3022](https://github.com/hadolint/hadolint/wiki/DL3022)   | `COPY --from` should reference a previously defined `FROM` alias.                                                                                   |
| [DL3023](https://github.com/hadolint/hadolint/wiki/DL3023)   | `COPY --from` cannot reference its own `FROM` alias.                                                                                                |
| [DL3024](https://github.com/hadolint/hadolint/wiki/DL3024)   | `FROM` aliases (stage names) must be unique.                                                                                                        |
| [DL3025](https://github.com/hadolint/hadolint/wiki/DL3025)   | Use arguments JSON notation for CMD and ENTRYPOINT arguments.                                                                                       |
| [DL4000](https://github.com/hadolint/hadolint/wiki/DL4000)   | MAINTAINER is deprecated.                                                                                                                           |
| [DL4001](https://github.com/hadolint/hadolint/wiki/DL4001)   | Either use Wget or Curl but not both.                                                                                                               |
| [DL4003](https://github.com/hadolint/hadolint/wiki/DL4003)   | Multiple `CMD` instructions found.                                                                                                                  |
| [DL4004](https://github.com/hadolint/hadolint/wiki/DL4004)   | Multiple `ENTRYPOINT` instructions found.                                                                                                           |
| [DL4005](https://github.com/hadolint/hadolint/wiki/DL4005)   | Use `SHELL` to change the default shell.                                                                                                            |
| [DL4006](https://github.com/hadolint/hadolint/wiki/DL4006)   | Set the `SHELL` option -o pipefail before `RUN` with a pipe in it.                                                                                  |

### AST

Dockerfile syntax is fully described in the [Dockerfile reference](https://docs.docker.com/engine/reference/builder/). 
Just take a look at [moby/buildkit](https://github.com/moby/buildkit/tree/master/frontend/dockerfile/parser) in the language-docker project to see the AST definition.

## Contribution
Contributions are of course always welcome!

1. Fork zabio3/godolint (https://github.com/zabio3/godolint/fork)
2. Run `go get` to install dependencies
3. Create a feature branch
4. Commit your changes
5. Run test using `go test ./...`
6. Create a Pull Request

See [`CONTRIBUTING.md`](https://github.com/zabio3/godolint/blob/master/CONTRIBUTING.md) for details.