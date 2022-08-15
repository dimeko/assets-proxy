#!make
.PHONY: run build-client migrate migrate-down
include .env

CLIENT_DIR = ./client
CLIENT_BUILD_DIR = ./client/dist
DB_DRIVER=mysql

MYSQL_USER = ${DOCK_MYSQL_USER}
MYSQL_PASS = ${DOCK_MYSQL_PASS}
MYSQL_DB   = ${DOCK_MYSQL_DB}
MYSQL_HOST = localhost
MYSQL_PORT = ${DOCK_MYSQL_PORT}

run:
	docker-compose up -d
ifeq ($(wildcard $(CLIENT_BUILD_DIR)/.*),)
	npm run --prefix $(CLIENT_DIR) build
endif

build-client:
	npm run --prefix $(CLIENT_DIR) build

run-dev-client:
	npm run --prefix $(CLIENT_DIR) serve

migrate:
	migrate -path db/migrations -database "$(DB_DRIVER)://$(MYSQL_USER):$(MYSQL_PASS)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DB)" -verbose up

migrate-down:
	migrate -path db/migrations -database "$(DB_DRIVER)://$(MYSQL_USER):$(MYSQL_PASS)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DB)" -verbose down
