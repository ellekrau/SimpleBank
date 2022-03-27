# Migration
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate create -ext sql -dir src\database\migrations -seq init_schema

migrate -path src/database/migrations --database "postgresql://admin:P@ssw0rd@localhost:5432/simple_bank?sslmode=disable" --verbose up
migrate -path src/database/migrations --database "postgresql://admin:P@ssw0rd@localhost:5432/simple_bank?sslmode=disable" --verbose down