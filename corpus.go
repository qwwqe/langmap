package langmap

type Corpus struct {
	BaseTable
	Name       string `json:"name" db:"name"`
	MD5        string `json:"md5" db:"md5"`
	Browsable  string `json:"browsable" db:"browsable"`
	InstanceId uint   `json:"instance_id" db:"instance_id"`
}
