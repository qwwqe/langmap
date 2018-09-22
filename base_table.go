package langmap

type BaseTable struct {
	Id uint `json:"id" db:"id"`
}

func (t *BaseTable) GetId() uint { return t.Id }
