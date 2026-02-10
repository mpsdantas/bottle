package postgresql

import (
	"gorm.io/gorm"
)

var (
	ErrNotFound = gorm.ErrRecordNotFound
)
