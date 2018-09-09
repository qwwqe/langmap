package langmap

import (
	"github.com/jinzhu/gorm"
)

type Word struct {
	gorm.Model
	User
	Word        string       `gorm:"NOT NULL"`
	Definitions []Definition `json:",omitempty"`
}

func (w *Word) BeforeDelete(t *gorm.DB) error {
	db := t.Where("word_id = ?", w.ID).Delete(&Definition{})

	if db.Error != nil {
		return db.Error
	}

	return nil
}
