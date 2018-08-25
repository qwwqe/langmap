package langmap

import "github.com/jinzhu/gorm"

type Note struct {
	gorm.Model
	Title       string
	Type        int
	Comments    []Comment
	Definitions []*Definition `gorm:"many2many:note_definitions"`
	Tags        []*Tag        `gorm:"many2many:note_tags"`
	Collections []*Collection `gorm:"many2many:note_collections"`
}
