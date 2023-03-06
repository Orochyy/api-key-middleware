# api-key-middleware

## Prerequisites
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### 1.Create .env file
Rename .env.example to .env and set the environment variables
```bash
cp .env.example .env
```
### 2.Build the image
```bash
docker compose build
```
### 3.Run the server
```bash
PUBLIC_PORT=3000 docker compose up
```

### 4. Migrate the database
## 4.1 Open docker container bash
```bash
docker compose run --rm go_key_api bash
```

## 4.2 Run the migrations
```bash
migrate -source file://database/migrations -database "$DB_DRIVER://$DB_USERNAME:$DB_PASSWORD@tcp($DB_HOST:$DB_PORT)/$MYSQL_DATABASE" up
```
