# Go Backend Practice

A hands-on Go practice repository focused on backend and infrastructure engineering fundamentals.

This repo is organized as a series of small, focused sessions. Each session builds one concept at a time, starting from basic Go structure and gradually moving into file I/O, APIs, concurrency, testing, and tooling.

## Goals

The purpose of this repo is to practice Go in a way that maps to real backend engineering work:

* writing small, readable Go programs
* organizing code with modules and packages
* separating responsibilities across packages
* handling errors explicitly
* working with files, APIs, and services
* practicing testing and tooling
* building toward small backend services

## Repository Structure

```text
go-backend-practice/
├── README.md
└── sessions/
    ├── 01-task-counter/
    ├── 02-file-reader/
    ├── 03-task-struct/
    ├── 04-json-tasks/
    ├── 05-http-server/
    └── 06-service-layer/
```

Each session is self-contained and includes its own Go module. Later sessions depend on earlier ones via `go.mod` `replace` directives (for example, `05-http-server` and `06-service-layer` both reuse the `task` and `jsonutil` packages from `04-json-tasks`).

## Sessions

### 01 — Task Counter

Focus: Go modules, packages, exported functions, and basic project structure.

This session introduces how Go projects are organized using `go.mod`, packages, and simple reusable functions.

Key concepts:

* Go modules
* packages
* exported vs. unexported functions
* separating logic from program flow
* basic CLI output

### 02 — File Reader

Focus: file I/O and explicit error handling.

This session builds on the task counter by reading tasks from a `tasks.txt` file instead of hardcoding them in `main.go`.

Key concepts:

* reading files
* scanning line by line
* returning `([]string, error)`
* handling errors in the caller
* ignoring blank lines
* keeping `main` as the orchestration layer

### 03 — Task Struct

Focus: structs and sorting.

Tasks become a `Task` struct (`Name`, `Completed`, `Priority`) instead of plain strings, and are sorted by priority with a bubble sort before being reported on.

Key concepts:

* defining structs
* slices of structs
* implementing a basic sort algorithm by hand
* extending `counter`/`printer` to operate on structured data

### 04 — JSON Tasks

Focus: JSON encoding/decoding and file persistence.

This session adds a `jsonutil` package for marshaling `Task` structs to/from JSON, saving them to disk, and loading them back.

Key concepts:

* JSON struct tags
* `json.MarshalIndent` / `json.Unmarshal`
* reading and writing files with `os`
* round-tripping data through a `tasks.json` file
* returning wrapped errors from utility functions

### 05 — HTTP Server

Focus: serving data over HTTP.

Introduces a minimal `net/http` server that exposes an in-memory list of tasks as JSON over a `/tasks` endpoint, reusing the `task` package from session 04.

Key concepts:

* `http.HandlerFunc`
* writing JSON responses with `json.NewEncoder`
* setting response headers
* `http.ListenAndServe`
* cross-module reuse via `go.mod` `replace` directives

### 06 — Service Layer

Focus: separating business logic from HTTP transport.

Refactors the session 05 server to introduce a `TaskService` that loads tasks from `tasks.json`, keeping the HTTP handler thin and focused on request/response concerns.

Key concepts:

* service layer pattern
* dependency injection via constructor functions (`NewTaskService`, `NewTaskHandler`)
* restricting handlers to allowed HTTP methods
* separating I/O (loading tasks) from transport (serving them)
* consistent HTTP error responses

## How to Run a Session

Navigate into a session folder:

```bash
cd sessions/06-service-layer
```

Run the program:

```bash
go run .
```

For the HTTP server sessions (`05-http-server`, `06-service-layer`), this starts a server on `:8080`. In another terminal, query it:

```bash
curl http://localhost:8080/tasks
```

Format the code:

```bash
go fmt ./...
```

Check for suspicious issues:

```bash
go vet ./...
```

## Git Workflow

After completing or improving a session:

```bash
git status
git add .
git commit -m "Complete session 06 service layer"
git push
```

## Learning Notes

This repo is intentionally practice-focused. The code may start simple, but each session is meant to reinforce habits used in production Go services:

* clear package boundaries
* small functions
* explicit error handling
* readable program flow
* incremental improvements
* consistent use of Git

## Upcoming Topics

Planned practice areas include:

* routing (multiple endpoints, path parameters)
* request validation
* unit testing
* table-driven tests
* concurrency with goroutines
* channels
* context cancellation
* logging
* configuration
* Docker basics
* small service design

## Author

Built by Hamzat Olowu as part of a backend and infrastructure-focused Go learning track.
