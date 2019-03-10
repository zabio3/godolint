godolint
===
[![GitHub release](http://img.shields.io/github/release/zabio3/godolint.svg?style=flat-square)](https://github.com/zabio3/godolint/releases/latest)
[![Build Status](https://travis-ci.org/zabio3/godolint.svg?branch=master)](https://travis-ci.org/zabio3/godolint)
[![codecov](https://codecov.io/gh/zabio3/godolint/branch/master/graph/badge.svg)](https://codecov.io/gh/zabio3/godolint)
[![Go Report Card](https://goreportcard.com/badge/github.com/zabio3/godolint)](https://goreportcard.com/report/github.com/zabio3/godolint)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

**Comming soon !!**

A smarter Dockerfile linter that helps you build [best practice](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/) Docker images. 
The linter is parsing the Dockerfile into an AST and performs rules on top of the AST. 
(inspired by [hadolint](https://github.com/hadolint/hadolint/wiki))

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
| DL3000   | Use absolute WORKDIR.                                                                                                                               |
| DL3001   | For some bash commands it makes no sense running them in a Docker container like ssh, vim, shutdown, service, ps, free, top, kill, mount, ifconfig. |
| DL3002   | Last user should not be root.                                                                                                                       |
| DL3003   | Use WORKDIR to switch to a directory.                                                                                                               |
| DL3004   | Do not use sudo as it leads to unpredictable behavior. Use a tool like gosu to enforce root.                                                        |
| DL3005   | Do not use apt-get upgrade or dist-upgrade.                                                                                                         |
| DL3006   | Always tag the version of an image explicitly.      
| DL3007   | Using latest is prone to errors if the image will ever update. Pin the version explicitly to a release tag.                                         |                                                                                                |
| DL3008   | Pin versions in apt-get install.                                                                                                                    |
| DL3009 (Unimplemented) | Delete the apt-get lists after installing something.                                                                                                |
| DL3010 (Unimplemented) | Use ADD for extracting archives into an image.                                                                                                      |
| DL3011 (Unimplemented) | Valid UNIX ports range from 0 to 65535.                                                                                                             |
| DL3012 (Unimplemented) | Provide an email address or URL as maintainer.                                                                                                      |
| DL3013 (Unimplemented) | Pin versions in pip.                                                                                                                                |
| DL3014 (Unimplemented) | Use the `-y` switch.                                                                                                                                |
| DL3015 (Unimplemented) | Avoid additional packages by specifying --no-install-recommends.                                                                                    |
| DL3016 (Unimplemented) | Pin versions in `npm`.                                                                                                                              |
| DL3017 (Unimplemented) | Do not use `apk upgrade`.                                                                                                                           |
| DL3018 (Unimplemented) | Pin versions in apk add. Instead of `apk add <package>` use `apk add <package>=<version>`.                                                          |
| DL3019 (Unimplemented) | Use the `--no-cache` switch to avoid the need to use `--update` and remove `/var/cache/apk/*` when done installing packages.                        |
| DL3020 (Unimplemented) | Use `COPY` instead of `ADD` for files and folders.                                                                                                  |
| DL3021 (Unimplemented) | `COPY` with more than 2 arguments requires the last argument to end with `/`                                                                        |
| DL3022 (Unimplemented) | `COPY --from` should reference a previously defined `FROM` alias                                                                                    |
| DL3023 (Unimplemented) | `COPY --from` cannot reference its own `FROM` alias                                                                                                 |
| DL3024 (Unimplemented) | `FROM` aliases (stage names) must be unique                                                                                                         |
| DL3025 (Unimplemented) | Use arguments JSON notation for CMD and ENTRYPOINT arguments                                                                                        |
| DL3026 (Unimplemented) | Use only an allowed registry in the FROM image                                                                                                      |
| DL4000 (Unimplemented) | MAINTAINER is deprecated.                                                                                                                           |
| DL4001 (Unimplemented) | Either use Wget or Curl but not both.                                                                                                               |
| DL4003 (Unimplemented) | Multiple `CMD` instructions found.                                                                                                                  |
| DL4004 (Unimplemented) | Multiple `ENTRYPOINT` instructions found.                                                                                                           |
| DL4005 (Unimplemented) | Use `SHELL` to change the default shell.                                                                                                            |
| DL4006 (Unimplemented) | Set the `SHELL` option -o pipefail before `RUN` with a pipe in it                                                                                   |

### AST

Dockerfile syntax is fully described in the [Dockerfile reference](https://docs.docker.com/engine/reference/builder/). 
Just take a look at [moby/buildkit](https://github.com/moby/buildkit/tree/master/frontend/dockerfile/parser) in the language-docker project to see the AST definition.