build:
	docker compose build 

dev:
	docker compose up

e2e:
	docker compose up test-db-migration && docker compose up --build e2e 
