package langmap

type Lexica struct {
	BaseTable
	URI        string `json:"uri" db:"uri"`
	Name       string `json:"name" db:"name"`
	LanguageId uint   `json:"language_id" db:"language_id"`
}

func (Lexica) TableName() string { return "lexica" }
