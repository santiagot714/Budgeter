# Budgeter — Setup Guide

Instructions for setting up the development environment.
This file is updated as the project evolves.

---

## Development environment

### Requirements

| Tool           | Min version | Notes                               |
|----------------|-------------|-------------------------------------|
| WSL2           | 2.6+        | Distro: Ubuntu or equivalent        |
| Go             | 1.26.2      | Installed inside WSL                |
| Docker         | 29.0+       | Docker Desktop with WSL integration |
| Docker Compose | v2.40+      | Included with Docker Desktop        |
| Git            | any         | Inside WSL                          |
| Make           | 4.3+        | `sudo apt install make`             |
| golang-migrate | latest      | See installation below              |

### Installing Go in WSL (one-time setup)

```bash
# Download Go
curl -L https://go.dev/dl/go1.26.2.linux-amd64.tar.gz -o /tmp/go1.26.2.tar.gz

# Extract to home directory
mkdir -p $HOME/go-sdk
tar -C $HOME/go-sdk -xzf /tmp/go1.26.2.tar.gz

# Add to PATH in ~/.bashrc
echo '' >> ~/.bashrc
echo '# Go' >> ~/.bashrc
echo 'export GOROOT=$HOME/go-sdk/go' >> ~/.bashrc
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
echo 'export PATH=$PATH:$GOROOT/bin:$GOPATH/bin' >> ~/.bashrc

# Apply changes
source ~/.bashrc

# Verify
go version
```

### Installing golang-migrate (one-time setup)

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Verify
migrate --version
```

### Clone the project

```bash
git clone git@github.com:santiagot714/Budgeter.git
cd Budgeter
```

### Install Go dependencies

```bash
go mod download
```

---

## Local infrastructure (Docker)

```bash
docker compose up -d
```

---

## Environment variables

Copy `.env.example` to `.env` and fill in the required values:

```bash
cp .env.example .env
```

---

## Database migrations

```bash
# Apply all pending migrations
make migrate

# Revert last migration
make migrate-down

# Create a new migration
make migrate-create name=<migration_name>

# Check current migration version
make migrate-version

# Force a specific version (use after a dirty state)
make migrate-force version=<version_number>
```

---

## Useful commands (Makefile)

| Command                            | Description                        |
|------------------------------------|------------------------------------|
| `make migrate`                     | Apply all pending migrations       |
| `make migrate-down`                | Revert last migration              |
| `make migrate-create name=<name>`  | Create a new migration             |
| `make migrate-version`             | Show current migration version     |
| `make migrate-force version=<v>`   | Force version (fix dirty state)    |
