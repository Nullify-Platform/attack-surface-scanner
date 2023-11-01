# Test Server

This project contains a REST API with 2 endpoints.

- POST /api/users
- GET /api/users/{username}

The POST endpoint has no authentication.
The GET endpoint checks for an Authorization header

## Usage

```sh
go run ./cmd/server/...
```

or

```sh
docker-compose up
```
