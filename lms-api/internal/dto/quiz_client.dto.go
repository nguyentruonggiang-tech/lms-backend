package dto

import "lms-api/ent"

type QuestionPublic struct {
	ID           int    `json:"id"`
	QuizID       int    `json:"quiz_id"`
	QuestionText string `json:"question_text"`
	OptionA      string `json:"option_a"`
	OptionB      string `json:"option_b"`
	OptionC      string `json:"option_c"`
	OptionD      string `json:"option_d"`
}

type QuizWithQuestionsRes struct {
	Quiz      *ent.Quizzes      `json:"quiz"`
	Questions []*QuestionPublic `json:"questions"`
}

type QuizAnswerItem struct {
	QuestionID     int    `json:"questionId" binding:"required"`
	SelectedOption string `json:"selectedOption" binding:"required,oneof=A B C D"`
}

type QuizSubmitReq struct {
	Answers []QuizAnswerItem `json:"answers" binding:"required,min=1"`
}
