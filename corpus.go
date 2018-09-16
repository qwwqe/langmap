package langmap

type Corpus struct {
	Id         uint   `db:"id"`
	Name       string `db:"name"`
	MD5        string `db:"md_5"`
	Browsable  string `db:"browsable"`
	InstanceId uint   `db:"instance_id"`
}
