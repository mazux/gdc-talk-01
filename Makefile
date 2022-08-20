.SILENT:
.DEFAULT_GOAL := help

export USRID := $(shell id -u)
export GRPID := $(shell id -g)

.PHONY: help
help: ; $(info Usage:)
	echo "\033[0;32mmake config\033[0;0m             Make some pre-build configuration"
	echo ""

	echo "\033[0;32mmake build\033[0;0m              Build app"
	echo "\033[0;32mmake run\033[0;0m                Run app"
	echo ""

	echo "\033[0;32mmake docker-build\033[0;0m       Build app through docker"
	echo "\033[0;32mmake docker-run\033[0;0m         Run app through docker"
	echo ""

.PHONY: run
run: ;
	go run ./...

.PHONY: build
build: ; $(info building app...)
	go build -o ./app ./cmd/app/main.go
	echo "Built successfully!"

.PHONY: config
config: ; $(info Preparing configuration...)
	cp -i .env.dist .env
	echo "Done!"

.PHONY: docker-build
docker-build: ; $(info building app through docker...)
	docker build -f Dockerfile -t app .
	echo "Built successfully!"

.PHONY: docker-run
docker-run: ;
	docker run -it app ./binary
