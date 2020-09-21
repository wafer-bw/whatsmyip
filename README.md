# whatsmyip
Golang API running as a [Vercel Serverless Function](https://vercel.com/docs/serverless-functions) which returns your public IP address

![tests](https://github.com/wafer-bw/whatsmyip/workflows/tests/badge.svg)
<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Coverage-93%25-brightgreen.svg?longCache=true&style=flat)</a>
![CodeQL](https://github.com/wafer-bw/whatsmyip/workflows/CodeQL/badge.svg)
<a href='https://goreportcard.com/report/github.com/wafer-bw/whatsmyip' target='_blank'>![go report](https://goreportcard.com/badge/github.com/wafer-bw/whatsmyip)</a>

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
```bash
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
```

## References
* I learned a lot of this from [@codeallthethingz](https://github.com/codeallthethingz)
