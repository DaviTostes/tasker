# Tasker

Tasker is a command-line tool built in Go that provides an interactive terminal user interface (TUI) for refining task descriptions using AI. It leverages Google's Gemini AI model via Firebase Genkit to rewrite user-input tasks into clear, professional titles and descriptions suitable for issue trackers or project management tools.

## Features
- Interactive chat-like interface using Bubble Tea.
- AI-powered task rewriting following structured formats (e.g., `feature/name`, `fix/name`).
- Professional, unambiguous task titles and descriptions.

## Prerequisites
- Go 1.25.1 or later.
- Google AI API key (set as environment variable `GOOGLE_API_KEY` for Genkit authentication).

## Installation
1. Clone the repository:
   ```
   git clone &lt;repo-url&gt;
   cd tasker
   ```
2. Install dependencies:
   ```
   go mod tidy
   ```

## Usage
Build and run the application:
```
go build -o tasker cmd/main.go  && ./tasker
```
- Enter task descriptions in the input field.
- Press Enter to generate refined titles and descriptions.
- Use Ctrl+C or Esc to exit.

## Project Structure
- `cmd/main.go`: Application entry point.
- `internal/model/model.go`: TUI logic using Bubble Tea.
- `internal/gen/client.go`: AI generation logic with Firebase Genkit.
- `system.txt`: System prompt for AI task rewriting.
- `go.mod`: Go module dependencies.

## Coming Features
I will focus on improving the UI for task description refining to make it more intuitive and efficient:

1. **Understand Existing Code**: I will start by reviewing `internal/model/model.go` to grasp the current TUI structure using Bubble Tea.

2. **Enhance Input and Display**: I plan to add components from `github.com/charmbracelet/bubbles` such as spinners for AI loading states, lists for task history, or tables to clearly format output titles and descriptions.

3. **User Experience**: I will implement key bindings for saving tasks, copying to clipboard, or navigating conversation history to enable easier iterative refining.

4. **Visual Polish**: I intend to use Lipgloss for better styling of AI responses, such as bolding titles and sectioning descriptions.

5. **Testing**: I will write unit tests in `internal/` and run them with `go test ./...` to ensure reliability.

These changes will enhance the chat-like interface for seamless task interaction.

## License
MIT License - see [LICENSE](LICENSE) for details.
