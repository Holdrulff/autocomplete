# Project Tasks

This document breaks the autocomplete project into small, reviewable deliveries. Each task should be developed on a branch and submitted through a Pull Request.

Status legend:

- `COMPLETED`
- `IN PROGRESS`
- `NOT STARTED`

## Project foundation

### 0. Initialize the repository - COMPLETED

**Goal:** Create the initial Git history without publishing project files.

**Acceptance criteria:**

- The `main` branch exists locally and on GitHub.
- The local branch tracks `origin/main`.
- The original challenge document remains untracked.

### 1. Establish project documentation and governance - COMPLETED

**Goal:** Document the project plan, decision history, contribution workflow, and local-only files.

**Acceptance criteria:**

- `.gitignore` keeps `enunciado.docx` outside the public repository.
- `.gitattributes` defines portable line endings and binary files.
- `TASKS.md` decomposes the project into reviewable deliveries.
- `HISTORY.md` records relevant decisions and meaningful AI usage.
- The documentation is merged through Pull Request #1.

### 2. Split local product and cloud deployment milestones - IN PROGRESS

**Goal:** Separate product development from cloud deployment so each stage has a focused release and article.

**Acceptance criteria:**

- Milestone 1 ends with a complete Docker Compose release and local-product article.
- Milestone 2 starts only after the `v1.0.0` release and first article.
- Cloud concerns remain outside the application core through explicit interfaces.
- The decision and its consequences are recorded in `HISTORY.md`.

## Milestone 1 - Local Product

The first milestone delivers a complete application that anyone can clone and run with `docker compose up`. It does not include AWS resources or cloud deployment.

### 3. Learn Go foundations and bootstrap the module - NOT STARTED

**Goal:** Learn the Go concepts required by the project while creating the first executable package and test.

**Acceptance criteria:**

- The repository contains a Go 1.26 module with a deliberate package structure.
- Small exercises cover variables, functions, structs, slices, maps, errors, and interfaces.
- The first production package has a table-driven unit test.
- `go test ./...`, `go vet ./...`, and formatting checks pass.

### 4. Import the suggestion dataset - NOT STARTED

**Goal:** Build a reproducible Go command that imports popular technology tags from Stack Overflow.

**Acceptance criteria:**

- The command handles pagination, API quota information, and the `backoff` response field.
- Tags are normalized and duplicates are removed.
- A versioned snapshot contains up to 5,000 tags and their popularity scores.
- Dataset origin, generation date, and licensing attribution are documented.

### 5. Implement the linear search engine - NOT STARTED

**Goal:** Create a simple and correct reference implementation for prefix search.

**Acceptance criteria:**

- Queries are trimmed and compared case-insensitively while preserving punctuation.
- Queries shorter than 4 characters return no suggestions.
- Queries longer than 64 Unicode code points are rejected.
- Results are ordered by popularity and then alphabetically, with a maximum of 20 items.
- Unit tests cover validation, ranking, duplicates, Unicode, and punctuation.

### 6. Implement the trie and benchmarks - NOT STARTED

**Goal:** Build an optimized prefix index and compare it with the linear implementation.

**Acceptance criteria:**

- The trie uses Unicode code points and caches at most 20 ranked suggestions per node.
- Linear and trie implementations satisfy the same search interface.
- Equivalence tests show that both implementations return the same results.
- Benchmarks compare execution time and memory at multiple dataset sizes.
- The selected production strategy and trade-offs are recorded in `HISTORY.md`.

### 7. Expose the autocomplete HTTP service - NOT STARTED

**Goal:** Make the search engine available through a small, stateless Go HTTP service.

**Acceptance criteria:**

- `GET /v1/suggestions` implements the documented query and response contract.
- `GET /healthz` reports service readiness.
- Request validation, context cancellation, timeouts, and structured logs are implemented.
- Logs do not store the searched text.
- Handler and integration tests cover success and failure cases.

### 8. Build the GraphQL BFF - NOT STARTED

**Goal:** Provide the only API consumed by the frontend and isolate it from the autocomplete service.

**Acceptance criteria:**

- A schema-first `gqlgen` API exposes `suggestions(query: String!)`.
- The BFF calls the autocomplete service over the internal HTTP contract.
- Upstream timeouts and failures return safe GraphQL errors.
- Request context and correlation IDs are propagated.
- Resolver and contract tests cover empty, successful, and failing responses.

### 9. Learn web foundations - NOT STARTED

**Goal:** Learn the browser, HTML, CSS, JavaScript, and TypeScript concepts required before using React.

