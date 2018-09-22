package langmap

type DefinitionLink struct {
	BaseTable
	TypeId        uint `json:"type_id" db:"type_id"`
	Definition1Id uint `json:"definition1_id" db:"definition1_id"`
	Definition2Id uint `json:"definition2_id" db:"definition2_id"`
	InstanceId    uint `json:"instance_id" db:"instance_id"`
}
