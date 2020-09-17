protoc:
	rm -f spec/*.pb.go
	protoc spec/spec.proto --go_out=plugins=grpc:spec --experimental_allow_proto3_optional
	mv spec/whatsmyip/spec/* spec
	rm -rf spec/whatsmyip
.PHONY: protoc

test:
	go test -v -coverprofile=cover.out ./...
.PHONY: test

run:
	HTTP_PORT=8000 go run ./api/main.go
.PHONY: run
