package langmap

import "github.com/jinzhu/gorm"

type Definition struct {
	gorm.Model
	Meaning       string
	Pronunciation string
	WordID        int
	Word          Word
	Notes         []*Note        `gorm:"many2many:note_definitions"`
	Associations  []*Association `gorm:"many2many:association_definitions"`
}
