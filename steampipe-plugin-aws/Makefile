.PHONY: install build clean

build:
	GOPRIVATE="github.com/kaytu-io" CC=/usr/bin/musl-gcc GOPRIVATE="github.com/kaytu-io" GOOS=linux GOARCH=amd64 go build -v -o ./build/steampipe-plugin-aws.plugin *.go

docker-build:
	docker build -f docker/Dockerfile . -t steampipe-plugin-aws:0.0.1
	docker build -f docker/Dockerfile . -t steampipe-plugin-aws:0.0.1-$CI_PIPELINE_IID

docker-push:
	docker push steampipe-plugin-aws:0.0.1
	docker push steampipe-plugin-aws:0.0.1-$CI_PIPELINE_IID

podman-build:
	podman build -f docker/Dockerfile . -t steampipe-plugin-aws:0.0.1
	podman build -f docker/Dockerfile . -t steampipe-plugin-aws:0.0.1-$CI_PIPELINE_IID

podman-push:
	podman push steampipe-plugin-aws:0.0.1
	podman push steampipe-plugin-aws:0.0.1-$CI_PIPELINE_IID

clean:
	rm -r ./build

install:
	go build -v -o ~/.steampipe/plugins/hub.steampipe.io/plugins/turbot/aws@latest/steampipe-plugin-aws.plugin *.go
