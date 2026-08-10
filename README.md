# Go Backend Practice

A hands-on Go learning repository focused on backend and infrastructure engineering fundamentals.

This repository is organized as a sequence of focused practice sessions. The early sessions establish core Go concepts, while later sessions build on and reuse earlier modules to gradually evolve the code toward a small backend service architecture.

The goal is not to build the same application from scratch in every session. Instead, each session introduces a limited set of new concepts while reusing previously implemented code where appropriate.

## Goals

This repository is intended to build practical Go knowledge that maps to backend and infrastructure work, including:

* writing clear and idiomatic Go
* understanding Go modules, packages, and dependency management
* structuring applications around clear responsibilities
* handling errors explicitly
* modeling structured application data
* working with JSON and file persistence
* building HTTP APIs with the standard library
* separating transport, business logic, and storage concerns
* validating incoming requests
* writing unit and HTTP handler tests
* using interfaces and dependency injection
* gradually moving into concurrency, service reliability, and backend tooling

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
    ├── 06-service-layer/
    ├── 07-post-task/
    ├── 08-testing/
    ├── 09-handler-tests/
    └── 10-interfaces/
```

Each session has its own Go module.

Early sessions are mostly standalone. Later sessions intentionally reuse packages from previous sessions through Go module dependencies and local `replace` directives instead of copying already-established code.

This allows each session to focus on the new concept being introduced while keeping earlier implementations reusable.

## Sessions

### 01 — Task Counter

**Focus:** Go modules, packages, exported functions, and basic project structure.

Introduces the basic organization of a Go application and separates simple task-counting and printing behavior into packages.

Key concepts:

* `go.mod`
* packages
* exported vs. unexported identifiers
* slices
* basic function composition
* keeping `main` focused on orchestration

---

### 02 — File Reader

**Focus:** File I/O and explicit error handling.

Replaces hardcoded task input with tasks loaded from a text file.

Key concepts:

* `os.Open`
* `bufio.Scanner`
* `defer`
* returning `(value, error)`
* propagating errors to the caller
* handling recoverable failures without `panic`

---

### 03 — Task Struct

**Focus:** Structs and structured application data.

Moves from plain task strings to a `Task` model containing multiple related fields.

Key concepts:

* structs
* slices of structs
* modeling application entities
* operating on structured data
* sorting tasks by priority
* separating domain data from presentation logic

---

### 04 — JSON Tasks

**Focus:** JSON encoding, decoding, and file persistence.

Adds JSON representations for tasks and utilities for saving and loading structured task data.

Key concepts:

* JSON struct tags
* `json.MarshalIndent`
* `json.Unmarshal`
* file reads and writes
* serialization and deserialization
* reusable JSON utilities
* persistent task data

This session also becomes a reusable source of the `Task` model for several later sessions.

---

### 05 — HTTP Server

**Focus:** Serving data over HTTP.

Introduces Go's `net/http` package and exposes tasks through a `/tasks` endpoint.

Key concepts:

* `http.HandlerFunc`
* `http.ResponseWriter`
* `*http.Request`
* JSON HTTP responses
* response headers
* `http.ListenAndServe`
* HTTP status codes
* reusing packages from another Go module

---

### 06 — Service Layer

**Focus:** Separating HTTP transport from application logic.

Introduces a `TaskService` so the HTTP handler no longer directly handles task-loading behavior.

The application begins moving toward a layered flow:

```text
HTTP Handler
     ↓
Task Service
     ↓
Task Data
```

Key concepts:

* service layer
* separation of concerns
* constructor functions
* dependency injection
* thin HTTP handlers
* keeping transport concerns separate from application behavior

---

### 07 — POST Task and Validation

**Focus:** Accepting input and protecting application data.

Extends the HTTP application with task creation and request validation.

Key concepts:

* `POST /tasks`
* decoding request bodies
* validating incoming data
* rejecting malformed requests
* distinguishing validation failures from business-rule violations
* `201 Created`
* `400 Bad Request`
* duplicate-task handling
* HTTP method routing
* reusable service and handler packages

This session becomes an important reusable application layer for later testing sessions.

---

### 08 — Table-Driven Testing

**Focus:** Automated testing of application behavior.

Introduces Go's built-in `testing` package and applies tests to existing application behavior instead of recreating the application inside the testing session.

Key concepts:

* `_test.go` files
* `go test`
* table-driven tests
* test cases as data
* success and failure paths
* validation tests
* service-layer tests
* using temporary test resources where appropriate

The session primarily contains tests and imports application code implemented in earlier sessions.

---

### 09 — HTTP Handler Testing

**Focus:** Testing HTTP behavior without running a live server.

Uses existing task, service, and handler implementations and places them under an in-memory HTTP test harness.

Key concepts:

* `net/http/httptest`
* `httptest.NewRequest`
* `httptest.NewRecorder`
* Arrange / Act / Assert
* testing HTTP status codes
* testing response headers
* decoding response JSON before asserting application data
* testing handlers without binding to a real network port

The core request flow under test is:

```text
Test Setup
    ↓
