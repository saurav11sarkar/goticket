package common

import "testing"

func TestNewPagination(t *testing.T) {
	allowedSortFields := map[string]string{
		"email":     "email",
		"createdAt": "created_at",
	}

	tests := []struct {
		name                      string
		page, limit               int
		sortBy, sortOrder         string
		wantPage, wantLimit       int
		wantOffset                int
		wantSortBy, wantSortOrder string
	}{
		{
			name:          "defaults",
			wantPage:      1,
			wantLimit:     10,
			wantSortBy:    "created_at",
			wantSortOrder: "desc",
		},
		{
			name:          "valid values",
			page:          2,
			limit:         20,
			sortBy:        "email",
			sortOrder:     "ASC",
			wantPage:      2,
			wantLimit:     20,
			wantOffset:    20,
			wantSortBy:    "email",
			wantSortOrder: "asc",
		},
		{
			name:          "caps limit and rejects unsafe sort",
			page:          maxPage + 1,
			limit:         1000,
			sortBy:        "password",
			sortOrder:     "drop table users",
			wantPage:      maxPage,
			wantLimit:     100,
			wantOffset:    (maxPage - 1) * 100,
			wantSortBy:    "created_at",
			wantSortOrder: "desc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewPagination(
				tt.page,
				tt.limit,
				tt.sortBy,
				tt.sortOrder,
				allowedSortFields,
			)

			if got.Page != tt.wantPage || got.Limit != tt.wantLimit || got.Offset != tt.wantOffset {
				t.Errorf("pagination = page %d, limit %d, offset %d; want page %d, limit %d, offset %d",
					got.Page, got.Limit, got.Offset, tt.wantPage, tt.wantLimit, tt.wantOffset)
			}
			if got.SortBy != tt.wantSortBy || got.SortOrder != tt.wantSortOrder {
				t.Errorf("sort = %s %s; want %s %s",
					got.SortBy, got.SortOrder, tt.wantSortBy, tt.wantSortOrder)
			}
		})
	}
}

func TestNewMeta(t *testing.T) {
	meta := NewMeta(57, Pagination{Page: 2, Limit: 10})

	if meta.Total != 57 || meta.Page != 2 || meta.Limit != 10 || meta.TotalPages != 6 {
		t.Fatalf("NewMeta() = %+v, want total=57 page=2 limit=10 totalPages=6", meta)
	}
}
