package dto

import "time"

type DashboardOverview struct {
	TotalStudents     int `json:"totalStudents"`
	TotalCourses      int `json:"totalCourses"`
	TotalEnrollments  int `json:"totalEnrollments"`
	TotalCertificates int `json:"totalCertificates"`
}

type TopCourseReq struct {
	FromDate string `form:"fromDate"`
	ToDate   string `form:"toDate"`
	Limit    string `form:"limit"`
}

type TopCourseFilter struct {
	FromDate *time.Time
	ToDate   *time.Time
	Limit    int
}

type TopCourseItem struct {
	CourseID        int    `json:"courseId"`
	Title           string `json:"title"`
	EnrollmentCount int    `json:"enrollmentCount"`
}
