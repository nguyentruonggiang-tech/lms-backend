package usecase

import (
	"context"
	"lms-api/internal/dto"
)

type EnrollmentUsecase interface {
	Enroll(ctx context.Context, userID int, body dto.EnrollReq) (any, error)
	FindMyEnrollments(ctx context.Context, userID int, status, page, limit string) (any, error)
	FindMyCourse(ctx context.Context, userID, courseID int) (any, error)
	Cancel(ctx context.Context, userID, enrollmentID int) (any, error)
}
