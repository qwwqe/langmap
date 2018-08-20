package langmap

import "github.com/jinzhu/gorm"

type Word struct {
	gorm.Model
	Word        string `gorm:"UNIQUE;NOT NULL"`
	Definitions []Definition
}

type Definition struct {
	gorm.Model
	Meaning       string
	Pronunciation string
	WordID        int
	Word          Word
	Notes         []*Note        `gorm:"many2many:note_definitions"`
	Associations  []*Association `gorm:"many2many:association_defintions"`
}

type Comment struct {
	gorm.Model
	Comment string
}

type Note struct {
	gorm.Model
	Title       string
	Type        int
	Comments    []Comment
	Definitions []*Definition `gorm:"many2many:note_definitions"`
	Tags        []*Tag        `gorm:"many2many:note_tags"`
	Collections []*Collection `gorm:"many2many:note_collections"`
}

type Association struct {
	gorm.Model
	Type        int
	Definitions []*Definition `gorm:"many2many:association_definitions"`
}

type Tag struct {
	gorm.Model
	Name  string
	Notes []*Note `gorm:"many2many:note_tags"`
}

type Collection struct {
	gorm.Model
	Name        string
	Type        int
	NoteID      int
	PrimaryNote Note    `gorm:"foreignkey:NoteID"`
	Notes       []*Note `gorm:"many2many:note_collections"`
}
