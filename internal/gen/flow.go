package gen

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/googlegenai"
)

type Task struct {
	Title        string `json:"title"`
	Overview     string `json:"overview"`
	Requirements string `json:"requirements"`
	Criteria     string `json:"criteria"`
	Impact       string `json:"impact"`
}

func (t *Task) RenderMd() (string, error) {
	text := fmt.Sprint(
		"# ",
		t.Title,
		"\n ## Overview\n",
		t.Overview,
		"\n ## Requirements\n",
		t.Requirements,
		"\n ## Criteria\n",
		t.Criteria,
		"\n ## Impact\n",
		t.Impact,
	)

	return glamour.Render(text, "dark")
}

var (
	ctx = context.Background()
	g   = genkit.Init(
		ctx,
		genkit.WithPlugins(&googlegenai.GoogleAI{}),
		genkit.WithDefaultModel("googleai/gemini-2.5-flash"),
	)

	TaskFlow = genkit.DefineFlow(
		g,
		"taskFlow",
		func(ctx context.Context, msg string) (Task, error) {
			systemPrompt, err := os.ReadFile("system.txt")
			if err != nil {
				return Task{}, err
			}

			task, _, err := genkit.GenerateData[Task](ctx, g,
				ai.WithSystem(string(systemPrompt)),
				ai.WithPrompt(msg),
			)
			return *task, nil
		},
	)
)
