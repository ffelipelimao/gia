<p align="center">
    <h1>gia</h1>
  <img src="image/logo.png" alt="gia logo" width="300" />
</p>

A CLI tool for AI-driven task execution using the Gemini API

### Requirements
- Go 1.21+
- Git
- A Google AI Studio API key (`GEMINI_API_KEY`)

### Setup
1. Create a env at your .bashrc or .zshrc with your API key and your prompt
```bash
export GEMINI_API_KEY="your_api_key_here"
export GEMINI_PROMPT="your_prompt_to_create_commit"
```

### Installation
Install the binary globally (recommended):
```bash
go install github.com/ffelipelimao/gia@latest
```

### Usage

#### Commit Command
Generate and execute a git commit with AI assistance:

```bash
# Using the full command name
gia commit

# Using the alias
gia c
```

The command provides an interactive interface that:
1. Reads the repository's current `git diff`
2. Generates a commit message using AI
3. Presents you with options:
   - **`a`** - Accept and commit (executes the commit with the generated message)
   - **`r`** - Regenerate message (generates a new message using AI)
   - **`e`** - Edit manually (allows you to enter your own commit message)
   - **`q`** - Quit (exits without committing)

#### Example Interaction
```
📝 Generated commit message:
feat: add interactive commit message generation

Options:
  [a] Accept and commit
  [r] Regenerate message
  [e] Edit manually
  [q] Quit

Choose an option: a
✅ Commit executed successfully!
```

#### Help
Get help for any command:
```bash
gia --help
gia commit --help
gia c --help
```

### Project structure
- `gia.go`: CLI entry point with Cobra commands
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

# Test
# Another test
