package langmap

type Note struct {
	BaseTable
	Title      string `json:"title" db:"title"`
	Comment    string `json:"comment" db:"comment"`
	InstanceId uint   `json:"instance_id" db:"instance_id"`
}

func (Note) TableName() string { return "notes" }

func (n *Note) FromMap(m map[string]interface{}) {
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
