package main

import (
	"fmt"
	"log"

	core "github.com/ovijeet26/website-summarizer/src/core"
)

func main() {
	summary, err := Summarize("https://go.dev/")
	if err != nil {
		log.Fatal("Error summarizing:", err)
	}
	fmt.Println("Summary:\n", summary)
}

func Summarize(url string) (string, error) {
	website, err := core.NewWebsite(url)
	if err != nil {
		return "", fmt.Errorf("failed to create Website: %w", err)
	}

	messages := website.MessagesFor()

	response, err := core.ChatWithModel(messages)
	if err != nil {
		return "", fmt.Errorf("failed to chat with model: %w", err)
	}

	return response, nil
}
