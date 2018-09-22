package langmap

type DefinitionLinkType struct {
	BaseTable
	Name string `json:"name" db:"name"`
}

func (_ DefinitionLinkType) TableName() string { return "definition_link_types" }
