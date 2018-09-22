package langmap

type Wordlist struct {
	BaseTable
	Name       string `db:"name"`
	InstanceId uint   `db:"instance_id"`
}
