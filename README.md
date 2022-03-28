# Course
Backend Master Class [Golang + PostgreSQL + Kubernetes]

https://www.udemy.com/course/backend-master-class-golang-postgresql-kubernetes/

## .env
Create an .env file and set DB configuration  
`DB_NAME= `  
`DB_PASSWORD= `  
`DB_USER= `

## Migrations
Install  
 `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`

Create first migration  
`migrate create -ext sql -dir src\database\migrations -seq init_schema`

Up  
`migrate -path src/database/migrations --database "postgresql://admin:P@ssw0rd@localhost:5432/simple_bank?sslmode=disable" --verbose up`

Down  
`migrate -path src/database/migrations --database "postgresql://admin:P@ssw0rd@localhost:5432/simple_bank?sslmode=disable" --verbose down`

## Sqlc
Install  
`go install github.com/kyleconroy/sqlc/cmd/sqlc@latest`

Postgres documentation  
https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html

To run in windows using WSL:
- Download the binary: https://github.com/kyleconroy/sqlc/releases/download/v1.12.0/sqlc_1.12.0_linux_amd64.tar.gz
- Execute using WLS with this binary as executable  

`root@LAPTOP-PFVT9C7K:/mnt/c/Git/sqlc_1.12.0_linux_amd64# ./sqlc generate -f "../study/simple-bank/sqlc.yaml"`

## External packages
- https://github.com/lib/pq
- https://github.com/stretchr/testify
