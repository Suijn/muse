.PHONY: build-server run-server

build-server:
	docker compose build --no-cache server

run-server:
	docker compose up server 
