.PHONY: clear-swag swag-init

clear-swag:
	@echo "clearing docs directory"
	rm -f ./api/docs/*

swag-init: clear-swag
	@echo "Generating swagger docs..."
	swag init -d api -g cmd/server/main.go -o api/docs
