package langmap

type WordCreateForm struct {
	Word        string `json:"word"`
	Definitions []struct {
		Pronunciation string `json:"pronunciation"`
		Meaning       string `json:"meaning"`
	} `json:"definitions"`
}
