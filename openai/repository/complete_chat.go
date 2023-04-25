package repository

import (
	"encoding/json"
	"jargonjester/entity"
	"net/http"
)

const (
	completeChatURL = "/v1/chat/completions"
)

type completeChatPayload struct {
	Model    string           `json:"model"`
	Messages []entity.Message `json:"messages"`
}

type usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type choices struct {
	Message      entity.Message `json:"message"`
	FinishReason string         `json:"finish_reason"`
	Index        int            `json:"index"`
}

type completeChatResponse struct {
	Id      string    `json:"id"`
	Object  string    `json:"object"`
	Created int       `json:"created"`
	Model   string    `json:"model"`
	Usage   usage     `json:"usage"`
	Choices []choices `json:"choices"`
}

func (r *openaiRepository) CompleteChat(model string, messages []entity.Message) (entity.Message, error) {
	payload := completeChatPayload{
		Model:    model,
		Messages: messages,
	}

	var response completeChatResponse

	resp, err := r.sendRequest(http.MethodPost, completeChatURL, payload)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		return entity.Message{}, err
	}

	return response.Choices[0].Message, nil
}
