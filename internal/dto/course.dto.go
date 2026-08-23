package dto

type CourseCreateReq struct {
	CategoryID  int     `json:"categoryId" binding:"required"`
	Title       string  `json:"title" binding:"required"`
	Slug        string  `json:"slug" binding:"required"`
	Description *string `json:"description"`
	Thumbnail   *string `json:"thumbnail"`
	Price       float64 `json:"price"`
	Level       string  `json:"level" binding:"omitempty,oneof=beginner intermediate advanced"`
	Status      string  `json:"status" binding:"omitempty,oneof=draft published archived"`
}

type CourseUpdateReq struct {
	CategoryID  *int     `json:"categoryId"`
	Title       *string  `json:"title"`
	Slug        *string  `json:"slug"`
	Description *string  `json:"description"`
	Thumbnail   *string  `json:"thumbnail"`
	Price       *float64 `json:"price"`
	Level       *string  `json:"level" binding:"omitempty,oneof=beginner intermediate advanced"`
	Status      *string  `json:"status" binding:"omitempty,oneof=draft published archived"`
}

type CourseUpdateStatusReq struct {
	Status string `json:"status" binding:"required,oneof=draft published archived"`
}

type CoursePublicFilter struct {
	CategoryID *int
	Level      string
	MinPrice   *float64
	MaxPrice   *float64
	Q          string
}
