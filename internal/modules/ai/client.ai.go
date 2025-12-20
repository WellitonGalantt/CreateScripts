package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	apiKey     string
	baseUrl    string
	httpClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:  apiKey,
		baseUrl: "https://openrouter.ai",
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Client) Chat(model string, messages []Message) (*ChatResponse, error) {
	reqBody := ChatRequest{
		Model:    model,
		Messages: messages,
	}

	//Converte a struct para JSON
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("erro ao serializar requiusicao %w", err)
	}

	req, err := http.NewRequest("POST", c.baseUrl+"/api/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao executar requisição: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler resposta: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API retornou erro (status %d): %s", resp.StatusCode, string(body))
	}

	var chatResp ChatResponse
	if err := json.Unmarshal(body, &chatResp); err != nil {
		return nil, fmt.Errorf("erro ao deserializar resposta: %w", err)
	}

	return &chatResp, nil
}

// Método auxiliar para mensagens simples
func (c *Client) SendMessage(model, userMessage string) (string, error) {
	messages := []Message{
		{
			Role:    "user",
			Content: userMessage,
		},
	}

	resp, err := c.Chat(model, messages)
	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("nenhuma resposta recebida")
	}

	return resp.Choices[0].Message.Content, nil
}
