package langmap

type Note struct {
	BaseTable
	Title      string `db:"title" json:"title"`
	Comment    string `db:"comment" json:"comment"`
	InstanceId uint   `db:"instance_id" json:"instance_id"`
}

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
