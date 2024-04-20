// pkg/actions/actions.go
package actions

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

func GenerateCode(prompt string) error {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("OPENAI_API_KEY environment variable is not set")
	}

	client := openai.NewClient(apiKey)

	messages := []openai.ChatCompletionMessage{
		{
			Role: openai.ChatMessageRoleSystem,
			Content: `You are an AI that builds CLI tools for users.
You are able to call functions to build the tool.
Build what you can with the functions.
Also write a script to create the necessary AWS resources to run this and a script to deploy the tool there.
After you've done your work scaffolding the tool as much as you can, the user will take over and complete the work.`,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: prompt,
		},
	}

	functions := []openai.FunctionDefinition{
		// Define your functions here
	}

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:     openai.GPT4,
			Messages:  messages,
			Functions: functions,
		},
	)
	if err != nil {
		return fmt.Errorf("error creating chat completion: %v", err)
	}

	fmt.Println(resp.Choices[0].Message.Content)

	return nil
}
