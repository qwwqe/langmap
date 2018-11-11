package langmap

import (
	"io"
	"io/util"
)

type TaiwaneseMandarinTokenizerAdapter struct{}

func (a TaiwaneseMandarinTokenizerAdapter) Tokenize(r io.Reader) (*Corpus, error) {
	tokenizer := TaiwaneseMandarinTokenizer{}
	// tokenizer.LoadLexicon("mandarin_lexicon")
	return tokenizer.tokenize(r)
}
