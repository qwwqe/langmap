package langmap

import "github.com/go-gorp/gorp"

type BaseTable struct {
	Id uint `json:"id" db:"id"`
}

func (t BaseTable) GetId() uint              { return t.Id }
func (BaseTable) TableName() string          { return "" }
func (*BaseTable) Preload(*gorp.DbMap) error { return nil }
