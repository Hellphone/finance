run:
	docker-compose up --remove-orphans --build

migrateup:
	migrate -path db/migrations -database "postgresql://postgres:asdjk2j@finance_postgres:5432/finance?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:asdjk2j@finance_postgres:5432/finance?sslmode=disable" -verbose down

.PHONY:
	postgres migrateup migratedown