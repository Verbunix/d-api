docker network create public

docker-compose kill
docker-compose rm -f
docker-compose pull
docker-compose up -d
