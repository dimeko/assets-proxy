CLIENT_DIR = ./client
CLIENT_BUILD_DIR = ./client/dist

run:
	docker-compose up -d
ifeq ($(wildcard $(CLIENT_BUILD_DIR)/.*),)
	npm run --prefix $(CLIENT_DIR) build
endif
	go run main.go

build-client:
	npm run --prefix $(CLIENT_DIR) build

run-dev-client:
	npm run --prefix $(CLIENT_DIR) serve