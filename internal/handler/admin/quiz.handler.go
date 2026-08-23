package admin

import (
	"strconv"

	"lms-backend/internal/common/response"
	"lms-backend/internal/dto"
	"lms-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type QuizHandler struct {
	quizUsecase usecase.QuizUsecase
}

func NewQuizHandler(quizUsecase usecase.QuizUsecase) *QuizHandler {
	return &QuizHandler{quizUsecase: quizUsecase}
}

func (h *QuizHandler) Create(ctx *gin.Context) {
	lessonID, err := strconv.Atoi(ctx.Param("lesson_id"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid lesson_id"))
		return
	}

	var body dto.QuizCreateReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}

	data, err := h.quizUsecase.Create(ctx.Request.Context(), lessonID, body)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(data, "success", 0, ctx)
}

func (h *QuizHandler) FindByLessonID(ctx *gin.Context) {
	lessonID, err := strconv.Atoi(ctx.Param("lesson_id"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid lesson_id"))
		return
	}

	data, err := h.quizUsecase.FindByLessonID(ctx.Request.Context(), lessonID, ctx.Query("page"), ctx.Query("limit"))
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(data, "success", 0, ctx)
}

func (h *QuizHandler) FindByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid id"))
		return
	}

	data, err := h.quizUsecase.FindByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(data, "success", 0, ctx)
}

func (h *QuizHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid id"))
		return
	}

	var body dto.QuizUpdateReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}

	data, err := h.quizUsecase.Update(ctx.Request.Context(), id, body)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(data, "success", 0, ctx)
}

func (h *QuizHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(response.NewBadRequestException("invalid id"))
		return
	}

	data, err := h.quizUsecase.Delete(ctx.Request.Context(), id)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(data, "success", 0, ctx)
}
