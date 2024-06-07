# whatsmyip
Golang API running as a [Vercel Serverless Function](https://vercel.com/docs/functions/serverless-functions) which returns your public IP address

## Links
* [Homepage](https://whatsmyip.wafer-bw.vercel.app)
* [API](https://whatsmyip.wafer-bw.vercel.app/api) - supports the following `accept` headers:
    * `application/json`
    * `application/protobuf`
    * defaults to `text/plain`

## Prerequisites
* Required
    * [Go](https://golang.org/)
* Optional
    * [Node.js and npm](https://nodejs.org/en/) for running/deploying with [Vercel](https://vercel.com/)
    * [protoc-gen-go](https://developers.google.com/protocol-buffers/docs/gotutorial) for updating the `protobuf` API contract found [here](./spec/spec.proto)

## Setup
```sh
# Clone the repo and enter directory
git clone git@github.com:wafer-bw/whatsmyip.git
cd whatsmyip
# Get Vercel and login
npm i -g vercel
vercel login
# Link project to your vercel account
vercel
```

## Usage
If you don't have `make`, you can just run the commands found in the [Makefile](./Makefile) directly
```sh
# Run the webpage & API
make run
# Run tests & benchmarks
make test
# Run the API only
make api
# Update coverage badge
make coverage
# Update protobuf API contract
make protoc
# Run benchmarks
make benchmark
# Run formatting
make format
# Run linting
make lint
```

If you have issues with `make protoc`:
```sh
go get github.com/golang/protobuf/protoc-gen-go
export PATH=$PATH:$HOME/go/bin
export PATH=$PATH:/usr/local/go/bin
```
