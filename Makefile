.PHONY: build build-cli docker

build:
	export GOOS=linux
	export GOARCH=amd64
	CC=/usr/bin/musl-gcc GOPRIVATE="github.com/kaytu-io" GOOS=linux GOARCH=amd64 go build -v -ldflags "-linkmode external -extldflags '-static' -s -w" -tags musl -tags lambda.norpc -o ./build/kaytu-aws-describer ./main.go

docker:
	docker build -t kaytu-aws-describer:latest .

build-cli:
	export CGO_ENABLED=0
	export GOOS=linux
	export GOARCH=amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -ldflags "-w -extldflags -static" -o ./build/kaytu-aws-cli ./command/main.go
	cd build && zip ./kaytu-aws-cli.zip ./kaytu-aws-cli
	scp ./build/kaytu-aws-cli.zip steampipe:
	ssh steampipe unzip -o kaytu-aws-cli.zip
