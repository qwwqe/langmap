package langmap

type NoteTag struct {
	BaseTable
	NoteId     uint `json:"note_id" db:"note_id"`
	TagId      uint `json:"tag_id" db:"tag_id"`
	InstanceId uint `json:"instance_id" db:"instance_id"`
}

func (NoteTag) TableName() string { return "note_tags" }
