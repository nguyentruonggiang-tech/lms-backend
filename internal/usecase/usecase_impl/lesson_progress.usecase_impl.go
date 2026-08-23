package usecase_impl

import (
	"context"

	"lms-backend/internal/common/pagination"
	"lms-backend/internal/common/response"
	"lms-backend/internal/dto"
	"lms-backend/internal/repository"
	"lms-backend/internal/usecase"
)

type lessonProgressUsecase struct {
	lessonProgressRepository repository.LessonProgressRepository
	enrollmentRepository     repository.EnrollmentRepository
	lessonRepository         repository.LessonRepository
}

func NewLessonProgressUsecase(
	lessonProgressRepository repository.LessonProgressRepository,
	enrollmentRepository repository.EnrollmentRepository,
	lessonRepository repository.LessonRepository,
) usecase.LessonProgressUsecase {
	return &lessonProgressUsecase{
		lessonProgressRepository: lessonProgressRepository,
		enrollmentRepository:     enrollmentRepository,
		lessonRepository:         lessonRepository,
	}
}

func (u *lessonProgressUsecase) ListLessons(ctx context.Context, userID, courseID int) (any, error) {
	_, err := u.enrollmentRepository.FindByUserAndCourse(ctx, userID, courseID)
	if err != nil {
		return nil, response.NewForbiddenException("not enrolled in this course")
	}

	lessonList, err := u.lessonRepository.FindByCourseID(ctx, courseID, pagination.Query{Limit: 1000})
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	progressList, err := u.lessonProgressRepository.FindAllByUserAndCourse(ctx, userID, courseID)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	progressMap := make(map[int]*dto.LessonWithProgress, len(progressList))
	for _, p := range progressList {
		progressMap[p.LessonID] = &dto.LessonWithProgress{
			IsCompleted: p.IsCompleted,
			CompletedAt: p.CompletedAt,
		}
	}

	results := make([]dto.LessonWithProgress, 0, len(lessonList))
	for _, lesson := range lessonList {
		item := dto.LessonWithProgress{Lesson: lesson}
		if p, ok := progressMap[lesson.ID]; ok {
			item.IsCompleted = p.IsCompleted
			item.CompletedAt = p.CompletedAt
		}
		results = append(results, item)
	}

	return results, nil
}

func (u *lessonProgressUsecase) FindLesson(ctx context.Context, userID, lessonID int) (any, error) {
	lesson, err := u.lessonRepository.FindByID(ctx, lessonID)
	if err != nil {
		return nil, response.NewNotFoundException()
	}

	_, err = u.enrollmentRepository.FindByUserAndCourse(ctx, userID, lesson.CourseID)
	if err != nil {
		return nil, response.NewForbiddenException("not enrolled in this course")
	}

	result := dto.LessonWithProgress{Lesson: lesson}
	if p, err := u.lessonProgressRepository.FindByUserAndLesson(ctx, userID, lessonID); err == nil {
		result.IsCompleted = p.IsCompleted
		result.CompletedAt = p.CompletedAt
	}

	return result, nil
}

func (u *lessonProgressUsecase) Complete(ctx context.Context, userID, lessonID int) (any, error) {
	lesson, err := u.lessonRepository.FindByID(ctx, lessonID)
	if err != nil {
		return nil, response.NewNotFoundException("lesson not found")
	}

	_, err = u.enrollmentRepository.FindByUserAndCourse(ctx, userID, lesson.CourseID)
	if err != nil {
		return nil, response.NewForbiddenException("not enrolled in this course")
	}

	existing, err := u.lessonProgressRepository.FindByUserAndLesson(ctx, userID, lessonID)
	var result any
	if err != nil {
		result, err = u.lessonProgressRepository.Create(ctx, userID, lesson.CourseID, lessonID)
		if err != nil {
			return nil, response.NewBadRequestException(err.Error())
		}
	} else if !existing.IsCompleted {
		result, err = u.lessonProgressRepository.MarkComplete(ctx, existing.ID)
		if err != nil {
			return nil, response.NewBadRequestException(err.Error())
		}
	} else {
		result = existing
	}

	u.syncProgressPercent(ctx, userID, lesson.CourseID)

	return result, nil
}

func (u *lessonProgressUsecase) GetProgress(ctx context.Context, userID, courseID int) (any, error) {
	_, err := u.enrollmentRepository.FindByUserAndCourse(ctx, userID, courseID)
	if err != nil {
		return nil, response.NewForbiddenException("not enrolled in this course")
	}

	data, err := u.lessonProgressRepository.FindAllByUserAndCourse(ctx, userID, courseID)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}
	return data, nil
}

func (u *lessonProgressUsecase) syncProgressPercent(ctx context.Context, userID, courseID int) {
	completed, err := u.lessonProgressRepository.CountCompletedByUserAndCourse(ctx, userID, courseID)
	if err != nil {
		return
	}
	total, err := u.lessonRepository.CountByCourseID(ctx, courseID)
	if err != nil || total == 0 {
		return
	}
	percent := float64(completed) / float64(total) * 100
	_ = u.enrollmentRepository.UpdateProgressPercent(ctx, userID, courseID, percent)
}
