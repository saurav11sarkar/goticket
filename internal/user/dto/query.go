package dto

type UserQuery struct {
	SearchTerm string `query:"searchTerm"`
	Name       string `query:"name"`
	Email      string `query:"email"`
	Page       int    `query:"page"`
	Limit      int    `query:"limit"`
	SortBy     string `query:"sortBy"`
	SortOrder  string `query:"sortOrder"`
}
