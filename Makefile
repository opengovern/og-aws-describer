.PHONY: build build-cli

build:
	export CGO_ENABLED=0
	export GOOS=linux
	export GOARCH=amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -ldflags "-w -extldflags -static" -o ./build/kaytu-aws-describer ./main.go
	cd build && zip ./kaytu-aws-describer.zip ./kaytu-aws-describer
	aws s3 cp ./build/kaytu-aws-describer.zip s3://lambda-describe-binary/kaytu-aws-describer.zip --cli-read-timeout 300
	aws lambda update-function-code --function-name DescribeAWS --s3-bucket lambda-describe-binary --s3-key kaytu-aws-describer.zip --no-cli-pager --no-cli-auto-prompt
	aws lambda update-function-code --function-name kaytu-aws-describer --s3-bucket lambda-describe-binary --s3-key kaytu-aws-describer.zip --no-cli-pager --no-cli-auto-prompt

build-cli:
	export CGO_ENABLED=0
	export GOOS=linux
	export GOARCH=amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -ldflags "-w -extldflags -static" -o ./build/kaytu-aws-cli ./command/main.go
	cd build && zip ./kaytu-aws-cli.zip ./kaytu-aws-cli
	scp ./build/kaytu-aws-cli.zip steampipe:
	ssh steampipe unzip -o kaytu-aws-cli.zip
