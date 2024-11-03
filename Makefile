.DEFAULT_GOAL := all
IMAGE_NAME := tweet-analyzer
PORT ?= 5000
check-docker:
	@command -v docker >/dev/null 2>&1 || { \
		echo "Docker is not installed. Please install Docker to continue."; \
		exit 1; \
	}
build:
	@docker build --build-arg API_KEY=$(API_KEY) --build-arg PORT=$(PORT) -t $(IMAGE_NAME) .
run:
	@docker run -e API_KEY=$(API_KEY) -e PORT=$(PORT) --rm -it -p $(PORT):$(PORT) $(IMAGE_NAME)
all: check-docker build run
.PHONY: check-docker build run all
