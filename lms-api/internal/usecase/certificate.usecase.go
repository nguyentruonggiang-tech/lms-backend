package usecase

import "context"

type CertificateUsecase interface {
	CheckAndIssueCertificate(ctx context.Context, userID, courseID int)
	GetMyCertificates(ctx context.Context, userID int, page, limit string) (any, error)
	GetByCode(ctx context.Context, code string) (any, error)
}
