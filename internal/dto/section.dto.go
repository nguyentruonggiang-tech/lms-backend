package dto

type SectionCreateReq struct {
	Title     string `json:"title" binding:"required"`
	SortOrder *int   `json:"sort_order" binding:"omitempty"`
}

type SectionUpdateReq struct {
	Title     *string `json:"title" binding:"omitempty"`
	SortOrder *int    `json:"sort_order" binding:"omitempty"`
}

