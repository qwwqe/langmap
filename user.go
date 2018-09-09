package langmap

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type User struct {
	LanguageCode string
	UserID       uint
}

func UserLanguage(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	u := c.Request.Header.Get("User-Id")
	l := c.Request.Header.Get("Language-Code")

	return func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", u).Where("language_code = ?", l)
	}
}
