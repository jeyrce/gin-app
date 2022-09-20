module:=github.com/jeyrce/gin-app
hub:=$(shell echo "${module}" | awk -F/ '{print $$2}')
app:=$(shell echo "${module}" | awk -F/ '{print $$NF}')
version:=latest
buildAt:=$(shell date "+%Y-%m-%d_%H:%M:%S")
commitId:=$(shell git rev-parse --short HEAD)
branch:=$(shell git symbolic-ref --short -q HEAD)

.phony: all
all: swagger image


.phony: swagger
swagger:
	@echo "生成最新swagger描述文档"
	swag init -o api/v1/ -g url.go --ot go -d api/v1/ --instanceName V1
	swag init -o api/v2/ -g url.go --ot go -d api/v2/ --instanceName V2


.phony: image
image:
	@-docker buildx rm tmp
	docker buildx create --name tmp --bootstrap --use
	docker buildx build -t ${hub}/${app}:${version} \
		--build-arg module=${module} \
		--build-arg app=${app} \
		--build-arg commitId=${commitId} \
		--build-arg goProxy=${goProxy} \
		--platform linux/386,linux/amd64,linux/arm64 \
		--push \
		.
	@-docker buildx rm tmp

.phony: binary
binary:
	@echo ${hub}
	@rm -rf _out/*
	CGO_ENABLED=0 go build -ldflags " \
		-X '${module}/pkg/conf.metaCommitId=${commitId}' \
		-X '${module}/pkg/conf.metaBranch=${branch}' \
		-X '${module}/pkg/conf.metaVersion=${version}' \
		-X '${module}/pkg/conf.metaBuildAt=${buildAt}' \
		-X '${module}/pkg/conf.metaPoweredBy=${module}' \
	" \
	-o _out/${app} *.go


.phony: init
init:
	@bash init.sh 
