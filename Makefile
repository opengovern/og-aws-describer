.PHONY: build

build:
	export CGO_ENABLED=0
	export GOOS=linux
	export GOARCH=amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -ldflags "-w -extldflags -static" -o ./build/kaytu-aws-describer ./command/lambda/main.go
	cd build
	zip ./kaytu-aws-describer.zip ./kaytu-aws-describer
	aws s3 cp ./kaytu-aws-describer.zip s3://lambda-describe-binary/kaytu-aws-describer.zip
