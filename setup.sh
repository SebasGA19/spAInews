#!/usr/bin/env bash

mkdir -p /opt/spAInews

docker-compose -f ./docker-compose.yaml down --remove-orphans
docker-compose -f ./docker-compose.yaml up -d --build
chmod 777 spainews.sh
cp spainews.sh /etc/cron.daily/spainews