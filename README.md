# hangout-planner

A simple and open source hangout planner.

**Development**

- Backend (Go):
	- Default port: `:8080` (overrides via `PORT` env)
	- Entrypoint: [backend/cmd/api/main.go](backend/cmd/api/main.go)
	- Health check: `GET /api/health`

- Frontend (React + Vite + TypeScript):
	- Location: [client/](client)
	- Dev server: `http://localhost:5173` (Vite)
	- Dev proxy: `/api` → `http://localhost:8080`

## Run locally

In separate terminals:

1) Backend (dev)

```bash
cd backend
go run ./cmd/api
```

2) Frontend (dev)

```bash
cd client
yarn dev
# or: npm run dev / pnpm dev
```

Open the frontend at `http://localhost:5173`. The home page calls `/api/health` and displays the API status.

## Single-origin (serve UI from Go)

Build the client and copy static assets into the backend. Then start only the Go server and visit `http://localhost:8080`.

```bash
# From repo root
cd client
yarn build:copy

cd ../backend
go run ./cmd/api
# Open http://localhost:8080  (UI)
# API stays at http://localhost:8080/api/...
```

The Go server serves files from [backend/public](backend/public). SPA routes fall back to `index.html`.

## Notes

- The Vite dev server is configured to proxy API calls to the Go server, avoiding CORS during local development. See vite.config.ts.
- To change the backend port, set `PORT` before starting the Go server. Update the Vite proxy if you change host/port.
