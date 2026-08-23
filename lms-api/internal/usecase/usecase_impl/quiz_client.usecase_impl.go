package usecase_impl

import (
	"context"
	"math"

	"lms-api/internal/common/pagination"
	"lms-api/internal/common/response"
	"lms-api/internal/dto"
	"lms-api/internal/repository"
	"lms-api/internal/usecase"
)

type quizClientUsecase struct {
	quizRepository        repository.QuizRepository
	questionRepository    repository.QuestionRepository
	quizAttemptRepository repository.QuizAttemptRepository
	enrollmentRepository  repository.EnrollmentRepository
	certificateUsecase    usecase.CertificateUsecase
}

func NewQuizClientUsecase(
	quizRepository repository.QuizRepository,
	questionRepository repository.QuestionRepository,
	quizAttemptRepository repository.QuizAttemptRepository,
	enrollmentRepository repository.EnrollmentRepository,
	certificateUsecase usecase.CertificateUsecase,
) usecase.QuizClientUsecase {
	return &quizClientUsecase{
		quizRepository:        quizRepository,
		questionRepository:    questionRepository,
		quizAttemptRepository: quizAttemptRepository,
		enrollmentRepository:  enrollmentRepository,
		certificateUsecase:    certificateUsecase,
	}
}

func (u *quizClientUsecase) GetQuiz(ctx context.Context, userID, quizID int) (any, error) {
	quiz, err := u.quizRepository.FindByID(ctx, quizID)
	if err != nil {
		return nil, response.NewNotFoundException()
	}

	_, err = u.enrollmentRepository.FindByUserAndCourse(ctx, userID, quiz.CourseID)
	if err != nil {
		return nil, response.NewForbiddenException("not enrolled in this course")
	}

	allQuestions, err := u.questionRepository.FindAllByQuizID(ctx, quizID)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	publicQuestions := make([]*dto.QuestionPublic, 0, len(allQuestions))
	for _, q := range allQuestions {
		publicQuestions = append(publicQuestions, &dto.QuestionPublic{
			ID:           q.ID,
			QuizID:       q.QuizID,
			QuestionText: q.QuestionText,
			OptionA:      q.OptionA,
			OptionB:      q.OptionB,
			OptionC:      q.OptionC,
			OptionD:      q.OptionD,
		})
	}

	return dto.QuizWithQuestionsRes{
		Quiz:      quiz,
		Questions: publicQuestions,
	}, nil
}

func (u *quizClientUsecase) Submit(ctx context.Context, userID, quizID int, body dto.QuizSubmitReq) (any, error) {
	quiz, err := u.quizRepository.FindByID(ctx, quizID)
	if err != nil {
		return nil, response.NewNotFoundException()
	}

	_, err = u.enrollmentRepository.FindByUserAndCourse(ctx, userID, quiz.CourseID)
	if err != nil {
		return nil, response.NewForbiddenException("not enrolled in this course")
	}

	allQuestions, err := u.questionRepository.FindAllByQuizID(ctx, quizID)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	answerMap := make(map[int]string, len(body.Answers))
	for _, a := range body.Answers {
		answerMap[a.QuestionID] = a.SelectedOption
	}

	correct := 0
	for _, q := range allQuestions {
		if answerMap[q.ID] == string(q.CorrectOption) {
			correct++
		}
	}

	total := len(allQuestions)
	score := 0.0
	if total > 0 {
		score = float64(correct) / float64(total) * 100
	}
	isPassed := int(math.Round(score)) >= quiz.PassingScore

	attempt, err := u.quizAttemptRepository.Create(ctx, userID, quizID, total, correct, score, isPassed)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	go u.certificateUsecase.CheckAndIssueCertificate(ctx, userID, quiz.CourseID)

	return attempt, nil
}

func (u *quizClientUsecase) GetAttempts(ctx context.Context, userID, quizID int, page, limit string) (any, error) {
	quiz, err := u.quizRepository.FindByID(ctx, quizID)
	if err != nil {
		return nil, response.NewNotFoundException()
	}

	_, err = u.enrollmentRepository.FindByUserAndCourse(ctx, userID, quiz.CourseID)
	if err != nil {
		return nil, response.NewForbiddenException("not enrolled in this course")
	}

	query := pagination.Get(page, limit)

	data, err := u.quizAttemptRepository.FindByUserAndQuiz(ctx, userID, quizID, query)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	total, err := u.quizAttemptRepository.CountByUserAndQuiz(ctx, userID, quizID)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	return pagination.Response[any]{
		Items:     data,
		Page:      query.Page,
		Limit:     query.Limit,
		TotalItem: total,
		TotalPage: int(math.Ceil(float64(total) / float64(query.Limit))),
	}, nil
}
