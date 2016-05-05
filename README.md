Charon [![GoDoc](https://godoc.org/github.com/piotrkowalczuk/charon?status.svg)](http://godoc.org/github.com/piotrkowalczuk/charon)&nbsp;[![Build Status](https://travis-ci.org/piotrkowalczuk/charon.svg?branch=master)](https://travis-ci.org/piotrkowalczuk/charon)&nbsp;[![codecov.io](https://codecov.io/github/piotrkowalczuk/charon/coverage.svg?branch=master)](https://codecov.io/github/piotrkowalczuk/charon?branch=master)&nbsp;[![Code Climate](https://codeclimate.com/github/piotrkowalczuk/charon/badges/gpa.svg)](https://codeclimate.com/github/piotrkowalczuk/charon)&nbsp;[![Go Report Card](https://goreportcard.com/badge/github.com/piotrkowalczuk/charon)](https://goreportcard.com/report/github.com/piotrkowalczuk/charon)
=============

<img src="/data/logo/charon.png?raw=true" width="300">

## Quick Start

### Installation

```bash
$ go install github.com/piotrkowalczuk/charon/cmd/charond
$ go install github.com/piotrkowalczuk/charon/cmd/charonctl
```

### Superuser

```bash
$ charonctl register -noauth -r.superuser=true -r.username="j.snow@gmail.com" -r.password=123 -r.firstname=John -r.lastname=Snow
```

## Contribution

@TODO

### Documentation

@TODO

### TODO
- [x] Auth
    - [x] login
    - [x] logout
    - [x] is authenticated
    - [x] subject
    - [x] is granted
    - [x] belongs to
- [x] Permission
	- [x] get
    - [x] list
    - [x] register
- [x] Group
    - [x] get
    - [x] list
    - [x] modify
    - [x] delete
    - [x] create
    - [x] set permissions
    - [x] list permissions
- [x] User
    - [x] get
    - [x] list
    - [x] modify
    - [x] delete
    - [x] create
    - [x] set permissions
    - [x] set groups
    - [x] list permissions
    - [x] list groups
