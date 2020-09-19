# whatsmyip
Golang API running as a [Vercel Serverless Function](https://vercel.com/docs/serverless-functions) which returns your public IP address

![tests](https://github.com/wafer-bw/whatsmyip/workflows/tests/badge.svg)
<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Coverage-93%25-brightgreen.svg?longCache=true&style=flat)</a>
![CodeQL](https://github.com/wafer-bw/whatsmyip/workflows/CodeQL/badge.svg)
<a href='https://goreportcard.com/report/github.com/wafer-bw/whatsmyip' target='_blank'>![go report](https://goreportcard.com/badge/github.com/wafer-bw/whatsmyip)</a>

## Links
* Homepage - https://whatsmyip.wafer-bw.vercel.app
* API - https://whatsmyip.wafer-bw.vercel.app/api - supports the following `accept` headers:
    * `application/json`
    * `application/protobuf`

## Prerequisites
* [Go](https://golang.org/)
* [Node.js and npm](https://nodejs.org/en/) (for running [Vercel](https://vercel.com/))
* [protoc-gen-go](https://developers.google.com/protocol-buffers/docs/gotutorial) (if you wish to change the `protobuf` API contract within [./spec](./spec)).

## Setup
```bash
# Clone the repo and enter directory
git clone git@github.com:wafer-bw/whatsmyip.git
cd whatsmyip
# Get Vercel and login
npm i -g vercel
vercel login
```

## Usage
```bash
# Run the webpage & API as they would when deployed to Vercel
make run
# Run tests & benchmarks
make test
# Run just the API locally
make api
# Update coverage badge
make coverage
# Update protobuf API contract
make protoc
```
