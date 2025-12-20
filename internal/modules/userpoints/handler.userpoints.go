package userpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserPointsHandler struct {
	uc UserPointsUseCase
}

func NewUserPointsHandler(uc UserPointsUseCase) *UserPointsHandler {
	return &UserPointsHandler{
		uc: uc,
	}
}

func (h *UserPointsHandler) GetById(ctx *gin.Context) {
	userId := ctx.GetString("userID")
	result, err := h.uc.GetById(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "result": result})
}

func (h *UserPointsHandler) Debit(ctx *gin.Context) {
	var body CreditValuesDTOInput
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato JSON invalido!"})
		return
	}

	userId := ctx.GetString("userID")
	body.UserId = userId

	err := h.uc.Debit(body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true})

}

func (h *UserPointsHandler) Credt(ctx *gin.Context) {
	var body CreditValuesDTOInput
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato JSON invalido!"})
		return
	}

	userId := ctx.GetString("userID")
	body.UserId = userId

	err := h.uc.Credit(body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"sucess": true})
}

func (h *UserPointsHandler) GetBalance(ctx *gin.Context) {
	userId := ctx.GetString("userID")
	result, err := h.uc.GetBalance(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "balance": result})

}

func (h *UserPointsHandler) GetTransactions(ctx *gin.Context) {
	userId := ctx.GetString("userID")
	result, err := h.uc.GetTransactions(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "transactions": result})
}
