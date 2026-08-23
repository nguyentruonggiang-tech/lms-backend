package repository_impl

import (
	"context"
	"sort"

	"lms-api/ent"
	"lms-api/ent/enrollments"
	"lms-api/ent/users"
	"lms-api/internal/dto"
	"lms-api/internal/repository"
)

type dashboardRepository struct {
	client *ent.Client
}

func NewDashboardRepository(client *ent.Client) repository.DashboardRepository {
	return &dashboardRepository{client: client}
}

func (r *dashboardRepository) GetOverview(ctx context.Context) (dto.DashboardOverview, error) {
	totalStudents, err := r.client.Users.Query().
		Where(users.RoleEQ(users.RoleStudent), users.DeletedAtIsNil()).
		Count(ctx)
	if err != nil {
		return dto.DashboardOverview{}, err
	}

	totalCourses, err := r.client.Courses.Query().Count(ctx)
	if err != nil {
		return dto.DashboardOverview{}, err
	}

	totalEnrollments, err := r.client.Enrollments.Query().Count(ctx)
	if err != nil {
		return dto.DashboardOverview{}, err
	}

	totalCertificates, err := r.client.Certificates.Query().Count(ctx)
	if err != nil {
		return dto.DashboardOverview{}, err
	}

	return dto.DashboardOverview{
		TotalStudents:     totalStudents,
		TotalCourses:      totalCourses,
		TotalEnrollments:  totalEnrollments,
		TotalCertificates: totalCertificates,
	}, nil
}

type courseEnrollCount struct {
	CourseID int `json:"course_id"`
	Count    int `json:"count"`
}

func (r *dashboardRepository) GetTopCourses(ctx context.Context, filter dto.TopCourseFilter) ([]dto.TopCourseItem, error) {
	q := r.client.Enrollments.Query()
	if filter.FromDate != nil {
		q = q.Where(enrollments.EnrolledAtGTE(*filter.FromDate))
	}
	if filter.ToDate != nil {
		q = q.Where(enrollments.EnrolledAtLTE(*filter.ToDate))
	}

	var counts []courseEnrollCount
	if err := q.GroupBy(enrollments.FieldCourseID).
		Aggregate(ent.Count()).
		Scan(ctx, &counts); err != nil {
		return nil, err
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i].Count > counts[j].Count
	})

	if filter.Limit > len(counts) {
		filter.Limit = len(counts)
	}
	counts = counts[:filter.Limit]

	result := make([]dto.TopCourseItem, 0, len(counts))
	for _, c := range counts {
		course, err := r.client.Courses.Get(ctx, c.CourseID)
		if err != nil {
			continue
		}
		result = append(result, dto.TopCourseItem{
			CourseID:        c.CourseID,
			Title:           course.Title,
			EnrollmentCount: c.Count,
		})
	}

	return result, nil
}
