OUTPUT_DIR = __output/

.PHONY: all tex pdf clean

all: pdf

tex:
	@go run .

pdf: tex
	@mtxrun --path=$(OUTPUT_DIR) --script context cv.tex

clean:
	@rm -rf $(OUTPUT_DIR)
