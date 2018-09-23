package langmap

import (
	"bufio"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenizerService struct {
	BaseService
}

func (s *TokenizerService) Create(c *gin.Context) {
	type Data struct {
		LanguageId uint `form:"language_id"`
	}

	r := &Data{}

	if err := c.ShouldBind(r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": ErrJsonFailed,
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	log.Printf("%+v", r)

	file, err := c.FormFile("text")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "Failed to receive file",
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": "Failed to open received file",
			"errors": NewErrorsJSON([]error{err}),
		})
		return
	}

	t := JapaneseTokenizer{}
	scan := bufio.NewScanner(f)
	count := 0

	for scan.Scan() {
		tokens := t.Tokenize(scan.Text())
		count += len(tokens)
	}

	log.Println("parsed", count, "tokens")

	c.Status(http.StatusNoContent)
}

func (s *TokenizerService) Register() {
	ServiceRegister(s)
}
