# Code Editing Agent

A Go-based agent framework that uses Claude to help with code editing tasks. The agent can read files, list directories, and interact with the user through a chat interface.

## Features

- Interactive chat interface with Claude
- File system tools:
  - Read file contents
  - List files and directories
- Extensible tool system for adding new capabilities

## Prerequisites

- Go 1.24 or later
- Anthropic API key

## Setup

1. Clone the repository:
```bash
git clone https://github.com/dleon86/code-editing-agent.git
cd code-editing-agent
```

2. Create a `.env` file in the project root with your Anthropic API key:
```
ANTHROPIC_API_KEY=your_api_key_here
```

3. Install dependencies:
```bash
go mod tidy
```

## Usage

Run the agent:
```bash
go run cmd/agent/main.go
```

The agent will start an interactive chat session where you can:
- Ask questions about code
- Request file contents
- List directory contents
- Get help with code editing tasks

## Project Structure

```
.
├── cmd/
│   └── agent/         # Main application entry point
├── pkg/
│   ├── agent/         # Core agent functionality
│   └── tools/         # Tool implementations
└── README.md
```

## Adding New Tools

To add a new tool:
1. Create the tool implementation in `pkg/tools/`
2. Define the tool in `pkg/tools/tools.go`
3. Add it to the tools list in `cmd/agent/main.go`

## License

MIT 