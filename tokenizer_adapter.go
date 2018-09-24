package langmap

import (
	"io"
)

type TokenizerAdapter interface {
	Tokenize(io.Reader) (*Corpus, error)
}
