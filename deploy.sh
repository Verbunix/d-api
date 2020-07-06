#!/usr/bin/env sh

set -e

ENV=${ENV:-local}
echo "ENV: ${ENV}"

docker network create public
cd docker

if [ "$BUILD" = 1 ]
then
    docker-compose kill
    docker-compose rm -f
    docker-compose pull
    docker-compose up --build -d
else
    docker-compose kill
    docker-compose rm -f
    docker-compose pull
    docker-compose up -d
fi
