package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"jargonjester/domain"
	"net/http"
)

type openaiRepository struct {
	httpClient http.Client
	host       string
	apikey     string
}

func NewOpenaiRepository(host string, apikey string) domain.OpenaiRepository {
	return &openaiRepository{
		httpClient: http.Client{},
		host:       host,
		apikey:     apikey,
	}
}

func (r *openaiRepository) sendRequest(method, resource string, requestBody interface{}) (*http.Response, error) {
	payload := bytes.NewBuffer(nil)
	if requestBody != nil {
		json, _ := json.Marshal(requestBody)
		payload.Write(json)
	}

	req, _ := http.NewRequestWithContext(context.Background(), method, r.host+resource, payload)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+r.apikey)

	resp, _ := r.httpClient.Do(req)

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf(string(body))
	}

	return resp, nil
}
