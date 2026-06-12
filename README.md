# Monorepo Template

A monorepo skeleton combining a Go backend and a SvelteKit frontend.

## Overview

This project uses:

- Go + Gin for the backend API and production static file serving
- SvelteKit for the frontend application
- Separate development servers for frontend and backend
- A single production entrypoint where the Go server serves the built frontend for `/` and exposes API routes under `/api/`

## Project structure

```text
.
├── .env                  # Root environment variables for the Go server
├── .env.example          # Example environment file
├── go.mod                # Go module definition
├── server/               # Go backend source
│   ├── config/           # Environment/config loading
│   ├── handler/          # HTTP handlers
│   ├── router/           # Gin router setup
│   └── main.go           # Entry point
└── web/                  # SvelteKit frontend source
    ├── src/              # Svelte app source
    ├── package.json      # Frontend scripts and dependencies
    └── build/            # Production frontend build output
```

## Requirements

- Go 1.22+
- Node.js 20+
- npm

## Environment configuration

Create a root-level `.env` file based on `.env.example`:

```env
ENV=dev
PORT=8080
FRONTEND_DIR=./web/build
```

### Variables

- `ENV`: set to `dev` for development or `prod` for production behavior
- `PORT`: the port the Go backend listens on
- `FRONTEND_DIR`: the directory containing the built frontend files to serve in production

> In development, the SvelteKit app runs separately and does not need the backend to serve its assets.

## Development workflow

Development uses two separate processes:

1. Frontend dev server (SvelteKit)
2. Backend dev server (Go)

### 1) Start the frontend

```bash
cd web
npm install
npm run dev
```

By default, the SvelteKit dev server runs on:

- http://localhost:5173

### 2) Start the backend

In a separate terminal:

```bash
go run ./server
```

By default, the Go server runs on:

- http://localhost:8080

### Development behavior

- The frontend is served by Vite on its own port
- The backend is served by Gin on its own port
- API requests should target the backend server, for example:
  - http://localhost:8080/api/health

## Production workflow

### 1) Build the frontend

```bash
cd web
npm install
npm run build
```

This generates the static frontend output in:

- `web/build`

### 2) Configure production environment

Set the root `.env` file for production, for example:

```env
ENV=prod
PORT=5005
FRONTEND_DIR=./web/build
```

### 3) Start the backend

```bash
go run ./server
```

The Go server will now:

- serve the built frontend at `/`
- expose API routes under `/api/`
- return the frontend app for client-side routes such as `/some/page`

### Production endpoints

- Frontend app: http://localhost:5005/
- API health: http://localhost:5005/api/health

## API routes

The backend currently exposes:

- `GET /api/health` → returns a health check JSON response

You can extend this by adding new handlers in `server/handler/` and registering them in `server/router/router.go`.

## Useful commands

### Install dependencies

```bash
cd web && npm install
```

```bash
go mod download
```

### Run tests

```bash
go test ./...
```

### Frontend checks

```bash
cd web
npm run check
npm run build
```

## Notes

- In development, frontend and backend are intentionally separate.
- In production, the Go server acts as the single entrypoint and serves the built frontend for browser navigation.
- If you change the frontend build output location, update `FRONTEND_DIR` in `.env` accordingly.
