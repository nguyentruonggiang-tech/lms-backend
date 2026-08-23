package usecase_impl

import (
	"context"
	"fmt"
	"math"

	"github.com/google/uuid"

	"lms-api/internal/common/pagination"
	"lms-api/internal/common/rabbitmq"
	"lms-api/internal/common/response"
	"lms-api/internal/repository"
	"lms-api/internal/usecase"
)

type certificateIssuedPayload struct {
	UserID  int    `json:"userId"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type certificateUsecase struct {
	certificateRepository repository.CertificateRepository
	enrollmentRepository  repository.EnrollmentRepository
	quizAttemptRepository repository.QuizAttemptRepository
	rabbitmq              *rabbitmq.RabbitMQ
}

func NewCertificateUsecase(
	certificateRepository repository.CertificateRepository,
	enrollmentRepository repository.EnrollmentRepository,
	quizAttemptRepository repository.QuizAttemptRepository,
	rabbitmq *rabbitmq.RabbitMQ,
) usecase.CertificateUsecase {
	return &certificateUsecase{
		certificateRepository: certificateRepository,
		enrollmentRepository:  enrollmentRepository,
		quizAttemptRepository: quizAttemptRepository,
		rabbitmq:              rabbitmq,
	}
}

func (u *certificateUsecase) CheckAndIssueCertificate(ctx context.Context, userID, courseID int) {
	exists, err := u.certificateRepository.ExistsByUserAndCourse(ctx, userID, courseID)
	if err != nil || exists {
		return
	}

	enrollment, err := u.enrollmentRepository.FindByUserAndCourse(ctx, userID, courseID)
	if err != nil || enrollment.ProgressPercent < 100 {
		return
	}

	passed, err := u.quizAttemptRepository.HasPassedAllByUserAndCourse(ctx, userID, courseID)
	if err != nil || !passed {
		return
	}

	code := uuid.New().String()
	cert, err := u.certificateRepository.Create(ctx, userID, courseID, code)
	if err != nil {
		return
	}

	_ = u.enrollmentRepository.CompleteEnrollment(ctx, userID, courseID)

	_ = u.rabbitmq.Send(ctx, "notification.certificate_issued", certificateIssuedPayload{
		UserID:  userID,
		Title:   fmt.Sprintf("Chúc mừng bạn đã hoàn thành khóa học \"%s\"", enrollment.Edges.Courses.Title),
		Content: fmt.Sprintf("Chứng chỉ của bạn có mã: %s", cert.Code),
	})
}

func (u *certificateUsecase) GetMyCertificates(ctx context.Context, userID int, page, limit string) (any, error) {
	query := pagination.Get(page, limit)

	data, err := u.certificateRepository.FindAllByUserID(ctx, userID, query)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	total, err := u.certificateRepository.CountByUserID(ctx, userID)
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

func (u *certificateUsecase) GetByCode(ctx context.Context, code string) (any, error) {
	data, err := u.certificateRepository.FindByCode(ctx, code)
	if err != nil {
		return nil, response.NewNotFoundException()
	}
	return data, nil
}
