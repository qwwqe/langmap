package langmap

import "github.com/go-gorp/gorp"

type Note struct {
	BaseTable
	Title      string `json:"title" db:"title"`
	Comment    string `json:"comment" db:"comment"`
	InstanceId uint   `json:"instance_id" db:"instance_id"`
}

func (Note) TableName() string { return "notes" }

func LoadNotes(db *gorp.DbMap) ([]*Note, error) {
	r := []*Note{}
	q := "select * from " + Note{}.TableName()

	if _, err := db.Select(&r, q); err != nil {
		return nil, err
	}

	for _, i := range r {
		i.Preload(db)
	}

	return r, nil
}

func (n *Note) Inject(m map[string]interface{}) {
	for k, v := range m {
		switch k {
		case "id":
			n.Id = uint(v.(float64))

		case "title":
			n.Title = v.(string)

		case "comment":
			n.Comment = v.(string)

		case "instance_id":
			n.InstanceId = uint(v.(float64))

		}
	}
}
