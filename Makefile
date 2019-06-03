# Copyright (C) 2019 Nicolas Lamirault <nicolas.lamirault@gmail.com>

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

APP = bbox_exporter

VERSION=$(shell \
        grep "const Version" version/version.go \
        |awk -F'=' '{print $$2}' \
        |sed -e "s/[^0-9.]//g" \
	|sed -e "s/ //g")

SHELL = /bin/bash

DIR = $(shell pwd)

DOCKER = docker

GO = go

# GOX = gox -os="linux darwin windows freebsd openbsd netbsd"
GOX = gox -osarch="linux/amd64" -osarch="linux/arm64" -osarch="linux/arm" -osarch="darwin/amd64" -osarch="windows/amd64"
GOX_ARGS = "-output={{.Dir}}-$(VERSION)_{{.OS}}_{{.Arch}}"

EXE = $(shell ls bbox_exporter-*)
BINTRAY_URI = https://api.bintray.com
BINTRAY_USERNAME = nlamirault
BINTRAY_REPOSITORY= oss

NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

MAKE_COLOR=\033[33;01m%-20s\033[0m

MAIN = github.com/nlamirault/bbox_exporter

PACKAGE=$(APP)-$(VERSION)
ARCHIVE=$(PACKAGE).tar

.DEFAULT_GOAL := help

.PHONY: help
help:
	@echo -e "$(OK_COLOR)==== $(APP) [$(VERSION)] ====$(NO_COLOR)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "$(MAKE_COLOR) : %s\n", $$1, $$2}'

clean: ## Cleanup
	@echo -e "$(OK_COLOR)[$(APP)] Cleanup$(NO_COLOR)"
	@rm -fr $(EXE) $(APP) $(APP)-*.tar.gz

.PHONY: build
build: ## Make binary
	@echo -e "$(OK_COLOR)[$(APP)] Build $(NO_COLOR)"
	@$(GO) build .

.PHONY: test
test: ## Launch unit tests
	@echo -e "$(OK_COLOR)[$(APP)] Launch unit tests $(NO_COLOR)"
	@$(GO) test

.PHONY: run
run: ## Start exporter
	@echo -e "$(OK_COLOR)[$(APP)] Start the exporter $(NO_COLOR)"
	@./$(APP) -log.level DEBUG

.PHONY: lint
lint: ## Launch golint
	@$(foreach file,$(SRCS),golint $(file) || exit;)

.PHONY: vet
vet: ## Launch go vet
	@$(foreach file,$(SRCS),$(GO) vet $(file) || exit;)

.PHONY: errcheck
errcheck: ## Launch go errcheck
	@echo -e "$(OK_COLOR)[$(APP)] Go Errcheck $(NO_COLOR)"
	@$(foreach pkg,$(PKGS),errcheck $(pkg) || exit;)

.PHONY: coverage
coverage: ## Launch code coverage
	@$(foreach pkg,$(PKGS),$(GO) test -cover $(pkg) || exit;)

docker-build: ## Build Docker image
	@echo -e "$(OK_COLOR)Docker build $(APP):$(VERSION)$(NO_COLOR)"
	@docker build -t $(APP):$(VERSION) .

docker-run: ## Run the Docker image
	@echo -e "$(OK_COLOR)Docker run $(APP):$(VERSION)$(NO_COLOR)"
	@docker run --rm=true \
		$(APP):$(VERSION) \
		-log.level debug -bbox https://mabbox.bytel.fr

gox: ## Make all binaries
	@echo -e "$(OK_COLOR)[$(APP)] Create binaries $(NO_COLOR)"
	$(GOX) $(GOX_ARGS) github.com/nlamirault/bbox_exporter

.PHONY: binaries
binaries: ## Upload all binaries
	@echo -e "$(OK_COLOR)[$(APP)] Upload binaries to Bintray $(NO_COLOR)"
	for i in $(EXE); do \
		curl -T $$i \
			-u$(BINTRAY_USERNAME):$(BINTRAY_APIKEY) \
			"$(BINTRAY_URI)/content/$(BINTRAY_USERNAME)/$(BINTRAY_REPOSITORY)/$(APP)/${VERSION}/$$i;publish=1"; \
        done

# for goprojectile
.PHONY: gopath
gopath:
	@echo `pwd`:`pwd`/vendor