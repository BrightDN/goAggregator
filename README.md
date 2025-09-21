# goAggregator

A simple RSS aggregator. The gator CLI fetches and stores RSS posts and lets you browse them. Accounts are username-based.

## Prerequisites

- Go 1.20+ installed and on PATH
  - echo 'export PATH="$PATH:$(go env GOPATH)/bin"' >> ~/.bashrc && source ~/.bashrc
- PostgreSQL installed and running

## Install

- Latest:
  - go install github.com/BrightDN/goAggregator/cmd/gator@latest
- Specific version:
  - go install github.com/BrightDN/goAggregator/cmd/gator@v0.1.0

## Configure

gator reads config from ~/.gatorconfig.json.

Create it:
```bash
cat > ~/.gatorconfig.json <<'JSON'
{
  "db_url": "postgres://user:pass@localhost:5432/gator?sslmode=disable",
  "current_user_name": "default"
}
JSON
```
Fields:
- db_url: PostgreSQL connection string
- current_user_name: your gator username

## Database setup (Goose)

This project uses Goose for schema migrations.

- Install Goose:
  - go install github.com/pressly/goose/v3/cmd/goose@latest

- Create DB (if needed):
  - createdb gator

- Run migrations:
  - goose -dir ./sql/schema postgres "postgres://user:pass@localhost:5432/gator?sslmode=disable" up

- Reset data only (does NOT run migrations):
  - gator reset

If you prefer manual setup, apply the SQL files in sql/schema/ in order.

## Usage

Common commands:
- Login user:
  - gator login <username>
- Create user:
  - gator register <username>
- Add a feed:
  - gator addfeed <feed_name> <feed_url>
- Follow a feed:
  - gator follow <feed_url>
- Aggregate (fetch):
  - gator agg <timestring>
- Browse posts:
  - gator browse <OPTIONAL: limit> (Defaults to 2)

## Development

- Build locally:
  - go build -o gator ./cmd/gator
  - ./gator help

## Notes

- Binaries install to $GOBIN (or $GOPATH)/bin. Ensure it’s on PATH.
- Config file: ~/.gatorconfig.json (edit directly or via CLI that updates CurrentUserName).
- Repo: https://github.com/BrightDN/goAggregator