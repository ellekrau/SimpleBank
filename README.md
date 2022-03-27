# Course
Backend Master Class [Golang + PostgreSQL + Kubernetes]

https://www.udemy.com/course/backend-master-class-golang-postgresql-kubernetes/

## Migrations
### Install package
`go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`

### Create first migration
`migrate create -ext sql -dir src\database\migrations -seq init_schema`

### Up
`migrate -path src/database/migrations --database "postgresql://admin:P@ssw0rd@localhost:5432/simple_bank?sslmode=disable" --verbose up`

### Down
`migrate -path src/database/migrations --database "postgresql://admin:P@ssw0rd@localhost:5432/simple_bank?sslmode=disable" --verbose down`
