package repository_impl

import (
	"context"
	"lms-backend/ent"
	"lms-backend/ent/users"
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
