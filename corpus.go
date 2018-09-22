package langmap

type Corpus struct {
	BaseTable
	Name       string `db:"name"`
	MD5        string `db:"md_5"`
	Browsable  string `db:"browsable"`
	InstanceId uint   `db:"instance_id"`
}
