DOCKERFILE = ./docker/Dockerfile
IMAGE_NAME ?= cv-builder
IMAGE_TAG ?= latest
OUTPUT_DIR = __output/

.PHONY: all image publish tex pdf clean

all: pdf

image:
	@docker build -f $(DOCKERFILE) -t $(IMAGE_NAME):$(IMAGE_TAG) .

publish: image
	@docker push $(IMAGE_NAME):$(IMAGE_TAG)

tex:
	@go run .

pdf: tex
	@mtxrun --path=$(OUTPUT_DIR) --script context cv.tex

clean:
	@rm -rf $(OUTPUT_DIR)
