run:
	docker-compose up --remove-orphans --build

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=asdjk2j -d postgres:12-alpine

migrateup:
	migrate -path migrations -database "postgresql://postgres:asdjk2j@finance_postgres:5432/finance?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations -database "postgresql://postgres:asdjk2j@finance_postgres:5432/finance?sslmode=disable" -verbose down

.PHONY:
	postgres migrateup migratedown sqlc