# Reporting API

## Overview

A Golang-based API for generating analytical reports for an e-commerce system.

## Prerequisites

- Go 1.18+
- PostgreSQL
- Redis (for caching)

## Setup

1. Clone the repository.
2. Run `go mod tidy` to install dependencies.
3. Update the `.env` file with your configuration.
4. Run database migrations using the scripts in `db/migrations/`.

## Run the Application

```bash
go run cmd/main.go
```
