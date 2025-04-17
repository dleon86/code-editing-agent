package tools

import (
	"agent/pkg/agent"
)

var ReadFileDefinition = agent.ToolDefinition{
	Name:        "read_file",
	Description: "Read the contents of a given relative file path. Use this when you want to see what's inside a file. Do not use this with directory names.",
	InputSchema: *ReadFileInputSchema,
	Function:    ReadFile,
}

var ListFilesDefinition = agent.ToolDefinition{
	Name:        "list_files",
	Description: "List files and directories at a given path. If no path is provided, lists files in the current directory.",
	InputSchema: *ListFilesInputSchema,
	Function:    ListFiles,
}
