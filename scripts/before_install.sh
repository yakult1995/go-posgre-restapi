#!/bin/sh

tree

docker-compose -f /home/ec2-user/go-posgre-restapi/docker-compose.yml down > /dev/null 2>&1

mv /home/ec2-user/go-posgre-restapi/docker/go/.env /tmp/go.env > /dev/null 2>&1
mv /home/ec2-user/go-posgre-restapi/docker/pgweb/.env /tmp/pgweb.env > /dev/null 2>&1
mv /home/ec2-user/go-posgre-restapi/docker/posgre/.env /tmp/posgre.env > /dev/null 2>&1