**Acceptance criteria:**

- A semantic search form is built without React first.
- Exercises cover the box model, responsive layout, DOM events, modules, and asynchronous JavaScript.
- TypeScript exercises cover primitive types, objects, unions, functions, and narrowing.
- The learning artifacts are either directly useful to the product or kept outside production code.

### 10. Create the React interface - NOT STARTED

**Goal:** Build a responsive, semantic search page with React and TypeScript.

**Acceptance criteria:**

- The page works on mobile and desktop layouts.
- Suggestions display the matching prefix in bold.
- The list shows approximately 10 rows and scrolls to reveal up to 20 results.
- Selecting a suggestion updates the search input.
- Components have focused tests for visible behavior.

### 11. Add asynchronous behavior and accessibility - NOT STARTED

**Goal:** Make the autocomplete fast, race-safe, and usable with keyboard, mouse, and touch.

**Acceptance criteria:**

- Requests start only after 4 characters and use a 150 ms debounce.
- Obsolete requests are cancelled and stale responses cannot replace newer results.
- The input follows the ARIA combobox/listbox interaction pattern.
- Arrow keys, Enter, Escape, mouse, and touch interactions work.
- Empty and recoverable error states are presented appropriately.

### 12. Learn Docker foundations - NOT STARTED

**Goal:** Understand how images, containers, layers, ports, networks, and Compose package the application.

**Acceptance criteria:**

- The differences between an image, container, Dockerfile, and Compose file can be explained.
- Each service has an intentional multi-stage Dockerfile.
- Configuration is provided through environment variables rather than image contents.
- Health checks and container dependencies are understood and documented.

### 13. Integrate the complete local application - NOT STARTED

**Goal:** Run and verify the complete system through a single local command.

**Acceptance criteria:**

- `docker compose up` starts the frontend, GraphQL BFF, and autocomplete service.
- Only the frontend/reverse proxy is exposed at `http://localhost:3000`.
- `/graphql` is proxied to the BFF, which reaches the backend on the private Compose network.
- All containers become healthy and shut down cleanly.
- Native non-Docker development remains documented and functional.

### 14. Automate quality checks - NOT STARTED

**Goal:** Validate every Pull Request and the complete browser-to-backend flow.

**Acceptance criteria:**

- CI runs Go formatting, tests, race detection, and `go vet`.
- CI runs ESLint, TypeScript checks, frontend tests, and production builds.
- GraphQL generated artifacts are checked for drift.
- Containers build and Docker Compose end-to-end tests pass.
- Load and Lighthouse results satisfy the documented local thresholds.

### 15. Publish the local release and first article - NOT STARTED

**Goal:** Release a reproducible portfolio project and explain its complete local architecture in English.

**Acceptance criteria:**

- README includes requirements, setup, architecture, testing, screenshots or a GIF, and troubleshooting.
- A fresh clone works with `docker compose up` using only documented prerequisites.
- GitHub release `v1.0.0` is published.
- The article explains the algorithm, APIs, frontend, Docker, CI, benchmarks, and trade-offs.
- AWS is clearly identified as a future milestone rather than a missing v1 feature.

## Milestone 2 - Cloud Deployment

This milestone starts only after Milestone 1 and its article are complete.

### 16. Adapt the services for AWS - NOT STARTED

**Goal:** Add AWS transport adapters without changing the application core or GraphQL contract.

**Acceptance criteria:**

- Go Lambda adapters reuse the existing search and BFF application layers.
- AWS SAM defines ARM64 functions, API Gateway, private S3, and CloudFront.
- The GraphQL Lambda has least-privilege permission to invoke the private autocomplete Lambda.
- The Docker Compose workflow remains functional.

### 17. Automate cloud delivery and operations - NOT STARTED

**Goal:** Deploy safely and observe the application while keeping monthly cost below USD 5.

**Acceptance criteria:**

- GitHub Actions deploys through AWS OIDC without long-lived credentials.
- Logs, retention, concurrency, throttling, and budget alerts are configured.
- Rollback and smoke-test procedures are documented.
- The application is available through a public HTTPS CloudFront URL.

### 18. Publish the cloud release and second article - NOT STARTED

**Goal:** Explain how and why the local architecture evolved for AWS.

**Acceptance criteria:**

- Warm and cold latency, error rate, and estimated cost are recorded.
- Local and cloud data flows are compared with diagrams and measurements.
- GitHub release `v2.0.0` is published.
- The second article covers serverless trade-offs, IAM, observability, deployment, and cost.
