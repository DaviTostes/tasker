package gen

import (
	"context"
	"os"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/googlegenai"
)

func GenerateTask(msg string) (string, error) {
	ctx := context.Background()

	g := genkit.Init(
		ctx,
		genkit.WithPlugins(&googlegenai.GoogleAI{}),
		genkit.WithDefaultModel("googleai/gemini-2.5-flash"),
	)

	systemPrompt, err := os.ReadFile("system.txt")
	if err != nil {
		return "", err
	}

	resp, err := genkit.Generate(ctx, g,
		ai.WithSystem(string(systemPrompt)),
		ai.WithPrompt(msg),
	)
	if err != nil {
		return "", err
	}

	return resp.Text(), nil
}
