#!/bin/bash
echo "trabalhando dependencias do projeto."
go mod tidy
echo "Building da aplicação Go"
go mod download && go mod verify
go build -o app .

echo "subir postgres e pgadmin"
docker-compose up -d

sleep 15

echo "pegar ip do posgres no container"
PG_CONTAINER_ID=$(docker ps -q -f name=postgres_container_wex)
PG_CONTAINER_IP=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $PG_CONTAINER_ID)

echo "IP do container PostgreSQL: $PG_CONTAINER_IP"


# docker run -d --name wex-app-teste-container \
#   -e DB_HOST=$PG_CONTAINER_IP \
#   -e DB_PORT=5432 \
#   -e DB_NAME=wex \
#   -e DB_USER=wex \
#   -e DB_PASSWORD=wex \
#   -p 8888:8888 \
#   wex-app-teste


# docker exec -it wex-app-teste-container env
DB_HOST=$PG_CONTAINER_IP DB_PORT=5432 DB_NAME=wex DB_USER=wex DB_PASSWORD=wex ./app