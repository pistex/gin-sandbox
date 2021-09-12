e2e:
	docker compose up test-db-migration && docker compose up --build e2e 
