package dto

type EnrollReq struct {
	CourseID int `json:"courseId" binding:"required"`
}