Task Service
    ↓
Task Handler
    ↓
Fake HTTP Request
    ↓
Response Recorder
    ↓
Assertions
```

---

### 10 — Interfaces and Repository Abstraction

**Focus:** Decoupling business logic from storage implementations.

Introduces a repository interface between the service layer and task storage.

The architecture evolves toward:

```text
HTTP Handler
     ↓
Task Service
     ↓
TaskRepository
   ↙          ↘
JSON          Memory
Repository    Repository
```

Key concepts:

* Go interfaces
* implicit interface satisfaction
* repository pattern
* dependency inversion
* constructor-based dependency injection
* in-memory repository implementations
* JSON-backed repository implementations
* custom errors
* mock repositories
* testing through abstractions
* swapping implementations without changing callers

Unlike Java, a Go type does not explicitly declare that it `implements` an interface. A type satisfies an interface automatically when it provides the required method set.

## Cross-Session Reuse

Later sessions intentionally import code from earlier sessions.

For example, a testing-focused session may import an existing:

```text
Task model
TaskService
TaskHandler
```

instead of recreating those components.

Some session modules use local replacements such as:

```go
replace github.com/CHA0sTIG3R/go-backend-practice/sessions/07-post-task => ../07-post-task
```

This keeps the repository incremental: new sessions can focus on the concept being learned rather than duplicating previously completed work.

## Running the Sessions

Not every session is intended to run the same way.

### Runnable programs and services

Navigate into the appropriate session:

```bash
cd sessions/07-post-task
```

Then:

```bash
go run .
```

For HTTP sessions, use another terminal to send requests, for example:

```bash
curl http://localhost:8080/tasks
```

### Test-focused sessions

Sessions such as `08-testing` and `09-handler-tests` primarily exercise code through tests.

Run:

```bash
go test ./...
```

For verbose test output:

```bash
go test -v ./...
```

## Go Tooling Checks

Before considering a session complete:

```bash
go fmt ./...
go test ./...
go vet ./...
```

Not every early session contains tests, so `go test ./...` becomes increasingly useful as the repository progresses.

## Git Workflow

After completing or improving a session:

```bash
git status
git add .
git commit -m "Complete session XX <topic>"
git push
```

Commits are kept incremental so the repository also records how the design changes as new concepts are introduced.

## Learning Approach

This repository is intentionally practice-focused.

The code is expected to evolve as new Go and backend concepts are introduced. Earlier implementations may be simple by design, while later sessions add stronger abstractions and testing around them.

The main engineering habits being reinforced are:

* clear package boundaries
* explicit error handling
* small and focused functions
* separation of concerns
* dependency injection
* programming against abstractions where useful
* reusable code instead of unnecessary duplication
* automated testing
* incremental refactoring
* consistent Git usage

The goal is not to introduce abstractions as early as possible. New structure is introduced when the exercises create a reason for it.

## Current Progress

Completed topics include:

* Go modules and packages
* explicit error handling
* file I/O
* structs and structured data
* JSON serialization
* HTTP servers
* service-layer separation
* request validation
* task creation
* unit testing
* table-driven tests
* HTTP handler testing
* interfaces
* repository abstractions
* dependency injection
* in-memory and JSON-backed storage implementations
* mocks and custom errors

## Upcoming Topics

The next phase will move further into Go concepts that are especially relevant to backend and infrastructure engineering, including:

* goroutines
* synchronization
* channels
* data races and safe shared state
* bounded concurrency and worker patterns
* context cancellation and timeouts
* graceful shutdown
* structured logging
* application configuration
* database-backed repositories
* service reliability patterns
* Docker and backend tooling

These topics will continue building on existing code where that reuse makes sense rather than restarting from an empty project each session.

## Author

Built by Hamzat Olowu as a hands-on Go learning track focused on backend and infrastructure engineering.
