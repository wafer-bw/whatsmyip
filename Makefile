protoc:
	rm -f spec/*.pb.go
	protoc spec/spec.proto --go_out=spec
	mv spec/whatsmyip/spec/* spec
	rm -rf spec/whatsmyip
.PHONY: protoc

test:
	go test -coverprofile=cover.out ./...
.PHONY: test

api:
	go run main.go
.PHONY: api

run:
	vercel dev
.PHONY: run

coverage:
	gopherbadger -md="README.md" -png=false -prefix ""
.PHONY: coverage
