package langmap

import (
	"io"
	"io/ioutil"

	kagome "github.com/ikawaha/kagome/tokenizer"
)

const (
	JapaneseTokenSymbol = "記号"
)

type JapaneseTokenizerAdapter struct{}

func (JapaneseTokenizerAdapter) IsSemantic(t kagome.Token) bool {
	if t.Class == kagome.UNKNOWN {
		return false
	}

	if t.Pos() == JapaneseTokenSymbol {
		return false
	}

	return true
}

func (a JapaneseTokenizerAdapter) Tokenize(r io.Reader) (*Corpus, error) {
	// really, I just want to be able to pass io.Reader to kagome
	// but it does not support io.Reader.
	// submit a patch?
	text, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	c := &Corpus{
		CorpusWords: make([]*CorpusWord, 0),
	}

	tokenizer := kagome.NewWithDic(kagome.SysDicIPA())
	for position, token := range tokenizer.Tokenize(string(text)) {
		if token.Class == kagome.DUMMY {
			continue
		}

		w := &CorpusWord{
			Semantic: a.IsSemantic(token),
			Position: uint(position),
			Word:     &Word{Word: token.Surface},
		}

		c.CorpusWords = append(c.CorpusWords, w)
	}

	return c, nil
}
