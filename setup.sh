#!/usr/bin/env bash

mkdir -p /opt/spAInews

docker-compose -f ./docker-compose.yaml down --remove-orphans
docker-compose -f ./docker-compose.yaml up -d --build
chmod 777 spainews.sh
cp spainews.sh /opt/spAInews
cp spainews.service /etc/systemd/system
systemctl daemon-reload
systemctl enable spainews
systemctl start spainews