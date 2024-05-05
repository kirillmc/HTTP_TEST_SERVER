LOCAL_BIN:=$(CURDIR)/bin

docker-build:
	docker buildx build --no-cache --platform linux/amd64 -t http_test_server:v0.0.1 .