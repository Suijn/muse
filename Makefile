.PHONY: build-server build-tests run-server run-tests

build-api:
	docker compose build --no-cache api

build-tests:
	docker compose build --no-cache tests

run-api:
	docker compose up api

run-tests:
	docker compose up tests
