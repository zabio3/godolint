godolint
===
[![GitHub release](http://img.shields.io/github/release/zabio3/godolint.svg?style=flat-square)](https://github.com/zabio3/godolint/releases/latest)
[![Build Status](https://travis-ci.org/zabio3/godolint.svg?branch=master)](https://travis-ci.org/zabio3/godolint)
[![Go Report Card](https://goreportcard.com/badge/github.com/zabio3/godolint)](https://goreportcard.com/report/github.com/zabio3/godolint)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

A smarter Dockerfile linter that helps you build [best practice](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/) Docker images. 
The linter is parsing the Dockerfile into an AST and performs rules on top of the AST. 
(Affected tools: [hadolint](https://github.com/hadolint/hadolint/wiki))

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

Or you can use go get,

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


### AST

Dockerfile syntax is fully described in the [Dockerfile reference](https://docs.docker.com/engine/reference/builder/). 
Just take a look at [moby/buildkit](https://github.com/moby/buildkit/tree/master/frontend/dockerfile/parser) in the language-docker project to see the AST definition.