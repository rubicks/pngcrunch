pngcrunch [![Build Status](https://travis-ci.org/rubicks/pngcrunch.svg?branch=master)](https://travis-ci.org/rubicks/pngcrunch)
======
Minify your png files.

### Install Go

`pngcrunch` is written in Go.

If you know already know Go, then you can stop reading because you already know how this
works.

If you don't know Go, the rest of this README is for you. You should

* install Go

* `export` your `GOPATH`

* learn how Go works; i.e., `go get`, `go build`, `go test`, and `go install`

Go works really hard to make your like easy, including the following...

### Get it

    $ go get github.com/rubicks/pngcrunch

### Test it

Simply

    $ go test github.com/rubicks/pngcrunch

or

    $ cd $GOPATH/src/github.com/rubicks/pngcrunch && go test

### Build it

in a temporary directory

    $ cd $(mktemp -d) && go build github.com/rubicks/pngcrunch

or

    $ cd $GOPATH/src/github.com/rubicks/pngcrunch && go build

### Install it

to your `PATH`, assuming you already

    export PATH=$PATH:$GOPATH/bin

somewhere, you can

    $ go install github.com/rubicks/pngcrunch

### Use it

Instructions forthcoming, I promise.

### Dockerize it

The `Dockerfile` for pngcrunch inherits from the
[docker hub official repo for golang](https://registry.hub.docker.com/_/golang). After
you

    $ go get github.com/rubicks/pngcrunch

you can

    $ docker build -t pngcrunch-golang-app $GOPATH/src/github.com/rubicks/pngcrunch

and

    $ docker run -it --rm pngcrunch-golang-app

just as you would normally.
