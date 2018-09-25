package langmap

import (
	"bufio"
	"github.com/derekparker/trie"
	"io"
	"strings"
)

// TaiwaneseMandarinTokenizer describes a type that can be used in conjunction with an arbitrary lexicon to tokenize arbitrary text.
// This being the case, however, it is designed specifically with Traditional Taiwanese Mandarin in mind,
// and therefore the heuristics employed in the tokenization process will
// likely have little efficacy if applied to other languages.
//
// The general tokenization algorithm, as well as the various heuristics
// employed in disambiguating tokenizations of apparently equal likelihood,
// are adopted from MMSEG, written by 蔡志浩 (Chih-Hao Tsai). Both are
// described below.
//
// Tokenization is chiefly accomplished via a 3-Depth Maximum-Matching
// Algorithm, the basic axiom of which is that given a string of characters
// beginning at N, the most likely tokenization is the longest 3-word
// sequence beginning at N. This depth is not fixed, and can be adjusted
// as the user pleases, keeping in mind that deeper is not necessarily
// better.
//
// When multiple 3-word sequences of equal length are found, various
// heuristics are then employed to resolve disambiguation. These are as
// implemented as follows and applied in the order described.
//
// 1) Greatest average word length
// 2) Smallest variance of word lengths
// 3) Largest sum of morphemic freedom of single-character words.
//    Morphemic freedom here is approximated by frequency count.
//
// TODO (felix): put this in a separate package?
type TaiwaneseMandarinTokenizer struct {
	lexicon *Trie
}

// (t TaiwaneseMandarinTokenizer) LoadLexicon builds a lexicon Trie from a stream of <word, frequency> pairs
func (t TaiwaneseMandarinTokenizer) LoadLexicon(s *Scanner) error {
	if t.lexicon == nil {
		t.lexicon = Trie.New()
	}

	for s.Scan() {
		if splits := strings.Fields(s.Text()); len(splits) != 2 {
			continue
		}

		t.lexicon.Add(splits[0], splits[1])
	}

	return nil
}
