# Autocomplete

A technology autocomplete application based on the most popular Stack Overflow tags.

The backend loads a snapshot of the tags into a trie, searches for terms by prefix, and returns suggestions ordered by popularity. The frontend provides debounced search, keyboard navigation, and light and dark themes.

## Technologies

- Go
- GraphQL with gqlgen
- React, TypeScript, and Vite
- Nginx
- Docker Compose

## Running the application

Prerequisite: Docker with Docker Compose.

From the project root, run:

```bash
docker compose up --build
```

Open the application at:

```text
http://localhost:8080
```

To stop the application:

```bash
docker compose down
```

## GraphQL API

The endpoint is available at `POST /query`.

Example:

```graphql
query {
  autocomplete(prefix: "java") {
    value
    score
  }
}
```

## Development

Run the backend tests and checks:

```bash
go test ./...
go vet ./...
```

Run the frontend locally:

```bash
npm --prefix ./frontend install
npm --prefix ./frontend run dev
```

In another terminal, start the backend:

```bash
go run ./cmd/server
```

Vite serves the frontend at `http://localhost:5173` and proxies `/query` to the local backend.

## Architecture

```text
React → Nginx → GraphQL → autocomplete service → trie → tag snapshot
```

The data was obtained through the Stack Exchange API. Stack Overflow user contributions are licensed under CC BY-SA.
