#!make
.PHONY: run build-client migrate migrate-down build-prod
include .env

CLIENT_DIR = ./client
CLIENT_BUILD_DIR = ./client/dist

BUILD_DATE = $(shell date -u)

MYSQL_USER = ${DOCK_MYSQL_USER}
MYSQL_PASS = ${DOCK_MYSQL_PASS}
MYSQL_DB   = ${DOCK_MYSQL_DB}
MYSQL_PORT = ${DOCK_MYSQL_PORT}

CURRENT_TAG=$(shell git describe --tags)

run:
	docker-compose up -d
ifeq ($(wildcard $(CLIENT_BUILD_DIR)/.*),)
	npm run --prefix $(CLIENT_DIR) build
endif

build-client:
	npm i --prefix $(CLIENT_DIR)
	npm run --prefix $(CLIENT_DIR) build

run-dev-client:
	npm run --prefix $(CLIENT_DIR) serve

seed:
	docker exec -i $(DOCK_MYSQL_CONTAINER_NAME) mysql -u $(MYSQL_USER) -p$(MYSQL_PASS) $(DOCK_MYSQL_DB) < $(shell pwd)/db/$(seed_folder)/user_seeds.sql

build-prod:
	docker build -f $(shell pwd)/build/Dockerfile.build --build-arg GIT_COMMIT=$(CURRENT_TAG) --build-arg BUILD_DATE="$(shell date -u)" --build-arg BUILD_ARCH=amd64 -t assets-proxy:$(CURRENT_TAG) .
