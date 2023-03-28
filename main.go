package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "context"
    "github.com/lsparrish/chatbot-for-study/utils"
    "github.com/sashabaranov/go-openai"
)

func main() {
    // Get API key from .env file
    apiKey := utils.GetApiKey()

    // Create OpenAI client
    client := openai.NewClient(apiKey)

    // Set up reader for user input
    reader := bufio.NewReader(os.Stdin)

    // Start chat loop
    for {
        // Get user input
        fmt.Print("You: ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        // Check for exit command
        if strings.ToLower(input) == "exit" {
            fmt.Println("Goodbye!")
            return
        }

        // Generate response using OpenAI API
        ctx := context.Background()
        completions, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
            Messages: []openai.ChatCompletionMessage{
                { 
                    Content: "You are a helpful code generator.",
                    Role: "system",
                },
                {
                    Content: input,
                    Role: "user",
                },
                
            },
            Model:       openai.GPT3Dot5Turbo,
            MaxTokens:   1000,
            Temperature: 0.7,
        })

        if err != nil {
            log.Fatal("Error making OpenAI API request: ", err)
        }


        // Print response
        response := completions.Choices[0].Message.Content
      
        fmt.Println("Bot:", response)
    }
}
