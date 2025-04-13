package postgresql

import (
	"gorm.io/gorm"
)

var (
	NotFound = gorm.ErrRecordNotFound
)
