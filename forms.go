package langmap

type WordForm struct {
	Word        string `json:"word"`
	Definitions []struct {
		Pronunciation string `json:"pronunciation"`
		Meaning       string `json:"meaning"`
	} `json:"definitions"`
}

type NoteForm struct {
	Title   string   `json:"title"`
	Type    int      `json:"type"`
	Comment []string `json:"comment"`
	Tags    []string `json:"tags"`
}
