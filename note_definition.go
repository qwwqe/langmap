package langmap

type NoteDefinition struct {
	BaseTable
	NoteId       uint `db:"note_id"`
	DefinitionId uint `db:"definition_id"`
	InstanceId   uint `db:"instance_id"`
}
