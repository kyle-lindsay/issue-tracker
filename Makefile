.PHONY: build dev

build:
	@echo "Building binary..."
	go build -o issue-tracker

dev: build
	@echo "App compiled. To load autocomplete, run this command:"
	@echo ""
	@echo "		export PATH=\$$PATH:\$$(pwd) && source <(issue-tracker completion bash)"
	@echo ""
