package usecase_impl

import (
	"context"
	"math"

	"github.com/google/uuid"

	"lms-api/internal/common/pagination"
	"lms-api/internal/common/response"
	"lms-api/internal/repository"
	"lms-api/internal/usecase"
)

type certificateUsecase struct {
	certificateRepository repository.CertificateRepository
	enrollmentRepository  repository.EnrollmentRepository
	quizAttemptRepository repository.QuizAttemptRepository
}

func NewCertificateUsecase(
	certificateRepository repository.CertificateRepository,
	enrollmentRepository repository.EnrollmentRepository,
	quizAttemptRepository repository.QuizAttemptRepository,
) usecase.CertificateUsecase {
	return &certificateUsecase{
		certificateRepository: certificateRepository,
		enrollmentRepository:  enrollmentRepository,
		quizAttemptRepository: quizAttemptRepository,
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
	_, err = u.certificateRepository.Create(ctx, userID, courseID, code)
	if err != nil {
		return
	}

	_ = u.enrollmentRepository.CompleteEnrollment(ctx, userID, courseID)
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
