package langmap

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	Name  string
	Notes []*Note `gorm:"many2many:note_tags"`
}
