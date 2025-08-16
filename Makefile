.POSIX:

.PHONY: all
all: build format validate

.PHONY: install
install:
	command -v tofu >/dev/null 2>&1 || go install github.com/opentofu/tofu@latest
	command -v terraform-docs >/dev/null 2>&1 || go install github.com/terraform-docs/terraform-docs@v0.16.0
	command -v tfsec >/dev/null 2>&1 || go install github.com/aquasecurity/tfsec/cmd/tfsec@latest
	command -v tflint >/dev/null 2>&1 || go install github.com/terraform-linters/tflint@v0.38.1

.PHONY: build
build: install generate

.PHONY: generate
generate:
	go run main.go

.PHONY: format
format:
	tofu fmt

.PHONY: init
init:
	tofu init -no-color

.PHONY: validate
validate: init
	tofu fmt --check
	tofu validate -no-color
	tflint --no-color