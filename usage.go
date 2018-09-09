package langmap

import "github.com/jinzhu/gorm"

type Usage struct {
	gorm.Model
	User
	Sentence     string `gorm:"UNIQUE_INDEX:idx_sentence_definition"`
	Positions    string
	DefinitionID uint        `gorm:"UNIQUE_INDEX:idx_sentence_definition" json:",omitempty" sql:"type:integer REFERENCES definitions(id) ON DELETE RESTRICT ON UPDATE RESTRICT"`
	Definition   *Definition `json:",omitempty"`
}
