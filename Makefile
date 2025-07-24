run:
	docker compose up -d

stop:
	docker compose down

logs:
	docker compose logs -f

reload-backend:
	docker compose restart backend

reload-frontend:
	docker compose restart frontend

reload-db:
	docker compose restart cockroachdb 