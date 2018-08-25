package langmap

import "github.com/jinzhu/gorm"

type Collection struct {
	gorm.Model
	Name        string
	Type        int
	NoteID      int
	PrimaryNote Note    `gorm:"foreignkey:NoteID"`
	Notes       []*Note `gorm:"many2many:note_collections"`
}
