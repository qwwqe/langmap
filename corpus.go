package langmap

import "github.com/go-gorp/gorp"

type Corpus struct {
	BaseTable
	Name        string        `json:"name" db:"name"`
	MD5         string        `json:"md5" db:"md5"`
	Browsable   string        `json:"browsable" db:"browsable"`
	InstanceId  uint          `json:"instance_id" db:"instance_id"`
	CorpusWords []*CorpusWord `json:"corpus_words,omitempty" db:"-"`
}

func (Corpus) TableName() string { return "corpora" }

func (i *Corpus) FromMap(m map[string]interface{}) {
	for k, v := range m {
		switch k {
		case "id":
			i.Id = uint(v.(float64))

		case "name":
			i.Name = v.(string)

		case "md5":
			i.MD5 = v.(string)

		case "browsable":
			i.Browsable = v.(string)

		case "instance_id":
			i.InstanceId = uint(v.(float64))

		}
	}
}

func (t *Corpus) PostInsert(s gorp.SqlExecutor) error {
	for _, w := range t.CorpusWords {
		w.CorpusId = t.Id

		if err := s.Insert(w); err != nil {
			return err
		}
	}

	return nil
}
