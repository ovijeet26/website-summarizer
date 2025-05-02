package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func ChatWithModel(messages []Message) (string, error) {
	url := "http://localhost:11434/v1/chat/completions"

	chatReq := ChatRequest{
		Model:    "llama3.2",
		Messages: messages,
	}

	jsonData, err := json.Marshal(chatReq)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer ollama")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var chatResp ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&chatResp); err != nil {
		return "", err
	}

	if len(chatResp.Choices) == 0 {
		return "", fmt.Errorf("no choices returned")
	}

	return chatResp.Choices[0].Message.Content, nil
}
