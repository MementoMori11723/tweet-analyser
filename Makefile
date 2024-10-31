.DEFAULT_GOAL := all
IMAGE_NAME := tweet-analyzer
check-docker:
	@command -v docker >/dev/null 2>&1 || { \
		echo "Docker is not installed. Please install Docker to continue."; \
		exit 1; \
	}
build:
	@docker build -t $(IMAGE_NAME) .
run:
	@docker run --rm -it -p 5000:5000 $(IMAGE_NAME)
all: check-docker build run
.PHONY: check-docker build run all
