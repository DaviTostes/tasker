package main

import (
	"log"
	"tasker/internal/gen"
	"tasker/internal/tui"
)

func main() {
	if err := gen.ReadSystemPrompt(); err != nil {
		log.Fatalln(err)
	}

	tui.Run()
}
