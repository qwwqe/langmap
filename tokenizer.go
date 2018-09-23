package langmap

import (
	"github.com/ikawaha/kagome/tokenizer"
)

type JapaneseTokenizer struct {
}

func (JapaneseTokenizer) Tokenize(input string) []tokenizer.Token {
	t := tokenizer.NewWithDic(tokenizer.SysDicIPA())
	return t.Analyze(input, tokenizer.Normal)
}
