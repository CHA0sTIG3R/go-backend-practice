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
├── notes/
└── sessions/
    ├── 01-task-counter/
    └── 02-file-reader/
```

Each session is self-contained and includes its own Go module.

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

## How to Run a Session

Navigate into a session folder:

```bash
cd sessions/02-file-reader
```

Run the program:

```bash
go run .
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
git commit -m "Complete session 02 file reader"
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

* HTTP servers
* JSON APIs
* routing
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
