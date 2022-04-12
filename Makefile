GO              ?= go
GOFLAGS         :=
REPO_PATH       := github.com/jchauncey/allaclone
GIT_VERSION     := $(shell git describe --tags --dirty)
GIT_COMMIT      := $(shell git rev-parse HEAD)
BUILD_DATE      := $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
LDFLAGS         := -s -X $(REPO_PATH)/cmd/version.GitVersion=$(GIT_VERSION) -X $(REPO_PATH)/cmd/version.GitCommit=$(GIT_COMMIT) -X $(REPO_PATH)/cmd/version.BuildDate=$(BUILD_DATE)
BINDIR          := $(CURDIR)/bin
CLI_NAME        ?= allaclone


.PHONY: clean
clean:
	rm -rf bin/

all: clean bin/$(CLI_NAME)

bin/$(CLI_NAME):
	CGO_ENABLED=0 $(GO) build -ldflags '$(LDFLAGS)' -o $(BINDIR)/$(CLI_NAME)

.PHONY: clean-db
clean-db:
	docker rm -f database

.PHONY: db
db:
	docker build -t alla-db -f Dockerfile.db .

.PHONY: run-db
run-db: clean
	docker run -p 3306:3306 --name database -e MARIADB_ROOT_PASSWORD=password -d alla-db --port 3306
	sleep 120

.PHONY: lint
lint:
	golangci-lint -v run ./... --timeout 5m

.PHONY: install-deps
install-deps:
	go install github.com/vektra/mockery/v2@v2.9.4
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.42.1
	go install github.com/jstemmer/go-junit-report@v0.9.1
	go install github.com/t-yuki/gocover-cobertura@master
	go install golang.org/x/tools/...@latest