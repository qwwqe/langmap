package langmap

import "github.com/go-gorp/gorp"

type Collection struct {
	BaseTable
	Name       string `json:"name" db:"name"`
	InstanceId uint   `json:"instance_id" db:"instance_id"`
}

func (Collection) TableName() string { return "collections" }

func LoadCollections(db *gorp.DbMap, f Filter) ([]*Collection, error) {
	r := []*Collection{}

	if _, err := db.Select(&r, SelectQuery(Collection{}, f), f.Values...); err != nil {
		return nil, err
	}

	for _, i := range r {
		i.Preload(db)
	}

	return r, nil
}
