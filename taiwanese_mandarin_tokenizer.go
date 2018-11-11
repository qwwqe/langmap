package langmap

import (
	"bufio"
	"github.com/derekparker/trie"
	"io"
	"io/ioutil"
	"strings"
	"unicode/utf8"
)

const (
	SentenceDelimiter  = "。"
	MaxDepth           = 3
	SimpleSegmentation = false
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
// TODO (felix): test Trie creation speed
// TODO (felix): fix sentence identification (requires intelligent parsing of quotations)
type TaiwaneseMandarinTokenizer struct {
	lexicon *Trie
}

type wordSequence struct {
	words []string
}

func (seq wordSequence) numWords() {
	return len(seq.words)
}

func (seq wordSequence) numBytes() {
	count := 0
	for _, word := range seq.wordSequence {
		count += len(word)
	}
}

func (seq wordSequence) addWord(w string) {

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

// (t TaiwaneseMandarinTokenizer) Tokenize tokenizes a text using the currently loaded lexicon
func (t TaiwaneseMandarinTokenizer) Tokenize(r io.Reader) (*Corpus, error) {
	corpus := &Corpus{
		CorpusWords: make([]CorpusWord, 0),
	}

	// Dump reader contexts to a bytestring (backwards approach...?)
	text, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	// Check for UTF-8 validity
	if !utf8.Valid(text) {
		return nil, err
	}

	textIndex := 0

	// This loop represents cycles of MaxDepth segmentations.
	// At the beginning of each iteration of this loop,
	// one and only one of the following is true:
	// a) r points to the beginning of a lexical sequence (the start of an n-word sequence)
	// b) r points to the beginning of a non-lexical sequence (a 'leader')
	// c) r points to nothing (break)
	for textIndex < len(text) {
		// Consume leading non-lexical characters
		leaderIndex := textIndex
		for w := 0; leaderIndex < len(text); leaderIndex += w {
			runeValue, width := utf8.DecodeRune(text[leaderIndex:])
			_, lexical := t.lexicon.Find(string(runeValue))
			if lexical {
				break
			}
			w = width
		}

		// Concatenate leading non-lexical characters into a single string
		if leaderIndex != textIndex {
			word := CorpusWord{
				Semantic: false,
				Position: len(corpus.corpusWords),
				Word:     text[index:leaderIndex],
			}
			corpus.CorpusWords = append(corpus.CorpusWords, word)
			textIndex = leaderIndex
		}

		// Compile n-depth candidates
		candidates := make([][]string, 1)
		for len(candidates) > 0 && len(candidates[0]) < MaxDepth {
			candidate := candidates[0]
			candidates = candidates[1:] // safe to due declared capacity of 1

			// Sequence offset from textIndex
			seqOffset := 0
			for _, word := range candidate {
				seqOffset += len(word)
			}

			// Compile candidates for following word
			var leader string
			leaderIndex := textIndex + seqOffset
			for w := 0; leaderIndex < len(text); leaderIndex += w {
				runeValue, width := utf8.DecodeRune(text[textIndex+seqOffset:])
				leader = text[textIndex+seqOffset : leaderIndex+width]
				_, lexical := t.lexicon.Find(leader)
				if !lexical {
					break
				}
				candidates = append(candidates, append(candidate, leader))
				w = width
			}
		}

		// Filter by greatest total length
		var filtered [][]string
		maxLength := -1
		for _, candidate := range candidates {
			length := 0
			for _, word := range candidate {
				length += len(word)
			}

			if length > maxLength {
				maxLength = length
				filtered = [][]string{candidate}
			} else if length == maxLength {
				filtered = append(filtered, candidate)
			}
		}
		candidates = filtered

		// Filter by greatest average word length
		maxAverage := -1.0
		for _, candidate := range candidates {
			average := float64(maxLength) / float64(len(candidate))

			if average > maxAverage {
				maxAverage = average
				filtered = [][]string{candidate}
			} else if average == maxAverage {
				filtered = append(filtered, candidate)
			}
		}
		candidates = filtered

		// Filter by smallest variance of word lengths
		leastVariance := -1.0
		for _, candidate := range candidates {
			squaredDifferenceSum := 0
			for _, word := range candidate {
				squaredDifferenceSum += math.Pow(len(word)-maxAverage, 2)
			}
			variance := squaredDifferenceSum / float64(len(candidate))
			if variance < leastVariance || leastVariance < 0 {
				leastVariance = variance
				filtered = [][]string{candidate}
			} else if variance == leastVariance {
				filtered = append(filtered, candidate)
			}
		}
		candidates = filtered

		// Filter by largest sum of morphemic freedom of single-character words
		maxFreedom := -1.0
		for _, candidate := range candidates {
			freedom := 0
			for _, word := range candidate {
				f, _ = t.lexicon.Find(word)
				freedom += f
			}

			// This is the last heuristic, don't select multiple candidates
			if freedom > maxFreedom {
				maxFreedom = freedom
				filtered = [][]string{candidate}
			}
		}
		candidates = filtered

		// Add words to corpus. If no candidates exist... we're at the end of the text?
		if len(candidates) < 1 || len(candidates[0]) < 1 {
			break
		}

		for _, word := range candidates[0] {
			corpusWord := CorpusWord{
				Semantic: true,
				Position: len(corpus.corpusWords),
				Word:     word,
			}
			corpus.CorpusWords = append(corpus.CorpusWords, corpusWord)
			textIndex += len(word)
		}

	}

	return corpus, nil

}
