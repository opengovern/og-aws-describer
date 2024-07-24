.PHONY: build build-cli docker

lambda-build:
	CC=/usr/bin/musl-gcc GOPRIVATE="github.com/kaytu-io" GOOS=linux GOARCH=amd64 go build -v -ldflags "-linkmode external -extldflags '-static' -s -w" -tags musl,lambda.norpc -o ./build/kaytu-aws-describer ./lambda/main.go

functions-build:
	CC=/usr/bin/musl-gcc GOPRIVATE="github.com/kaytu-io" GOOS=linux GOARCH=amd64 go build -v -ldflags "-linkmode external -extldflags '-static' -s -w" -tags musl -o ./azfunction/kaytu-aws-describer ./azfunction/main.go

local-build:
	CC=/usr/bin/musl-gcc GOPRIVATE="github.com/kaytu-io" GOOS=linux GOARCH=amd64 go build -v -ldflags "-linkmode external -extldflags '-static' -s -w" -tags musl -o ./local/kaytu-aws-describer ./local/main/main.go

lambda-docker:
	docker build -t 435670955331.dkr.ecr.us-east-2.amazonaws.com/kaytu-aws-describer:latest .
	docker push 435670955331.dkr.ecr.us-east-2.amazonaws.com/kaytu-aws-describer:latest

lambda-update:
	aws lambda update-function-code --function-name kaytu-aws-describer --image-uri 435670955331.dkr.ecr.us-east-2.amazonaws.com/kaytu-aws-describer:latest --region us-east-2

build-cli:
	export CGO_ENABLED=0
	export GOOS=linux
	export GOARCH=amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -ldflags "-w -extldflags -static" -o ./build/kaytu-aws-cli ./command/main.go
	cd build && zip ./kaytu-aws-cli.zip ./kaytu-aws-cli
	scp ./build/kaytu-aws-cli.zip steampipe:
	ssh steampipe unzip -o kaytu-aws-cli.zip
