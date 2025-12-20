package ai

type ChatService struct {
	aiClient *Client
}

func NewChatService(aiClient *Client) *ChatService {
	return &ChatService{
		aiClient: aiClient,
	}
}

func (s *ChatService) ProcessMessage(message string) (string, error) {
	// Aqui você pode adicionar lógica de negócio antes/depois
	// de chamar a IA: validações, logging, cache, etc.

	model := "openai/gpt-3.5-turbo"
	response, err := s.aiClient.SendMessage(model, message)
	if err != nil {
		return "", err
	}

	return response, nil
}
