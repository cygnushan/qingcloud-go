# Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a Apache
# license that can be found in the LICENSE file.

GO_LDFLAGS:=--ldflags '-w -s -extldflags "-static"'
GO:=docker run --rm -v $(shell go env GOPATH):/go -w /go/src/$(shell go list) golang:alpine go

default:
	go generate ./pkg/...
	go fmt ./...
	go test ./...

build-docker: Dockerfile
	$(GO) fmt ./...
	docker build -t chai2010/qingcloud-go .
	@docker image prune -f 1>/dev/null 2>&1

run-docker:
	docker run --rm -it -v `pwd`:/root -w /root chai2010/qingcloud-go qcli

hello:
	go fmt ./...
	go run hello.go

demo_mgo:
	go run examples/mongo/DescribeMongos/DescribeMongos.go

demo_cluster:
	go run examples/cluster/DescribeClusterNodes/DescribeClusterNodes.go

# go run examples/cluster/DescribeClusters/DescribeClusters.go

demo:
	go run examples/nic/DescribeNics/DescribeNics.go

init-vendor:
	govendor init
	govendor add +external
	@echo "ok"

update-vendor:
	govendor update +external
	govendor list
	@echo "ok"

remove-unused-vendor:
	govendor remove +u
	@echo "ok"

tools:
	go get github.com/kardianos/govendor
	docker pull golang:alpine
	@echo "ok"

generate:
	cd ./api && make
	go generate ./pkg/...
	go fmt  ./...
	go test ./...

fmt:
	go fmt ./...

test:
	go generate ./pkg/...
	go fmt  ./...
	go vet ./...
	go test ./...

dist:
	go generate ./pkg/...

	mkdir -p ./bin/qingcloud-cli-windows-amd64
	mkdir -p ./bin/qingcloud-cli-linux-amd64
	mkdir -p ./bin/qingcloud-cli-darwin-amd64

	GOOS=windows GOARCH=amd64 go build -o ./bin/qingcloud-cli-windows-amd64/qcli.exe ./cmd/qcli
	GOOS=linux   GOARCH=amd64 go build -o ./bin/qingcloud-cli-linux-amd64/qcli       ./cmd/qcli
	GOOS=darwin  GOARCH=amd64 go build -o ./bin/qingcloud-cli-darwin-amd64/qcli      ./cmd/qcli

	cp README.md ./bin/qingcloud-cli-windows-amd64/
	cp README.md ./bin/qingcloud-cli-linux-amd64/
	cp README.md ./bin/qingcloud-cli-darwin-amd64/

	cp CHANGES.md ./bin/qingcloud-cli-windows-amd64/
	cp CHANGES.md ./bin/qingcloud-cli-linux-amd64/
	cp CHANGES.md ./bin/qingcloud-cli-darwin-amd64/

	cp LICENSE ./bin/qingcloud-cli-windows-amd64/
	cp LICENSE ./bin/qingcloud-cli-linux-amd64/
	cp LICENSE ./bin/qingcloud-cli-darwin-amd64/

	cd ./bin && zip -r qingcloud-cli-windows-amd64.zip qingcloud-cli-windows-amd64
	cd ./bin && zip -r qingcloud-cli-linux-amd64.zip   qingcloud-cli-linux-amd64
	cd ./bin && zip -r qingcloud-cli-darwin-amd64.zip  qingcloud-cli-darwin-amd64

	@echo "ok"

clean:
	make -C ./pkg/api clean
	make -C ./pkg/cmd/qcli/api clean
