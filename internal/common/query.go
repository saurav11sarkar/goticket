package common

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func SearchScope(searchTerm string, fields ...string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		searchTerm = strings.TrimSpace(searchTerm)
		if searchTerm == "" || len(fields) == 0 {
			return db
		}

		conditions := make([]string, 0, len(fields))
		args := make([]any, 0, len(fields))
		for _, field := range fields {
			// Field names must always be supplied by server code, never by the client.
			conditions = append(conditions, fmt.Sprintf("%s ILIKE ?", field))
			args = append(args, "%"+searchTerm+"%")
		}

		return db.Where("("+strings.Join(conditions, " OR ")+")", args...)
	}
}
