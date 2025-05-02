package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// OpenAI API endpoint and model constants
const (
	OpenAIEndpoint = "http://localhost:11434/v1/chat/completions"
	ModelName      = "llama3.2"
	APIKey         = "ollama"
)

func ChatWithModel(messages []Message) (string, error) {
	// Set up the request
	chatReq := ChatRequest{
		Model:    ModelName,
		Messages: messages,
	}

	// Convert request to JSON
	jsonData, err := json.Marshal(chatReq)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", OpenAIEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+APIKey)

	// Send the request
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Decode the response
	var chatResp ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&chatResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	if len(chatResp.Choices) == 0 {
		return "", fmt.Errorf("no choices returned from the model")
	}

	return chatResp.Choices[0].Message.Content, nil
}
