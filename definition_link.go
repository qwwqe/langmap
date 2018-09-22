package langmap

type DefinitionLink struct {
	BaseTable
	TypeId        uint `db:"type_id"`
	Definition1Id uint `db:"definition1_id"`
	Definition2Id uint `db:"definition2_id"`
	InstanceId    uint `db:"instance_id"`
}
