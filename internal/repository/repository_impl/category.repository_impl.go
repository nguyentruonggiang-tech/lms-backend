package repository_impl

import (
	"context"
	"lms-backend/ent"
	"lms-backend/internal/common/pagination"
	"lms-backend/internal/dto"
	"lms-backend/internal/repository"
)

type categoryRepository struct {
	client *ent.Client
}

func NewCategoryRepository(client *ent.Client) repository.CategoryRepository {
	return &categoryRepository{client: client}
}

func (r *categoryRepository) Create(ctx context.Context, body dto.CategoryCreateReq) (*ent.Categories, error) {
	q := r.client.Categories.Create().
		SetName(body.Name).
		SetSlug(body.Slug)

	if body.Description != nil {
		q = q.SetDescription(*body.Description)
	}

	return q.Save(ctx)
}

func (r *categoryRepository) FindAll(ctx context.Context, query pagination.Query) ([]*ent.Categories, error) {
	return r.client.Categories.Query().
		Limit(query.Limit).
		Offset(query.Offset).
		All(ctx)
}

func (r *categoryRepository) Count(ctx context.Context) (int, error) {
	return r.client.Categories.Query().Count(ctx)
}

func (r *categoryRepository) FindByID(ctx context.Context, id int) (*ent.Categories, error) {
	return r.client.Categories.Get(ctx, id)
}

func (r *categoryRepository) Update(ctx context.Context, id int, body dto.CategoryUpdateReq) (*ent.Categories, error) {
	q := r.client.Categories.UpdateOneID(id)

	if body.Name != nil {
		q = q.SetName(*body.Name)
	}
	if body.Slug != nil {
		q = q.SetSlug(*body.Slug)
	}
	if body.Description != nil {
		q = q.SetDescription(*body.Description)
	}

	return q.Save(ctx)
}

func (r *categoryRepository) Delete(ctx context.Context, id int) error {
	return r.client.Categories.DeleteOneID(id).Exec(ctx)
}
