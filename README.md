# SkyFlare

SkyFlare is a cloud service with a workflow similar to Cloudflare.

## Prerequisites

- **Go** must be installed on your system.

### Installing Go

**On macOS (with Homebrew):**
```bash
brew install go
```

**On Windows:**
1. Download the Go installer from [golang.org/dl](https://golang.org/dl/)
2. Run the installer and follow the setup wizard
3. Verify installation by opening Command Prompt and running:
```cmd
go version
```

**On Linux:**
```bash
# Ubuntu/Debian
sudo apt update
sudo apt install golang-go

# Or using snap
sudo snap install go --classic
```

## Development

### Starting the Application

To start the application in development mode, run:

```bash
make dev
```

### Running Tests

#### Unit Tests

To execute unit tests:

```bash
make unit-test
```

#### All Tests

To run all tests:

```bash
make test
```

For verbose output during testing:

```bash
make test-verbose
```
