GO := go

detect-gopath:
    ifndef $(GOPATH)
        GOPATH=$(shell go env GOPATH)
        export GOPATH
    endif

generate-grpc: detect-gopath
	PATH=$(GOPATH)/bin:${PATH} && \
	$(GO) get github.com/golang/protobuf/protoc-gen-go && \
	$(GO) get google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.0 && \
	protoc --go_out=. \
 		--go_opt=paths=source_relative \
    	--go-grpc_out=. \
    	--go-grpc_opt=paths=source_relative \
    	protos/stock.proto

generate-php-client:
	protoc --php_out=./examples/php --grpc_out=./examples/php \
	--plugin=protoc-gen-grpc=/home/cmuench/Workspaces/playground/grpc/bins/opt/grpc_php_plugin \
 	protos/stock.proto

build-server: detect-gopath
	$(GO) build server.go

build-client: detect-gopath
	$(GO) build client.go

run-server: detect-gopath
	$(GO) run server.go

run-client: detect-gopath
	$(GO) run client.go

all: detect-gopath generate-grpc generate-php-client build-server build-client
