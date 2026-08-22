package repository

import (
	"context"
	"lms-backend/ent"
	"lms-backend/internal/dto"
)

type UserRepository interface {
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	CreateUser(ctx context.Context, req dto.AuthRegisterReq) (*ent.Users, error)
	FindUserByEmail(ctx context.Context, email string) (*ent.Users, error)
	FindUserById(ctx context.Context, id int) (*ent.Users, error)
	UpdatePassword(ctx context.Context, id int, hashedPassword string) error
}
