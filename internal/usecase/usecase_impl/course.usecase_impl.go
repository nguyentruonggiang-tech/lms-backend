package usecase_impl

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"math"
	"time"

	"lms-backend/internal/common/cache"
	"lms-backend/internal/common/elastic"
	"lms-backend/internal/common/pagination"
	"lms-backend/internal/common/response"
	"lms-backend/internal/dto"
	"lms-backend/internal/repository"
	"lms-backend/internal/usecase"
)

type courseUsecase struct {
	courseRepository repository.CourseRepository
	lessonRepository repository.LessonRepository
	redisClient      *cache.RedisClient
	esClient         *elastic.ElasticClient
}

func NewCourseUsecase(
	courseRepository repository.CourseRepository,
	lessonRepository repository.LessonRepository,
	redisClient *cache.RedisClient,
	esClient *elastic.ElasticClient,
) usecase.CourseUsecase {
	return &courseUsecase{
		courseRepository: courseRepository,
		lessonRepository: lessonRepository,
		redisClient:      redisClient,
		esClient:         esClient,
	}
}

func (u *courseUsecase) indexCourse(ctx context.Context, id int) {
	data, err := u.courseRepository.FindByID(ctx, id)
	if err != nil {
		return
	}
	categoryName := ""
	if data.Edges.Categories != nil {
		categoryName = data.Edges.Categories.Name
	}
	u.esClient.IndexCourse(ctx, elastic.CourseDoc{
		ID:           data.ID,
		Title:        data.Title,
		Description:  data.Description,
		CategoryID:   data.CategoryID,
		CategoryName: categoryName,
		Level:        data.Level.String(),
		Price:        data.Price,
		Status:       data.Status.String(),
		CreatedAt:    data.CreatedAt.Format(time.RFC3339),
	})
}

func (u *courseUsecase) Create(ctx context.Context, body dto.CourseCreateReq) (any, error) {
	data, err := u.courseRepository.Create(ctx, body)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}
	u.redisClient.DelByPattern(ctx, "courses:*")
	u.indexCourse(ctx, data.ID)
	return data, nil
}

func (u *courseUsecase) FindAll(ctx context.Context, page, limit string) (any, error) {
	query := pagination.Get(page, limit)

	data, err := u.courseRepository.FindAll(ctx, query)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	total, err := u.courseRepository.Count(ctx)
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

func (u *courseUsecase) FindByID(ctx context.Context, id int) (any, error) {
	data, err := u.courseRepository.FindByID(ctx, id)
	if err != nil {
		return nil, response.NewNotFoundException()
	}
	return data, nil
}

func (u *courseUsecase) Update(ctx context.Context, id int, body dto.CourseUpdateReq) (any, error) {
	data, err := u.courseRepository.Update(ctx, id, body)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}
	u.redisClient.DelByPattern(ctx, "courses:*")
	u.indexCourse(ctx, data.ID)
	return data, nil
}

func (u *courseUsecase) UpdateStatus(ctx context.Context, id int, body dto.CourseUpdateStatusReq) (any, error) {
	data, err := u.courseRepository.UpdateStatus(ctx, id, body.Status)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}
	u.redisClient.DelByPattern(ctx, "courses:*")
	u.indexCourse(ctx, data.ID)
	return data, nil
}

func (u *courseUsecase) Delete(ctx context.Context, id int) (any, error) {
	err := u.courseRepository.Delete(ctx, id)
	if err != nil {
		return nil, response.NewNotFoundException()
	}
	u.redisClient.DelByPattern(ctx, "courses:*")
	u.esClient.DeleteCourse(ctx, id)
	return true, nil
}

func (u *courseUsecase) FindAllPublished(ctx context.Context, filter dto.CoursePublicFilter, page, limit string) (any, error) {
	cacheKey := courseListKey(filter, page, limit)
	if cached, err := u.redisClient.Get(ctx, cacheKey); err == nil {
		var result pagination.Response[any]
		if json.Unmarshal([]byte(cached), &result) == nil {
			return result, nil
		}
	}

	query := pagination.Get(page, limit)

	data, err := u.courseRepository.FindAllPublished(ctx, filter, query)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	total, err := u.courseRepository.CountPublished(ctx, filter)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	result := pagination.Response[any]{
		Items:     data,
		Page:      query.Page,
		Limit:     query.Limit,
		TotalItem: total,
		TotalPage: int(math.Ceil(float64(total) / float64(query.Limit))),
	}

	if b, err := json.Marshal(result); err == nil {
		u.redisClient.Set(ctx, cacheKey, string(b), 60*time.Second)
	}

	return result, nil
}

func (u *courseUsecase) SearchPublished(ctx context.Context, filter dto.CoursePublicFilter, page, limit string) (any, error) {
	query := pagination.Get(page, limit)

	if filter.Q == "" {
		data, err := u.courseRepository.SearchPublished(ctx, filter, query)
		if err != nil {
			return nil, response.NewBadRequestException(err.Error())
		}
		total, err := u.courseRepository.CountSearch(ctx, filter)
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

	docs, total, err := u.esClient.SearchCourses(ctx, filter.Q, query.Offset, query.Limit)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	if len(docs) == 0 {
		return pagination.Response[any]{
			Items:     []any{},
			Page:      query.Page,
			Limit:     query.Limit,
			TotalItem: 0,
			TotalPage: 0,
		}, nil
	}

	return pagination.Response[any]{
		Items:     docs,
		Page:      query.Page,
		Limit:     query.Limit,
		TotalItem: total,
		TotalPage: int(math.Ceil(float64(total) / float64(query.Limit))),
	}, nil
}

func (u *courseUsecase) FindPublishedByID(ctx context.Context, id int) (any, error) {
	cacheKey := fmt.Sprintf("courses:detail:%d", id)
	if cached, err := u.redisClient.Get(ctx, cacheKey); err == nil {
		var result any
		if json.Unmarshal([]byte(cached), &result) == nil {
			return result, nil
		}
	}

	data, err := u.courseRepository.FindPublishedByID(ctx, id)
	if err != nil {
		return nil, response.NewNotFoundException()
	}

	if b, err := json.Marshal(data); err == nil {
		u.redisClient.Set(ctx, cacheKey, string(b), 60*time.Second)
	}

	return data, nil
}

func courseListKey(filter dto.CoursePublicFilter, page, limit string) string {
	raw := fmt.Sprintf("%v|%v|%s|%v|%v|%s|%s",
		filter.CategoryID, filter.Level, filter.Q,
		filter.MinPrice, filter.MaxPrice, page, limit,
	)
	h := md5.Sum([]byte(raw))
	return fmt.Sprintf("courses:list:%x", h)
}

func (u *courseUsecase) FindPreviewLessons(ctx context.Context, courseID int) (any, error) {
	_, err := u.courseRepository.FindPublishedByID(ctx, courseID)
	if err != nil {
		return nil, response.NewNotFoundException()
	}

	data, err := u.lessonRepository.FindPreviewByCourseID(ctx, courseID)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}
	return data, nil
}

