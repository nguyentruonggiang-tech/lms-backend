package admin

import (
	"strconv"

	"lms-backend/internal/common/response"
	"lms-backend/internal/dto"
	"lms-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type QuestionHandler struct {
	questionUsecase usecase.QuestionUsecase
}

func NewQuestionHandler(questionUsecase usecase.QuestionUsecase) *QuestionHandler {
	return &QuestionHandler{questionUsecase: questionUsecase}
}

func (h *QuestionHandler) Create(ctx *gin.Context) {
	quizID, err := strconv.Atoi(ctx.Param("quiz_id"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid quiz_id"))
		return
	}

	var body dto.QuestionCreateReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}

	data, err := h.questionUsecase.Create(ctx.Request.Context(), quizID, body)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(data, "success", 0, ctx)
}

func (h *QuestionHandler) FindByQuizID(ctx *gin.Context) {
	quizID, err := strconv.Atoi(ctx.Param("quiz_id"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid quiz_id"))
		return
	}

	data, err := h.questionUsecase.FindByQuizID(ctx.Request.Context(), quizID, ctx.Query("page"), ctx.Query("limit"))
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(data, "success", 0, ctx)
}

func (h *QuestionHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid id"))
		return
	}

	var body dto.QuestionUpdateReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}

	data, err := h.questionUsecase.Update(ctx.Request.Context(), id, body)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(data, "success", 0, ctx)
}

func (h *QuestionHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid id"))
		return
	}

	data, err := h.questionUsecase.Delete(ctx.Request.Context(), id)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(data, "success", 0, ctx)
}
