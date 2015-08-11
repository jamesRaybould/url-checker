#URL checker
A quick and dirty way of hitting a set of URLs written in Golang

##Usage
`urlChecker --csvPath=urls.csv`

Where `urls.csv` is a list of valid URLs with a line break after each URL

##Running
- Make sure that you have `golang` installed (using `brew install go` would be my recommended way on OSX)
- Change into the directory on the command line
- run `export GOPATH=$PWD && export GOBIN=$PWD/bin`
- `go get url-checker`
- `go run src/url-checker/urlChecker.go`

##Building
As above but run `go build src/github.com/jamesRaybould/url-checker/urlChecker.go`

This will give you an executable in the current directory that can be run on *any* osx machine
