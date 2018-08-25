package langmap

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	Comment string
}
