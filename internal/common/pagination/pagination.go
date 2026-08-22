package pagination

import "strconv"

type Query struct {
	Page   int
	Limit  int
	Offset int
}

type Response[T any] struct {
	Items     T   `json:"items"`
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	TotalItem int `json:"totalItem"`
	TotalPage int `json:"totalPage"`
}

func Get(pageStr, limitStr string) Query {
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}
	return Query{
		Page:   page,
		Limit:  limit,
		Offset: (page - 1) * limit,
	}
}
