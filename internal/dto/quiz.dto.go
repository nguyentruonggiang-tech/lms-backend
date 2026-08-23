package dto

type QuizCreateReq struct {
	Title        string `json:"title" binding:"required"`
	PassingScore *int   `json:"passing_score" binding:"omitempty"`
}

type QuizUpdateReq struct {
	Title        *string `json:"title" binding:"omitempty"`
	PassingScore *int    `json:"passing_score" binding:"omitempty"`
}
