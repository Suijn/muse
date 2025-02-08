.PHONY: build-server build-tests run-server run-tests

build-server:
	docker compose build --no-cache server

build-tests:
	docker compose build --no-cache tests

run-server:
	docker compose up server 

run-tests:
	docker compose up tests
