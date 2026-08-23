package repository_impl

import (
	"context"
	"lms-backend/ent"
	"lms-backend/ent/users"
	"lms-backend/internal/common/pagination"
	"lms-backend/internal/dto"
	"lms-backend/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type userRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) repository.UserRepository {
	return &userRepository{client: client}
}

func (r *userRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	return r.client.Users.Query().
		Where(users.EmailEQ(email), users.DeletedAtIsNil()).
		Exist(ctx)
}

func (r *userRepository) CreateUser(ctx context.Context, req dto.AuthRegisterReq) (*ent.Users, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return r.client.Users.Create().
		SetEmail(req.Email).
		SetPassword(string(hashed)).
		SetFullName(req.FullName).
		Save(ctx)
}

func (r *userRepository) FindUserByEmail(ctx context.Context, email string) (*ent.Users, error) {
	return r.client.Users.Query().
		Where(users.EmailEQ(email), users.DeletedAtIsNil()).
		Only(ctx)
}

func (r *userRepository) FindUserById(ctx context.Context, id int) (*ent.Users, error) {
	return r.client.Users.Query().
		Where(users.IDEQ(id), users.DeletedAtIsNil()).
		Only(ctx)
}

func (r *userRepository) UpdatePassword(ctx context.Context, id int, hashedPassword string) error {
	return r.client.Users.UpdateOneID(id).SetPassword(hashedPassword).Exec(ctx)
}

func (r *userRepository) FindAll(ctx context.Context, query pagination.Query, role, status, keyword string) ([]*ent.Users, error) {
	q := r.client.Users.Query().Where(users.DeletedAtIsNil())
	if role != "" {
		q = q.Where(users.RoleEQ(users.Role(role)))
	}
	if status != "" {
		q = q.Where(users.StatusEQ(users.Status(status)))
	}
	if keyword != "" {
		q = q.Where(users.Or(
			users.EmailContainsFold(keyword),
			users.FullNameContainsFold(keyword),
		))
	}
	return q.Limit(query.Limit).Offset(query.Offset).All(ctx)
}

func (r *userRepository) CountAll(ctx context.Context, role, status, keyword string) (int, error) {
	q := r.client.Users.Query().Where(users.DeletedAtIsNil())
	if role != "" {
		q = q.Where(users.RoleEQ(users.Role(role)))
	}
	if status != "" {
		q = q.Where(users.StatusEQ(users.Status(status)))
	}
	if keyword != "" {
		q = q.Where(users.Or(
			users.EmailContainsFold(keyword),
			users.FullNameContainsFold(keyword),
		))
	}
	return q.Count(ctx)
}

func (r *userRepository) UpdateStatus(ctx context.Context, id int, status string) (*ent.Users, error) {
	return r.client.Users.UpdateOneID(id).SetStatus(users.Status(status)).Save(ctx)
}

func (r *userRepository) UpdateRole(ctx context.Context, id int, role string) (*ent.Users, error) {
	return r.client.Users.UpdateOneID(id).SetRole(users.Role(role)).Save(ctx)
}
