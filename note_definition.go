package langmap

type NoteDefinition struct {
	BaseTable
	NoteId       uint `json:"note_id" db:"note_id"`
	DefinitionId uint `json:"definition_id" db:"definition_id"`
	InstanceId   uint `json:"instance_id" db:"instance_id"`
}
