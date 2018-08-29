package langmap

import "github.com/jinzhu/gorm"

type Word struct {
	gorm.Model
	Word        string       `gorm:"UNIQUE;NOT NULL"`
	Definitions []Definition `json:",omitempty"`
}
