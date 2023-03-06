# api-key-middleware

## Prerequisites
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### 1.Create .env file
Rename .env.example to .env and set the environment variables
```bash
mv .env.example .env
```
### 2.Build the image
```bash
docker compose build
```
### 3.Run the server
```bash
PUBLIC_PORT=3000 docker compose up
```

[//]: # (### 4. Migration)

[//]: # (```mysql )

[//]: # (```)
