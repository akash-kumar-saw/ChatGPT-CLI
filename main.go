package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	gpt3 "github.com/PullRequestInc/go-gpt3"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func getResponse(context context.Context, client gpt3.Client, question string) {
	err := client.CompletionStreamWithEngine(context, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		Prompt: []string{
			question,
		},
		MaxTokens:   gpt3.IntPtr(3000),
		Temperature: gpt3.Float32Ptr(0),
	}, func(response *gpt3.CompletionResponse) {
		fmt.Printf(response.Choices[0].Text)
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}

type NullWriter int

func (NullWriter) Write([]byte) (int, error) {
	return 0, nil
}

func main() {
	log.SetOutput(new(NullWriter))

	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	apiKey := viper.GetString("API_KEY")
	if apiKey == "" {
		fmt.Println("ChatGPT API Key is Missing")
		return
	}

	fmt.Printf("***** ChatGPT - Press 'Q' to Exit *****")

	context := context.Background()
	client := gpt3.NewClient(apiKey)

	rootCmd := &cobra.Command{
		Use:   "ChatGPT - Command Line Interface (CLI)",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			scanner := bufio.NewScanner(os.Stdin)
			run := true

			for run {
				fmt.Printf("\nYou : ")
				if !scanner.Scan() {
					break
				}

				question := scanner.Text()
				if question == "" {
					continue
				}
				if question == "Q" || question == "q" {
					break
				}
				getResponse(context, client, question)
			}
		},
	}
	rootCmd.Execute()
	fmt.Printf("\n***** Thanks for Using ChatGPT *****")
}
