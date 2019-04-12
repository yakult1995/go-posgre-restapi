#!/bin/sh

mv /tmp/go.env /home/ec2-user/go-posgre-restapi/docker/go/.env
mv /tmp/pgweb.env /home/ec2-user/go-posgre-restapi/docker/pgweb/.env
mv /tmp/posgre.env /home/ec2-user/go-posgre-restapi/docker/posgre/.env

docker-compose -f /home/ec2-user/go-posgre-restapi/docker-compose.yml build
docker-compose -f /home/ec2-user/go-posgre-restapi/docker-compose.yml up -d
