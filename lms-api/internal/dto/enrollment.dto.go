package dto

type EnrollReq struct {
	CourseID int `json:"courseId" binding:"required"`
}

type AdminEnrollmentFilter struct {
	CourseID *int
	UserID   *int
	Status   string
}

type EnrollmentUpdateStatusReq struct {
	Status string `json:"status" binding:"required,oneof=active completed cancelled"`
}
