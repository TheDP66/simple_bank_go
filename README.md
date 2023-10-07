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

Note:

- if you get this error while migrating

```bash
error: Dirty database version 2. Fix and force version.
```

Open database (dbeaver) > field schema_migrations > uncheck the dirty field

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

## Create a new query flow

1. Create a new query in db/query/user.sql
2. Run "make sqlc"
3. Buat test db/sqlc/user_test.go
4. Run "make mock"
   - Update Querier interface

## Dockerize application

# Dockerfile

1. Create Dockerfile
2. Run Dockerfile to create images

```bash
docker build -t simplebank:latest .
```

3. Create a new network in docker

```bash
docker network create bank-network
```

4. Connect container to network (so postgres container can communicate with simplebank container)

```bash
docker network connect bank-network postgres3.17
```

5. Run created images

- For production use

```bash
docker run -t --name simplebank_prod --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:admin@postgres3.17:5432/simple_bank?sslmode=disable" simplebank:latest
```

# docker-compose.yaml

1. Create file docker-compose.yaml
2. Run using this command

```bash
docker compose up
```

- note: docker-compose.yaml run in group not like docker file
