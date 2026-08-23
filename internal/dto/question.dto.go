package dto

type QuestionCreateReq struct {
	QuestionText  string `json:"question_text" binding:"required"`
	OptionA       string `json:"option_a" binding:"required"`
	OptionB       string `json:"option_b" binding:"required"`
	OptionC       string `json:"option_c" binding:"required"`
	OptionD       string `json:"option_d" binding:"required"`
	CorrectOption string `json:"correct_option" binding:"required,oneof=A B C D"`
}

type QuestionUpdateReq struct {
	QuestionText  *string `json:"question_text" binding:"omitempty"`
	OptionA       *string `json:"option_a" binding:"omitempty"`
	OptionB       *string `json:"option_b" binding:"omitempty"`
	OptionC       *string `json:"option_c" binding:"omitempty"`
	OptionD       *string `json:"option_d" binding:"omitempty"`
	CorrectOption *string `json:"correct_option" binding:"omitempty,oneof=A B C D"`
}
