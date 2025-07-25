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

rebuild-frontend:
	docker compose build frontend

rebuild-backend:
	docker compose build backend

rebuild:
	docker compose build
	docker compose up -d