package gpt4

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

func Generate(fileName string, fileContents string, outputType string) {
	client := openai.NewClient("your-key-here")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			ResponseFormat: &openai.ChatCompletionResponseFormat{
				Type: openai.ChatCompletionResponseFormatTypeJSONObject,
			},
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: `
						You are a TypeScript developer assistant who writes tests using Jest.
						You are given a typescript file name and the contents of the file.
						You will be instructed to write either unit tests or mocks for the file.
						Your expected output is json in the following format: 
						{
							"file_name": "{originalFileName}.{mock/test}.ts",
							"file_contents": "{yourFileContents}"
						}
						Your input will always include the original file name and contents as well as the desired output type (mock or test)
					`,
				},
				{
					Role: openai.ChatMessageRoleUser,
					// Message should look like this:
					// File name: {fileName}
					// File contents: {fileContents}
					// Desired output: {mock/test}
					Content: fmt.Sprintf(`
						File name: %s
						File contents: %s
						Desired output: %s
						`, fileName, fileContents, outputType),
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
