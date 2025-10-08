package inputs

import (
	"os"
	"slices"
	"strings"
)

var (
	path = "inputs.txt"
)

func Read() ([]string, error) {
	_, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(path)
	str := string(content)
	if str == "" {
		return []string{}, nil
	}

	splited := strings.Split(str, "\n")
	return splited, nil
}

func Add(newInput string, inputs []string) ([]string, error) {
	if slices.Contains(inputs, newInput) {
		return inputs, nil
	}

	inputs = append(inputs, newInput)
	joined := strings.Join(inputs, "\n")

	err := os.WriteFile("inputs.txt", []byte(joined), 0777)
	return inputs, err
}

func Get(inputs []string, i int, direction int) (string, int) {
	i += direction

	if i >= len(inputs) {
		return "", len(inputs)-1
	}

	if i < 0 {
		return inputs[0], 0
	}

	return inputs[i], i
}
