# Project Tasks

This document breaks the autocomplete project into small, reviewable deliveries. Each task should be developed on a branch and submitted through a Pull Request.

Status legend:

- `COMPLETED`
- `IN PROGRESS`
- `NOT STARTED`

## 0. Initialize the repository - COMPLETED

**Goal:** Create the initial Git history without publishing project files.

**Acceptance criteria:**

- The `main` branch exists locally and on GitHub.
- The local branch tracks `origin/main`.
- The original challenge document remains untracked.

## 1. Establish project documentation and governance - IN PROGRESS

**Goal:** Document the project plan, decision history, contribution workflow, and local-only files.

**Acceptance criteria:**

- `.gitignore` keeps `enunciado.docx` outside the public repository.
- `TASKS.md` decomposes the project into reviewable deliveries.
- `HISTORY.md` records relevant decisions and meaningful AI usage.
- The first documentation changes are merged through a Pull Request.

## 2. Import the suggestion dataset - NOT STARTED

**Goal:** Build a reproducible Go command that imports popular technology tags from Stack Overflow.

**Acceptance criteria:**

- The command handles pagination, API quota information, and the `backoff` response field.
- Tags are normalized and duplicates are removed.
- A versioned snapshot contains up to 5,000 tags and their popularity scores.
- Dataset origin, generation date, and licensing attribution are documented.

## 3. Implement the linear search engine - NOT STARTED

**Goal:** Create a simple and correct reference implementation for prefix search.

**Acceptance criteria:**

- Queries are trimmed and compared case-insensitively while preserving punctuation.
- Queries shorter than 4 characters return no suggestions.
- Queries longer than 64 Unicode code points are rejected.
- Results are ordered by popularity and then alphabetically, with a maximum of 20 items.
- Unit tests cover validation, ranking, duplicates, Unicode, and punctuation.

## 4. Implement the trie and benchmarks - NOT STARTED

**Goal:** Build an optimized prefix index and compare it with the linear implementation.

**Acceptance criteria:**

- The trie is built using Unicode code points.
- Each node caches at most the 20 highest-ranked suggestions for its prefix.
- Equivalence tests show that linear and trie searches return the same results.
- Benchmarks compare execution time and memory usage at multiple dataset sizes.
- The trade-offs and selected production strategy are recorded in `HISTORY.md`.

## 5. Expose the autocomplete HTTP service - NOT STARTED

**Goal:** Make the search engine available through a small Go HTTP service.

**Acceptance criteria:**

- `GET /v1/suggestions` implements the documented query and response contract.
- `GET /healthz` reports service readiness.
- Request validation, context cancellation, and timeouts are implemented.
- Logs are structured and do not store the searched text.
- Handler and integration tests cover success and failure cases.

## 6. Build the GraphQL BFF - NOT STARTED

**Goal:** Provide the only public API consumed by the frontend and isolate it from the autocomplete service.

**Acceptance criteria:**

- A schema-first `gqlgen` API exposes `suggestions(query: String!)`.
- The local adapter calls the autocomplete service over HTTP.
- Upstream timeouts and failures return safe GraphQL errors.
- Request context and correlation IDs are propagated.
- Resolver and contract tests cover empty, successful, and failing responses.

## 7. Create the React interface - NOT STARTED

**Goal:** Build a responsive, semantic search page with React and TypeScript.

**Acceptance criteria:**

- The page works on mobile and desktop layouts.
- Suggestions display the matching prefix in bold.
- The list displays approximately 10 rows and scrolls to reveal up to 20 results.
- Selecting a suggestion updates the search input.
- Components have focused tests for their visible behavior.

## 8. Add asynchronous behavior and accessibility - NOT STARTED

**Goal:** Make the autocomplete fast, race-safe, and usable with keyboard, mouse, and touch.

**Acceptance criteria:**

- Requests start only after 4 characters and use a 150 ms debounce.
- Obsolete requests are cancelled and stale responses cannot replace newer results.
- The input follows the ARIA combobox/listbox interaction pattern.
- Arrow keys, Enter, Escape, mouse, and touch interactions work.
- Empty and recoverable error states are presented appropriately.

## 9. Package and test the complete local application - NOT STARTED

**Goal:** Run and verify the complete system through a single local command.

**Acceptance criteria:**

- `docker compose up` starts the frontend, GraphQL BFF, and autocomplete service.
- Containers use health checks and multi-stage builds where appropriate.
- A browser test covers the complete frontend-to-backend flow.
- A local load test records error rate and warm-response latency.
- The documented setup works on Ubuntu or macOS.

## 10. Define and deploy the AWS infrastructure - NOT STARTED

**Goal:** Publish the application using a low-cost, serverless AWS architecture.

**Acceptance criteria:**

- AWS SAM defines ARM64 Go Lambdas, API Gateway, private S3, and CloudFront.
- The GraphQL Lambda can invoke only the private autocomplete Lambda.
- CloudFront serves the SPA and forwards `/graphql` without caching it.
- Logs expire after 7 days and concurrency limits reduce cost exposure.
- The public application works through a CloudFront HTTPS URL.

## 11. Automate CI and deployment - NOT STARTED

**Goal:** Validate every Pull Request and deploy approved changes without long-lived AWS credentials.

**Acceptance criteria:**

- Pull Requests run Go formatting checks, tests, race detection, lint, frontend tests, builds, and infrastructure validation.
- End-to-end tests exercise the Docker Compose environment.
- Merges to `main` deploy through GitHub Actions using AWS OIDC.
- Frontend assets are synchronized to S3 and CloudFront is invalidated.
- Rollback steps are documented and tested with a safe change.

## 12. Validate production and publish the technical article - NOT STARTED

**Goal:** Measure the deployed system and explain the complete engineering journey in English.

**Acceptance criteria:**

- AWS Budgets sends alerts at 80% and 100% of the USD 5 monthly limit.
- Warm and cold latency, error rate, and core algorithm benchmarks are recorded.
- Mobile Lighthouse performance and accessibility scores are at least 90.
- README includes setup, architecture, testing instructions, screenshots, and the live URL.
- The article explains algorithms, stack choices, trade-offs, cloud architecture, security, costs, and lessons learned.
