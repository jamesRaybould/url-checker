#URL checker
A quick and dirty way of hitting a set of URLs written in Golang

##Usage
`urlChecker --csvPath=urls.csv` or `urlChecker --api`

Where `urls.csv` is a list of valid URLs with a line break after each URL

The assumption is that the API you are calling matches the API this was orignally designed for.

##Developing
- Make sure that you have `golang` installed (using `brew install go` would be my recommended way on OSX)
- Change into the directory on the command line
- run `export GOPATH=$PWD && export GOBIN=$PWD/bin`
- `go get url-checker`
- `go run src/url-checker/urlChecker.go --api`

##Building
As above but run `go build src/github.com/jamesRaybould/url-checker/urlChecker.go`

This will give you an executable in the current directory that can be run on *any* osx machine
