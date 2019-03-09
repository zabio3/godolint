godolint
===
[![GitHub release](http://img.shields.io/github/release/zabio3/godolint.svg?style=flat-square)](release)
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

```
$ godolint testdata/src/DL3000_Dockerfile
testdata/src/DL3000_Dockerfile:3 DL3000 Use absolute WORKDIR

$ godolint testdata/src/DL3001_Dockerfile
testdata/src/DL3001_Dockerfile:6 DL3001 For some bash commands it makes no sense running them in a Docker container like `ssh`, `vim`, `shutdown`, `service`, `ps`, `free`, `top`, `kill`, `mount`, `ifconfig`
```

#### Options

```
Available options:
  --ignore RULECODE        A rule to ignore. If present, the ignore list in the
                           config file is ignored
```

##### Example

```
$ godolint --ignore DL3000 testdata/src/DL3000_Dockerfile
```

## Rules

An incomplete list of implemented rules.

| Rule                                                         | Description                                                                                                                                         |
|:-------------------------------------------------------------|:----------------------------------------------------------------------------------------------------------------------------------------------------|
| DL3000   | Use absolute WORKDIR.                                                                                                                               |
| DL3001   | For some bash commands it makes no sense running them in a Docker container like ssh, vim, shutdown, service, ps, free, top, kill, mount, ifconfig. |
| DL3002   | Last user should not be root.                                                                                                                       |
| DL3003   | Use WORKDIR to switch to a directory.                                                                                                               |