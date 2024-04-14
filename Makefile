PROJECT_DIR = $(CURDIR)
PROJECT_BIN = $(PROJECT_DIR)/bin
$(shell [ -f bin ] || mkdir -p $(PROJECT_BIN))
PATH := $(PROJECT_BIN):$(PATH)

GOLANGCI_LINT = $(PROJECT_BIN)/golangci-lint

.PHONY: .install-linter
.install-linter:
	### INSTALL GOLANGCI-LINT ###
	[ -f $(PROJECT_BIN)/golangci-lint ] || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(PROJECT_BIN) v1.46.2

.PHONY: lint
lint: .install-linter
	### RUN GOLANGCI-LINT ###
	$(GOLANGCI_LINT) run ./... --config=./.golangci.yml

.PHONY: lint-fast
lint-fast: .install-linter
	$(GOLANGCI_LINT) run ./... --fast --config=./.golangci.yml




.PHONY: build up integration_tests fake_banners

#up - stop - start because volume and tables in the database are created at the first startup,
#so the database may not be ready to accept the connection and the application will fall down :(
first-run: build up stop start integration_tests

test: up integration_tests

rebuild: build up

build:
	docker-compose  build

up:
	docker-compose up -d

start:
	docker-compose start

integration_tests:
	go test ./tests/integration/...

fake_banners:
	go test -run TestCreate_Happy ./tests/integration/admincreate_test.go

stop:
	docker-compose stop

down:
	docker-compose down


