# Womba Go CLI

Go client for the Womba AI test generation service.

## Features

- 🚀 Single binary, no dependencies
- ⚡ Fast HTTP client
- 🎨 Colored terminal output
- 🔒 Secure API authentication
- 📦 Easy installation

## Installation

### Option 1: Go Install (Recommended)

```bash
go install github.com/jtizdev/womba-go@latest
```

### Option 2: Download Binary

Download the latest binary for your platform from [Releases](https://github.com/jtizdev/womba-go/releases):

- **macOS (Intel)**: `womba-darwin-amd64`
- **macOS (M1/M2)**: `womba-darwin-arm64`
- **Linux**: `womba-linux-amd64`
- **Windows**: `womba-windows-amd64.exe`

```bash
# macOS/Linux
chmod +x womba-*
sudo mv womba-* /usr/local/bin/womba

# Windows - Add to PATH
```

### Option 3: Build from Source

```bash
git clone https://github.com/jtizdev/womba-go.git
cd womba-go
make build
sudo mv womba /usr/local/bin/
```

## Configuration

Set environment variables:

```bash
# Required
export WOMBA_API_URL="https://womba-api.up.railway.app"
export WOMBA_API_KEY="your-api-key-here"

# Add to ~/.bashrc or ~/.zshrc for persistence
echo 'export WOMBA_API_URL="https://womba-api.up.railway.app"' >> ~/.zshrc
echo 'export WOMBA_API_KEY="your-api-key-here"' >> ~/.zshrc
source ~/.zshrc
```

## Usage

### Generate Tests

Generate test cases for a Jira story:

```bash
# Basic generation (no upload)
womba generate -story PLAT-12991

# Generate and upload to Zephyr
womba generate -story PLAT-12991 -upload
```

### Check API Health

```bash
womba health
```

### Show Version

```bash
womba version
```

## Example Output

```
🚀 Generating tests for PLAT-12991...

✅ Successfully generated 8 test cases!
📊 Quality Score: 88.5/100
📁 Suggested Folder: Orchestration WS/POP ID Alignment
⏱️  Execution Time: 12.34s
🤖 AI Model: gpt-4o

Generated Test Cases:
================================================================================

1. Verify POP ID alignment in orchestration workflow
   Priority: High | Type: Functional
   Description: Test that POP IDs are correctly aligned...
   Steps: 5

2. Test POP ID validation with invalid inputs
   Priority: High | Type: Negative
   Description: Verify system handles invalid POP IDs...
   Steps: 4

...

🎉 Done!
```

## Commands

| Command | Description | Example |
|---------|-------------|---------|
| `generate` | Generate test cases | `womba generate -story PLAT-12991` |
| `health` | Check API status | `womba health` |
| `version` | Show CLI version | `womba version` |

## Flags

### generate

- `-story` (required) - Jira story key (e.g., PLAT-12991)
- `-upload` (optional) - Upload generated tests to Zephyr

## Development

### Prerequisites

- Go 1.21 or higher
- Make (optional)

### Build

```bash
# Build for current platform
make build

# Build for all platforms
make release

# Run tests
make test

# Format code
make fmt

# Clean build artifacts
make clean
```

### Project Structure

```
womba-go/
├── main.go              # CLI entry point
├── client/
│   └── womba_client.go  # HTTP client
├── go.mod               # Go dependencies
├── Makefile             # Build commands
└── README.md            # This file
```

## Architecture

```
┌─────────────┐
│  Go CLI     │
│  (womba-go) │
└──────┬──────┘
       │ HTTP
       ↓
┌─────────────┐
│  Womba API  │
│  (Python)   │
└─────────────┘
```

The Go CLI is a thin wrapper that calls the Womba API service via HTTP. All test generation logic is in the Python service.

## Troubleshooting

### Error: WOMBA_API_URL not set

```bash
export WOMBA_API_URL="https://womba-api.up.railway.app"
```

### Error: WOMBA_API_KEY not set

```bash
export WOMBA_API_KEY="your-api-key"
```

### Error: API error 401

Invalid API key. Check your `WOMBA_API_KEY`.

### Error: API error 403

API key is valid but doesn't have permission.

### Error: Request timeout

The API takes time to generate tests (usually 10-30s). The timeout is set to 120s. If it still times out, check API health:

```bash
womba health
```

## Contributing

1. Fork the repository
2. Create feature branch (`git checkout -b feature/amazing-feature`)
3. Commit changes (`git commit -m 'Add amazing feature'`)
4. Push to branch (`git push origin feature/amazing-feature`)
5. Open Pull Request

## License

MIT License - See [LICENSE](LICENSE) file

## Related Projects

- [womba](https://github.com/jtizdev/womba) - Python CLI (core)
- [womba-api](https://github.com/jtizdev/womba-api) - REST API service
- [womba-java](https://github.com/jtizdev/womba-java) - Java CLI
- [womba-node](https://github.com/jtizdev/womba-node) - Node.js CLI
- [womba-forge](https://github.com/jtizdev/womba-forge) - Atlassian Forge plugin

## Support

- **Issues**: https://github.com/jtizdev/womba-go/issues
- **Docs**: https://github.com/jtizdev/womba
- **API Docs**: https://womba-api.up.railway.app/docs

