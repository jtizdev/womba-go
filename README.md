# Womba Go CLI

AI-powered test generation for Jira stories.

## Installation

```bash
go install github.com/jtizdev/womba-go@latest
```

## Configuration

```bash
export WOMBA_API_URL="https://womba-api.up.railway.app"
export WOMBA_API_KEY="your-api-key"
```

## Usage

```bash
# Generate tests
womba generate -story PLAT-12991

# Generate and upload to Zephyr
womba generate -story PLAT-12991 -upload

# Check API health
womba health
```

## Support

- [Main Docs](https://github.com/jtizdev/womba)
- [Issues](https://github.com/jtizdev/womba-go/issues)
