package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

		body, err := ioutil.ReadAll(resp.Body)

		fmt.Println("joidwqjoid")
		fmt.Println(string(body))
		fmt.Println(resp.StatusCode)
		return nil, err
	}

	return resp, nil
}
