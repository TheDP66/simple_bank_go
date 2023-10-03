## Initialize

Run this command sequentially from root of this project

```bash
make postgres
make createdb
```

## Getting started

Follow this step to run this program

1. Run Docker Desktop
2. Run PostgresQL on docker. Run this command if you have any problem

```bash
sudo fuser -k 5432/tcp
```

3. Run dbeaver-ce
4. Start a connection to postgres - localhost:5432
5. Open Databases
6. Find simple_bank

## Design and generate database

Follow this step to design and generate database

1. Find file on /DB/Simple bank - dbdiagram.io, copy content
2. Open https://dbdiagram.io/ and paste content to left panel
3. After done, find Export button on upper bar then press Export to PostgreSQL

## Golang-migrate

Create a migration

```bash
migrate create -ext sql -dir db/migration -seq init_schema
```

Run migration

```bash
migrate -path db/migration -database "postgresql://root:admin@localhost:5432/simple_bank?sslmode=disable" -verbose up
```

## Generate query

Edit file on directory db/query. Follow the query formatting before creating a new query function. After done creating a new query you can run

```bash
sqlc generate
```

A new/updated file will be generated in folder db/sqlc

## Generate Mock

To generate a mock implementation run this command

```bash
mockgen -package mockdb -destination db/mock/store.go github.com/TheDP66/simple_bank_go/db/sqlc Store
```

If you got error run this command

```bash
go get github.com/golang/mock/
```
