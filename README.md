# Core Wallet Service

Simple wallet management service built with Go for managing user balances, top-up, and withdrawal operations.

## Tech Stack

- **Language**: Go (Golang)
- **Framework**: Fiber v2
- **Database**: PostgreSQL
- **ORM**: GORM
- **Cache**: Redis

## Features

- Wallet creation and balance management
- Top-up operations (request, confirm, reject)
- Withdrawal operations (in development)
- Transaction history tracking
- SHA256 checksum validation for data integrity

## Quick Start

```bash
# Install dependencies
go mod download

# Run the service
go run main.go
```

Service runs on `http://localhost:3000`

## Notes

- User management handled by separate service
- Authentication assumed to be handled externally
- Transaction pattern: Request (PENDING) → Confirm (SUCCESS) / Reject (FAILED)