package langmap

type Definition struct {
	Id            uint   `json:"id" db:"id"`
	Pronunciation string `json:"pronunciation" db:"pronunciation"`
	Meaning       string `json:"meaning" db:"meaning"`
	WordId        uint   `json:"word_id" db:"word_id"`
	InstanceId    uint   `json:"instance_id" db:"instance_id"`

	Word string `json:"word" db:"-"`
}
