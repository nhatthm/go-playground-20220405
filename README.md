# Playground 20220405

[![test-with-docker-compose](https://github.com/nhatthm/go-playground-20220405/actions/workflows/test-with-docker-compose.yaml/badge.svg)](https://github.com/nhatthm/go-playground-20220405/actions/workflows/test-with-docker-compose.yaml)
[![test-with-testcontainers](https://github.com/nhatthm/go-playground-20220405/actions/workflows/test-with-testcontainers.yaml/badge.svg)](https://github.com/nhatthm/go-playground-20220405/actions/workflows/test-with-testcontainers.yaml)

A showcase of differences between `docker-compose` and `testcontainers`.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Usage](#usage)

## Prerequisites

- `Go >= 1.18`

[<sub><sup>[table of contents]</sup></sub>](#table-of-contents)

## Usage

### Test with `docker-compose`

Run

```bash
docker compose up -d
make test-with-docker-compose DATABASE_DSN=mongodb://localhost:27017/playground
docker compose down
```

[<sub><sup>[table of contents]</sup></sub>](#table-of-contents)

### Test with `testcontainers`

Run

```bash
make test-with-testcontainers
```

[<sub><sup>[table of contents]</sup></sub>](#table-of-contents)
