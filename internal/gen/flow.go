package gen

import (
	"context"
	"fmt"

	"github.com/charmbracelet/glamour"
	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/googlegenai"
	"google.golang.org/genai"
)

type Task struct {
	Title        string `json:"title"`
	Overview     string `json:"overview"`
	Requirements string `json:"requirements"`
	Criteria     string `json:"criteria"`
	Impact       string `json:"impact"`
}

func (t *Task) GetText() string {
	return fmt.Sprint(
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
}

func (t *Task) RenderMd() (string, error) {
	return glamour.Render(t.GetText(), "dark")
}

var (
	ctx = context.Background()
	g   = genkit.Init(
		ctx,
		genkit.WithPlugins(&googlegenai.GoogleAI{}),
		genkit.WithDefaultModel("googleai/gemini-2.5-flash"),
	)

	TaskFlow = genkit.DefineFlow(g, "taskFlow",
		func(ctx context.Context, msg string) (Task, error) {
			prompt := genkit.LookupPrompt(g, "system")

			tBudget := int32(0)

			resp, err := prompt.Execute(ctx, ai.WithInput(map[string]any{"input": msg}), ai.WithConfig(&genai.GenerateContentConfig{
				ThinkingConfig: &genai.ThinkingConfig{
					ThinkingBudget: &tBudget,
				},
			}))
			if err != nil {
				return Task{}, err
			}

			var task Task
			if err := resp.Output(&task); err != nil {
				return Task{}, err
			}

			return task, err
		},
	)
)
