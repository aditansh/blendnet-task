## Tech Stack

- [Golang](https://go.dev): An open-source programming language supported by Google

  - [Fiber](https://gofiber.io): A fast and low memory footprint web framework for Go.
  - [Gorm](https://gorm.io): The fantastic ORM library for Golang, aims to be developer-friendly.
  - [go-redis](https://github.com/redis/go-redis): Go redis client.

- [Docker](https://www.docker.com): Docker is a platform for developing, shipping, and running applications using containers.
- [PostgreSQL](https://www.postgresql.org/): The worlds most advanced Open Source Relational Database.
- [Redis](https://redis.io): The world’s fastest in-memory database from the ones who built it.

# How To Run

## Prerequisites:

- [Docker](https://www.docker.com): Docker is a platform for developing, shipping, and running applications using containers.
- [Postman](https://www.postman.com)/[Apidog](https://apidog.com): A tool to test backend APIs without having to write frontends.

## Steps

1. Create a `.env` file in the root directory and add the following environment variables:

```bash
POSTGRES_HOST=postgres
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=admin
POSTGRES_DB=blendnet

CLIENT_ORIGIN=http://localhost:3000
PORT=:8080

ACCESS_TOKEN_SECRET=your-access-token-secret-key
REFRESH_TOKEN_SECRET=your-refresh-token-secret-key
ACCESS_TOKEN_EXPIRY=24h
REFRESH_TOKEN_EXPIRY=720h

API_KEY=your-alphavantage-api-key

REDIS_DB_ADDR=localhost:6379
REDIS_DB_PASSWORD=
```

2. Run the following command to start the backend:

```bash
docker-compose up --build -d
```

3.  Use postman and test the api at the endpoint `http://localhost/api`
