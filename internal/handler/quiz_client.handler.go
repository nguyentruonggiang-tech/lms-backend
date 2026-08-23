package handler

import (
	"strconv"

	"lms-backend/internal/common/helpers"
	"lms-backend/internal/common/response"
	"lms-backend/internal/dto"
	"lms-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type QuizClientHandler struct {
	quizClientUsecase usecase.QuizClientUsecase
}

func NewQuizClientHandler(quizClientUsecase usecase.QuizClientUsecase) *QuizClientHandler {
	return &QuizClientHandler{quizClientUsecase: quizClientUsecase}
}

func (h *QuizClientHandler) GetQuiz(ctx *gin.Context) {
	user, err := helpers.GetUser(ctx)
	if err != nil {
		ctx.Error(response.NewUnauthorizedException())
		return
	}

	quizID, err := strconv.Atoi(ctx.Param("quizId"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid quizId"))
		return
	}

	data, err := h.quizClientUsecase.GetQuiz(ctx.Request.Context(), user.ID, quizID)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(data, "success", 0, ctx)
}

func (h *QuizClientHandler) Submit(ctx *gin.Context) {
	user, err := helpers.GetUser(ctx)
	if err != nil {
		ctx.Error(response.NewUnauthorizedException())
		return
	}

	quizID, err := strconv.Atoi(ctx.Param("quizId"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid quizId"))
		return
	}

	var body dto.QuizSubmitReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}

	data, err := h.quizClientUsecase.Submit(ctx.Request.Context(), user.ID, quizID, body)
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(data, "submitted successfully", 0, ctx)
}

func (h *QuizClientHandler) GetAttempts(ctx *gin.Context) {
	user, err := helpers.GetUser(ctx)
	if err != nil {
		ctx.Error(response.NewUnauthorizedException())
		return
	}

	quizID, err := strconv.Atoi(ctx.Param("quizId"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid quizId"))
		return
	}

	data, err := h.quizClientUsecase.GetAttempts(ctx.Request.Context(), user.ID, quizID, ctx.Query("page"), ctx.Query("limit"))
	if err != nil {
		ctx.Error(err)
		return
	}
	response.Success(data, "success", 0, ctx)
}
