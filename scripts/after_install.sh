#!/bin/sh

cd /home/ec2-user/go-posgre-restapi

mv /tmp/go.env ./docker/go/.env
mv /tmp/pgweb.env ./docker/pgweb/.env
mv /tmp/posgre.env ./docker/posgre/.env

docker-compose build
docker-compose up -d
