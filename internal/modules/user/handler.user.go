package user

import (
	"net/http"
	"scriptmake/internal/apperror"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	userUseCase UserUseCase
	//jwtService
}

func NewUserHandler(userUseCase UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

var validate = validator.New()

func (h *UserHandler) Register(ctx *gin.Context) {
	var body RegisterUserDTOInput
	if err := ctx.ShouldBindJSON(body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato JSON invalido!"})
		return
	}

	// Validando usando tags JSON do struct
	if err := validate.Struct(body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userUseCase.Register(body)
	if err != nil {
		if err == apperror.ErrEmailAlreadyExist {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email ja esta registrado! Faça o login."})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"suecess": true, "message": "Usuário criado com sucesso!"})
}

func (h *UserHandler) Login(ctx *gin.Context) {
	var body LoginUserDTOInput
	if err := ctx.ShouldBindJSON(body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato JSON invalido!"})
		return
	}

	if err := validate.Struct(body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.userUseCase.Login(body)
	if err != nil {
		if err == apperror.ErrInvalidCredentials {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email ou senha invalido!"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"token": token})

}

func (h *UserHandler) Viewprofile(ctx *gin.Context) {
	userID := ctx.GetInt("userID")
	result, err := h.userUseCase.Viewprofile(userID)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"sucess": true, "result": result})
}
