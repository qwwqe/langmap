package langmap

import "github.com/jinzhu/gorm"

type Definition struct {
	gorm.Model
	User
	Meaning       string         `gorm:"UNIQUE_INDEX:idx_meaning_pronunciation_word"`
	Pronunciation string         `gorm:"UNIQUE_INDEX:idx_meaning_pronunciation_word"`
	WordID        uint           `gorm:"UNIQUE_INDEX:idx_meaning_pronunciation_word" json:",omitempty" sql:"type:integer REFERENCES words(id) ON DELETE RESTRICT ON UPDATE RESTRICT"`
	Word          *Word          `json:",omitempty"`
	Usages        []*Usage       `json:",omitempty"`
	Notes         []*Note        `gorm:"many2many:note_definitions" json:",omitempty"`
	Associations  []*Association `gorm:"many2many:association_definitions" json:",omitempty"`
}
