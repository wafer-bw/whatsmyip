protoc:
	rm -f spec/*.pb.go
	protoc spec/spec.proto --go_out=spec
	mv spec/whatsmyip/spec/* spec
	rm -rf spec/whatsmyip
.PHONY: protoc

test:
	go test -coverprofile=cover.out ./...
.PHONY: test

benchmark:
	go test -run=- -benchmem -bench . github.com/wafer-bw/whatsmyip/...
.PHONY: benchmark

lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run
.PHONY: lint

api:
	go run main.go
.PHONY: api

run:
	vercel dev
.PHONY: run
