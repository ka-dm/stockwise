reload-backend:
	docker compose restart backend

reload-frontend:
	docker compose restart frontend

reload-db:
	docker compose restart cockroachdb 