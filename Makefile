.PHONY: install clean gen build

default: build

install:
	@go install golang.org/x/tools/cmd/stringer@latest

clean:
	@rm -f gobgp
	@rm -f gobgpd

gen:
	@protoc -I api \
		-I /usr/local/include \
    	--go_out=api \
		--go_opt=paths=source_relative \
		--go-grpc_out=api \
		--go-grpc_opt=paths=source_relative \
		api/attribute_dmwg.proto
	@go generate ./pkg/packet/bgp/bgp.go

build:
	@go mod tidy
	@docker build --target export --output . .
