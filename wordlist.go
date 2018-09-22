package langmap

type Wordlist struct {
	BaseTable
	Name       string `json:"name" db:"name"`
	InstanceId uint   `json:"instance_id" db:"instance_id"`
}

func (_ Wordlist) TableName() string { return "wordlists" }
