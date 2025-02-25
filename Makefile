# Makefile for car-pooling-challenge
# vim: set ft=make ts=8 noet
# Copyright Cabify.com
# Licence MIT

# Variables
# UNAME		:= $(shell uname -s)

.EXPORT_ALL_VARIABLES:

# this is godly
# https://news.ycombinator.com/item?id=11939200
.PHONY: help
help:	### this screen. Keep it first target to be default
ifeq ($(UNAME), Linux)
	@grep -P '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
else
	@# this is not tested, but prepared in advance for you, Mac drivers
	@awk -F ':.*###' '$$0 ~ FS {printf "%15s%s\n", $$1 ":", $$2}' \
		$(MAKEFILE_LIST) | grep -v '@awk' | sort
endif

# Targets
#
.PHONY: debug
debug:	### Debug Makefile itself
	@echo $(UNAME)

.PHONY: build
build:
	CGO_ENABLED=0 go build -a -o target/bin/carpool ./cmd/car-pooling-server/main.go

.PHONY: run
run: build
	target/bin/carpool

.PHONY: dockerize
docker: build
	docker build -t car-pooling-challenge:latest .

.PHONY: test.acceptance
test.acceptance: docker
	CABIFY_CHALLENGE_TESTCASE=acceptance docker-compose up --abort-on-container-exit --always-recreate-deps --force-recreate

.PHONY: test
test: test.acceptance
