<p align="center">
    <h1>gia</h1>
  <img src="image/logo.png" alt="gia logo" width="300" />
</p>

A commit message generator using the Gemini API

### Requirements
- Go 1.21+
- Git
- A Google AI Studio API key (`GEMINI_API_KEY`)

### Setup
1. Create a `.env` file at the project root with:
```bash
GEMINI_API_KEY=your_api_key_here
```

### Usage
Install the binary globally (recommended):
```bash
go install github.com/ffelipelimao/gia@latest
```

Then, inside a Git repository with local changes, run:
```bash
gia
```

The command reads the repository's current `git diff` and prints a suggested commit message to stdout.

### Project structure
- `gia.go`: CLI entry point
- `internal/ai`: Google AI (Gemini) API integration
- `internal/exec`: Git command execution (diff collection and commit helper)
- `Makefile`: convenience target for `go run`

### Contributing
We welcome contributions! Steps:
1. (Optional) Open an issue describing the improvement or bug
2. Fork the repository
3. Create a branch from `main`:
```bash
git checkout -b feat/my-improvement
```
4. Make your changes
5. Ensure formatting and basic checks pass:
6. Update documentation if needed
7. Open a Pull Request explaining what and why

Additional guidelines:
- Keep commits using gia
- Prefer descriptive names for variables and functions
- Provide helpful error messages (the project follows this pattern)

### Credentials notes
- Do be stupid and not commit or expose secrets
- Use `.env` locally

### License
This project is distributed under the MIT License. See `LICENSE` for details.

