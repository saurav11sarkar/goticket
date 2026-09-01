package common

import "strings"

const (
	defaultPage  = 1
	defaultLimit = 10
	maxLimit     = 100
	maxPage      = 10_000
)

type Pagination struct {
	Page      int
	Limit     int
	Offset    int
	SortBy    string
	SortOrder string
}

type Meta struct {
	Total      int64 `json:"total"`
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalPages int   `json:"totalPages"`
}

func NewPagination(
	page int,
	limit int,
	sortBy string,
	sortOrder string,
	allowedSortFields map[string]string,
) Pagination {
	if page < 1 {
		page = defaultPage
	}
	if page > maxPage {
		page = maxPage
	}
	if limit < 1 {
		limit = defaultLimit
	}
	if limit > maxLimit {
		limit = maxLimit
	}

	column, ok := allowedSortFields[sortBy]
	if !ok {
		column = "created_at"
	}

	sortOrder = strings.ToLower(strings.TrimSpace(sortOrder))
	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "desc"
	}

	return Pagination{
		Page:      page,
		Limit:     limit,
		Offset:    (page - 1) * limit,
		SortBy:    column,
		SortOrder: sortOrder,
	}
}

func NewMeta(total int64, pagination Pagination) Meta {
	if pagination.Page < 1 {
		pagination.Page = defaultPage
	}
	if pagination.Limit < 1 {
		pagination.Limit = defaultLimit
	}

	totalPages := 0
	if total > 0 {
		totalPages = int((total + int64(pagination.Limit) - 1) / int64(pagination.Limit))
	}

	return Meta{
		Total:      total,
		Page:       pagination.Page,
		Limit:      pagination.Limit,
		TotalPages: totalPages,
	}
}
