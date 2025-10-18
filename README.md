# Womba CLI (Go)

> AI-powered test generation for Jira stories

## Install

```bash
go install github.com/jtizdev/womba-go@latest
```

## Usage

```bash
# Setup
export WOMBA_API_URL="https://womba-api.onrender.com"
export WOMBA_API_KEY="your-api-key"

# Generate tests
womba generate -story PLAT-12991

# Generate and upload to Zephyr
womba generate -story PLAT-12991 -upload
```

## License

MIT · [Womba](https://github.com/jtizdev/womba)
