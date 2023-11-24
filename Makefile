ifeq ($(POSTGRES_SETUP_TEST),)
	POSTGRES_SETUP_TEST := user=postgres password=postgres dbname=OzonBaze host=localhost port=5432 sslmode=disable
endif

MIGRATION_FOLDER=$(CURDIR)/migrations
INTEGRATION_TEST_PATH?=./test
UNIT_TEST_PATH_1 = ./internal/serv/server/
UNIT_TEST_PATH_2 = ./internal/serv/repository/postgres/


docker_start_components:
	docker-compose up;

build:
	docker-compose build

up-all:
	docker-compose up -d zookeeper kafka1 kafka2 kafka3

docker_stop:
	docker-compose down;

.PHONY: migration-create
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

.PHONY: test-migration-up
test-migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" up

.PHONY: test-migration-down
test-migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" down

.PHONY: unit-test
unit-test:
	go test $(UNIT_TEST_PATH_1)
	go test $(UNIT_TEST_PATH_2)

.PHONY: integration-test
integration-test:
	go test -tags=integration $(INTEGRATION_TEST_PATH)

.PHONY: truncate
truncate:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" down
