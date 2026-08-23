package dto

type LessonCreateReq struct {
	Title           string  `json:"title" binding:"required"`
	Content         *string `json:"content" binding:"omitempty"`
	VideoURL        *string `json:"video_url" binding:"omitempty"`
	DurationMinutes *int    `json:"duration_minutes" binding:"omitempty"`
	SortOrder       *int    `json:"sort_order" binding:"omitempty"`
	IsPreview       *bool   `json:"is_preview" binding:"omitempty"`
}

type LessonUpdateReq struct {
	Title           *string `json:"title" binding:"omitempty"`
	Content         *string `json:"content" binding:"omitempty"`
	VideoURL        *string `json:"video_url" binding:"omitempty"`
	DurationMinutes *int    `json:"duration_minutes" binding:"omitempty"`
	SortOrder       *int    `json:"sort_order" binding:"omitempty"`
	IsPreview       *bool   `json:"is_preview" binding:"omitempty"`
}
