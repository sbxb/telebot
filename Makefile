ifneq (,$(wildcard ./.env))
	include .env
	export
endif

.PHONY: run
run:
	go run .
