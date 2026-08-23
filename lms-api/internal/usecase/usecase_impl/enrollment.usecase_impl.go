package usecase_impl

import (
	"context"
	"fmt"
	"math"

	"lms-api/internal/common/pagination"
	"lms-api/internal/common/rabbitmq"
	"lms-api/internal/common/response"
	"lms-api/internal/dto"
	"lms-api/internal/repository"
	"lms-api/internal/usecase"
)

type enrollmentUsecase struct {
	enrollmentRepository repository.EnrollmentRepository
	courseRepository     repository.CourseRepository
	rabbitmq             *rabbitmq.RabbitMQ
}

func NewEnrollmentUsecase(enrollmentRepository repository.EnrollmentRepository, courseRepository repository.CourseRepository, rabbitmq *rabbitmq.RabbitMQ) usecase.EnrollmentUsecase {
	return &enrollmentUsecase{
		enrollmentRepository: enrollmentRepository,
		courseRepository:     courseRepository,
		rabbitmq:             rabbitmq,
	}
}

type courseEnrolledPayload struct {
	UserID  int    `json:"userId"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (u *enrollmentUsecase) Enroll(ctx context.Context, userID int, body dto.EnrollReq) (any, error) {
	course, err := u.courseRepository.FindPublishedByID(ctx, body.CourseID)
	if err != nil {
		return nil, response.NewNotFoundException("course not found")
	}

	exists, err := u.enrollmentRepository.ExistsByUserAndCourse(ctx, userID, body.CourseID)
	if err != nil {
		return nil, response.NewInternalServerErrorException(err.Error())
	}
	if exists {
		return nil, response.NewBadRequestException("already enrolled in this course")
	}

	data, err := u.enrollmentRepository.Create(ctx, userID, body.CourseID)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	_ = u.rabbitmq.Send(ctx, "notification.course_enrolled", courseEnrolledPayload{
		UserID:  userID,
		Title:   "Đăng ký khóa học thành công",
		Content: fmt.Sprintf("Bạn đã đăng ký thành công khóa học \"%s\"", course.Title),
	})

	return data, nil
}

func (u *enrollmentUsecase) FindMyEnrollments(ctx context.Context, userID int, status, page, limit string) (any, error) {
	query := pagination.Get(page, limit)

	data, err := u.enrollmentRepository.FindAllByUserID(ctx, userID, status, query)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	total, err := u.enrollmentRepository.CountByUserID(ctx, userID, status)
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

func (u *enrollmentUsecase) FindMyCourse(ctx context.Context, userID, courseID int) (any, error) {
	data, err := u.enrollmentRepository.FindByUserAndCourse(ctx, userID, courseID)
	if err != nil {
		return nil, response.NewNotFoundException()
	}
	return data, nil
}

func (u *enrollmentUsecase) Cancel(ctx context.Context, userID, enrollmentID int) (any, error) {
	enrollment, err := u.enrollmentRepository.FindByID(ctx, enrollmentID)
	if err != nil {
		return nil, response.NewNotFoundException()
	}

	if enrollment.UserID != userID {
		return nil, response.NewForbiddenException()
	}

	if err := u.enrollmentRepository.Delete(ctx, enrollmentID); err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}
	return true, nil
}
