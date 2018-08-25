package langmap

import "github.com/jinzhu/gorm"

type Association struct {
	gorm.Model
	Type        int
	Definitions []*Definition `gorm:"many2many:association_definitions"`
}
