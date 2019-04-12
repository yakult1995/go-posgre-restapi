#!/bin/sh

cd /home/ec2-user/go-posgre-restapi > /dev/null 2>&1

docker-compose down > /dev/null 2>&1

mv ./docker/go/.env /tmp/go.env > /dev/null 2>&1
mv ./docker/pgweb/.env /tmp/pgweb.env > /dev/null 2>&1
mv ./docker/posgre/.env /tmp/posgre.env > /dev/null 2>&1

echo printenv
