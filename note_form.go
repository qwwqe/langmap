package langmap

type NoteCreateForm struct {
	Title   string   `json:"title"`
	Type    int      `json:"type"`
	Comment []string `json:"comment"`
	Tags    []string `json:"tags"`
}
