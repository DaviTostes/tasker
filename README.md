# Tasker

Tasker is a command-line tool built in Go that provides an interactive terminal user interface (TUI) for refining task descriptions using AI. It leverages Google's Gemini AI model via Firebase Genkit to transform vague user inputs into professional, structured task documentation suitable for software engineering projects.

## Features

- **Interactive TUI**: Built with Bubble Tea for a polished terminal experience
- **AI-powered task refinement**: Uses Gemini 2.5 Flash to rewrite and structure task descriptions
- **Structured output**: Generates professional task documentation with title, overview, requirements, criteria, and impact
- **Editing capabilities**: Built-in vim-style editor (vimtea) for modifying generated content
- **Clipboard integration**: Quick copy functionality for generated tasks
- **Markdown rendering**: Beautiful presentation of task details using Glamour
- **Multi-language support**: Maintains the language of the original user input

## Task Structure

Each generated task includes:

- **Title**: Standardized format (feature/name, fix/name, refactor/name, task/name)
- **Overview**: Brief summary of the task's purpose and importance
- **Requirements**: Detailed steps, inputs, outputs, and constraints
- **Criteria**: 3-5 verifiable conditions for successful implementation
- **Impact**: Assessment of potential effects on the project system

## Prerequisites

- Go 1.25.1 or later
- Google AI API key (set as environment variable `GOOGLE_API_KEY` for Genkit authentication)

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/davitostes/tasker
   cd tasker
   ```
2. Install dependencies:
   ```
   go mod tidy
   ```

## Usage

Build and run the application:
```bash
go build -o tasker cmd/main.go && ./tasker
```

### Controls

- **Enter**: Generate refined task from input
- **Ctrl+C**: Quit application
- **Ctrl+Y**: Copy generated task to clipboard
- **Ctrl+E**: Edit generated task in vim-style editor
- **Ctrl+S**: Save/exit edit mode
- **Esc**: Exit edit mode (when in editor)

### Workflow

1. Enter a task description in the input field
2. Press Enter to generate a refined, structured task
3. Use Ctrl+E to edit the generated content if needed
4. Use Ctrl+S to save edits and return to the main interface
5. Use Ctrl+Y to copy the final task to clipboard for pasting into issue trackers or documentation

## Project Structure

- `cmd/main.go`: Application entry point
- `internal/tui/`: Terminal user interface implementation
  - `model.go`: TUI model and initialization
  - `functions.go`: Generation functionality
  - `update.go`: Update logic and key bindings
- `internal/gen/`: AI generation logic
  - `flow.go`: Firebase Genkit flow definition and Task struct
- `prompts/system.prompt`: AI system prompt for task refinement
- `go.mod`: Go module dependencies

## Dependencies

- **Bubble Tea**: TUI framework
- **Firebase Genkit**: AI orchestration
- **Gemini 2.5 Flash**: AI model
- **Glamour**: Markdown rendering
- **vimtea**: Built-in text editor
- **Charm libraries**: UI components (lipgloss, etc.)

## License

MIT License - see [LICENSE](LICENSE) for details.
