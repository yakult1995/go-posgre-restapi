#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" -d "$POSTGRES_DB"<<-EOSQL
     CREATE TABLE $POSTGRES_TABLE_NAME (
        id SERIAL NOT NULL PRIMARY KEY,
        name TEXT NOT NULL,
        email TEXT NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp,
        updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp
     )
EOSQL
