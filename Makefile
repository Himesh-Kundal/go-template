sqlc:
	sqlc generate

up:
	docker compose up --build

down:
	docker compose down -v

test:
	DATABASE_URL=postgres://app:secret@localhost:5432/app?sslmode=disable go test -v -count=1 ./tests/...
