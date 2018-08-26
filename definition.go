package langmap

import "github.com/jinzhu/gorm"

type Definition struct {
	gorm.Model
	Meaning       string
	Pronunciation string
	WordID        int            `json:",omitempty"`
	Word          *Word          `json:",omitempty"`
	Notes         []*Note        `gorm:"many2many:note_definitions" json:",omitempty"`
	Associations  []*Association `gorm:"many2many:association_definitions" json:",omitempty"`
}
