protoc:
	rm -f spec/*.pb.go
	protoc spec/spec.proto --go_out=spec
	mv spec/whatsmyip/spec/* spec
	rm -rf spec/whatsmyip
	make format
.PHONY: protoc

test:
	go test -coverprofile=cover.out ./...
	go tool cover -func=cover.out | grep "total:" | grep "(statements)"
.PHONY: test

benchmark:
	go test -run=- -benchmem -bench . github.com/wafer-bw/whatsmyip/...
.PHONY: benchmark

lint:
	go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest run
.PHONY: lint

format:
	go fix ./...
	go fmt ./...
.PHONY: format

api:
	go run main.go
.PHONY: api

run:
	vercel dev
.PHONY: run
