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

## Add a new feature
1. Create a new branch
```bash
git checkout -b ft/dbdocs
```
2. Add and commit all changes
3. Push to new branch
```bash
git push origin ft/dbdocs
```
<<<<<<< HEAD
=======
4. Go to github merge request > Create pull request > Choose Squash and merge > Confirm squash merge > Delete branch
>>>>>>> 4c4e6fe (Define gRPC API and generate Go code with protobuf)


## Create a new query flow

1. Create a new query in db/query/user.sql
2. Run "make sqlc"
3. Buat test db/sqlc/user_test.go
4. Run "make mock"
   - Update Querier interface

## Generate DB Documentation
1. sudo apt install nodejs
2. npm i -g dbdocs
3. dbdocs login
4. copy sql query from dbdiagram.io 
5. create doc/db.dbml then paste
7. dbdocs password --set dharma --project simple_bank
6. dbdocs build doc/db.dbml
8. visit dbdocs.io/... link
9. Convert dbml to sql
```bash
dbml2sql --postgres -o  doc/schema.sql doc/db.dbml
```

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

### GRPC

## OS Setup
1. Go to [here](https://grpc.io/docs/protoc-installation/) and follow the instruction to instal protoc
2. Click [here](https://grpc.io/docs/languages/go/quickstart/) and setup for Prerequisites
3. Install vscode-proto3 extension in vscode
4. In vscode > settings.json add this line
```bash
"protoc": {
      "options": [
          "--proto_path=protos",
      ]
   }
``` 
5. Click [here](https://stackoverflow.com/a/57730314) do the same but change .bash_profile to .bashrc

## gRPC Test Setup
1. Download file [here](https://github.com/ktr0731/evans/releases)
2. Extract and copy content to root project
3. Run chmod +x ./evans
4. Change script from evans to ./evans

## Create Proto Function
1. Make file rpc_create_user.proto
2. Create new file user.proto to define interface/class *optional
3. Create new file rpc_create_user.proto to define request and response
4. Import file and register function in service_simple_bank.proto
5. Run make proto