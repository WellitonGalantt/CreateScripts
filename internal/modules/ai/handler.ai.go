package ai

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AiHandler struct {
	chatService *ChatService
}

func NewAiHandler(chatService *ChatService) *AiHandler {
	return &AiHandler{
		chatService: chatService,
	}
}

func (h *AiHandler) TesteComunication(ctx *gin.Context) {
	response, err := h.chatService.ProcessMessage("Olá, como você está?")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{"success": true, "response": response})
}
