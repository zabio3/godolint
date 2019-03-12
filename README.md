godolint
===
[![GitHub release](http://img.shields.io/github/release/zabio3/godolint.svg?style=flat-square)](https://github.com/zabio3/godolint/releases/latest)
[![Build Status](https://travis-ci.org/zabio3/godolint.svg?branch=master)](https://travis-ci.org/zabio3/godolint)
[![codecov](https://codecov.io/gh/zabio3/godolint/branch/master/graph/badge.svg)](https://codecov.io/gh/zabio3/godolint)
[![Go Report Card](https://goreportcard.com/badge/github.com/zabio3/godolint)](https://goreportcard.com/report/github.com/zabio3/godolint)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

A Dockerfile linter that helps you build [best practice](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/) Docker images. 
For static analysis of AST, [moby/buildkit](https://github.com/moby/buildkit/tree/master/frontend/dockerfile/parser) is used, and lint check is done.
The tool is parsing the Dockerfile into an AST and performs rules on top of the AST.(inspired by [Haskell Dockerfile Linter](https://github.com/hadolint/hadolint))

## Usage

You can run godolint locally to lint your Dockerfile.

```
$ godolint <Dockerfile>
```

##### Example

To check Dockerfile

```
$ godolint testdata/src/DL3000_Dockerfile
testdata/src/DL3000_Dockerfile:3 DL3000 Use absolute WORKDIR

$ godolint testdata/src/DL3001_Dockerfile
testdata/src/DL3001_Dockerfile:6 DL3001 For some bash commands it makes no sense running them in a Docker container like `ssh`, `vim`, `shutdown`, `service`, `ps`, `free`, `top`, `kill`, `mount`, `ifconfig`
```

#### Options

You can set some options:

```
Available options:
  --ignore RULECODE        A rule to ignore. If present, the ignore list in the
                           config file is ignored
```

##### Example

To check Dockerfile (exclude specific rules).

```
$ godolint --ignore DL3000 testdata/src/DL3000_Dockerfile
```

## Install

You can download binary from release page and place it in $PATH directory.

Or you can use go get

```
$ go get github.com/zabio3/godolint/cmd/godolint
```

## Rules

An incomplete list of implemented rules.

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

### AST

Dockerfile syntax is fully described in the [Dockerfile reference](https://docs.docker.com/engine/reference/builder/). 
Just take a look at [moby/buildkit](https://github.com/moby/buildkit/tree/master/frontend/dockerfile/parser) in the language-docker project to see the AST definition.

## Development

### Release Build

Make sure you have installed the goreleaser tool and then you can release gosec as follows:

```
$ git tag 1.0.0
$ export GITHUB_TOKEN=<YOUR GITHUB TOKEN>
$ goreleaser
```