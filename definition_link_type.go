package langmap

type DefinitionLinkType struct {
	BaseTable
	Name string `json:"name" db:"name"`
}

func (DefinitionLinkType) TableName() string { return "definition_link_types" }
