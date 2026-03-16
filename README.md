# GO Auth Service

Minimal Go authentication service with OTP-based registration, login, JWT auth, and profile endpoints.

## Prerequisites

- Go 1.24+
- PostgreSQL

## Environment

Set these environment variables before running:

- `PORT` (default app port)
- `DB_URL` (PostgreSQL connection string)
- `JWT_SECRET` (secret key for JWT signing)

## Run

```bash
go mod tidy
go run ./cmd
```

## Docker

```bash
docker compose up --build
```

## Main Endpoints

- `POST /auth/send-otp`
- `POST /auth/register`
- `POST /auth/login`
- `GET /profile/me` (requires Bearer token)
- `PUT /profile/me` (requires Bearer token)