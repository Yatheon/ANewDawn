package main

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

/*
Prompt is the Discord message that the user sent to the bot.
*/
func GenerateGPT4Response(prompt string, author string) (string, error) {
	openAIToken := config.OpenAIToken
	if openAIToken == "" {
		return "", fmt.Errorf("OPENAI_API_KEY is not set")
	}

	// Print the prompt
	fmt.Println("Prompt:", author, ":", prompt)

	// Check if the prompt is too long
	if len(prompt) > 2048 {
		return "", fmt.Errorf("prompt is too long")
	}


	// Add additional information to the system message
	var additionalInfo string
	switch author {
	case "thelovinator":
		additionalInfo = "User (TheLovinator) is a programmer. Wants to live in the woods. Real name is Joakim. He made the bot."
	case "killyoy":
		additionalInfo = "User (KillYoy) likes to play video games. Real name is Andreas. Good at CSS."
	case "forgefilip":
		additionalInfo = "User (ForgeFilip) likes watches. Real name is Filip."
	case "plubplub":
		additionalInfo = "User (Piplup) likes to play WoW and Path of Exile. Real name is Axel. Is also called Bambi."
	case "nobot":
		additionalInfo = "User (Nobot) likes to play WoW. Real name is Gustav. Really good at programming."
	case "kao172":
		additionalInfo = "User (kao172) likes cars. Real name is Fredrik."
	}

	// Create a new client
	client := openai.NewClient(openAIToken)

	// System message
	var systemMessage string
	systemMessage = `You are in a Discord server. 
	You are Swedish. 
	Use Markdown for formatting. 
	Please respond with a short message. 
	`

	// Add additional information to the system message
	if additionalInfo != "" {
		systemMessage = fmt.Sprintf("%s\n%s", systemMessage, additionalInfo)
	}

	// Print the system message
	fmt.Println("System message:", systemMessage)

	// Create a completion
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4TurboPreview,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemMessage,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		return "", fmt.Errorf("failed to get response from OpenAI: %w", err)
	}

	ourResponse := resp.Choices[0].Message.Content

	fmt.Println("Response:", ourResponse)
	return ourResponse, nil
}
