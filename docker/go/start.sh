#!/bin/bash

set -e

until PGPASSWORD=$POSTGRES_ROOT_PASSWORD psql -h $POSTGRES_HOST -U $POSTGRES_ROOT_USER -d $POSTGRES_DB -c '\l'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 3
done

>&2 echo "Postgres is up - executing command"
cd /go/src/app
go get
go build -o main .
go run main.go handler.go user.go
