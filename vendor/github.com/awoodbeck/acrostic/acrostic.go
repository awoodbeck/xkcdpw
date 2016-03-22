//go:generate go run -tags=dev assets_generate.go

package acrostic

import (
	"bytes"
	"errors"

	"github.com/awoodbeck/acrostic/assets"
	"github.com/awoodbeck/acrostic/words"
)

// ErrBlankAcrostic is returned when a given acrostic has a length of zero characters.
var ErrBlankAcrostic = errors.New("acrostic may not be blank")

// ErrInvalidNumber is returned when the number of requested acrostics is less than 1.
var ErrInvalidNumber = errors.New("number of returned acrostics less than 1 is not valid")

// ErrUninitialized is returned when acrostics are requested from an uninitalized Acrostic object.
var ErrUninitialized = errors.New("acrostic object has not been initialized")

// NewAcrostic accepts pointers to an adjective and noun word lists, and returns a pointer to
// a populated Acrostic object.
//
// If the adjectives word list pointer is nil, the default adjectives word list will be used.
// Likewise for the nouns word list pointer.
func NewAcrostic(adjs, nouns *words.Words) (*Acrostic, error) {
	var err error

	if adjs == nil {
		adjs, err = getWordList("adjectives.txt")
		if err != nil {
			return nil, err
		}
	}
	if nouns == nil {
		nouns, err = getWordList("nouns.txt")
		if err != nil {
			return nil, err
		}
	}

	acro := &Acrostic{
		adjectives: adjs,
		nouns:      nouns,
	}

	return acro, nil
}

// Acrostic maintains a list of adjectives and nouns, and returns an acrostical
// phrase for a given word.
type Acrostic struct {
	adjectives *words.Words
	nouns      *words.Words
}

// GenerateAcrostics accepts an integer indicating the number of phrases
// to return, and returns a string slice with the results.
func (a *Acrostic) GenerateAcrostics(acro string, num int) ([]string, error) {
	var err error
	var o []string
	var w string
	acroLen := len(acro)

	switch {
	case a.adjectives == nil || a.nouns == nil:
		return o, ErrUninitialized
	case acroLen == 0:
		return o, ErrBlankAcrostic
	case num < 1:
		return o, ErrInvalidNumber
	}

	for i := 0; i < num; i++ {
		var p []byte
		for j := 0; j < acroLen; j++ {
			switch {
			case j == acroLen-1:
				w, err = a.nouns.GetRandomWord(acro[j])
				if err != nil {
					return o, err
				}
				p = append(p, []byte(" "+w)...)
			case j > 0:
				p = append(p, ' ')
				fallthrough
			default:
				w, err = a.nouns.GetRandomWord(acro[j])
				if err != nil {
					return o, err
				}
				p = append(p, []byte(w)...)
			}
		}
		o = append(o, string(p))
	}

	return o, nil
}

// GenerateRandomAcrostics accepts an acrostic length and an integer indicating the number
// of acrostics to return.
func (a *Acrostic) GenerateRandomAcrostics(length, num int) ([]string, error) {
	var o []string
	var err error

	switch {
	case a.adjectives == nil || a.nouns == nil:
		return o, ErrUninitialized
	case length < 1:
		return o, ErrInvalidNumber
	case num < 1:
		return o, ErrInvalidNumber
	}

	acro := make([]byte, length)

	for i := 0; i < length-1; i++ {
		acro[i], err = a.adjectives.GetRandomKey()
		if err != nil {
			return o, err
		}
	}

	acro[length-1], err = a.nouns.GetRandomKey()
	if err != nil {
		return o, err
	}

	return a.GenerateAcrostics(string(acro), num)
}

// GenerateRandomPhrases accepts two integers: words per phrase, and number of phrases.
// It returns a string slice matching the number of phrases.
func (a *Acrostic) GenerateRandomPhrases(words, num int) ([]string, error) {
	var o []string
	switch {
	case a.adjectives == nil || a.nouns == nil:
		return o, ErrUninitialized
	case words < 1:
		return o, ErrInvalidNumber
	case num < 1:
		return o, ErrInvalidNumber
	}

	for i := 0; i < num; i++ {
		s, err := a.GenerateRandomAcrostics(words, 1)
		if err != nil {
			return o, err
		}
		o = append(o, s[0])
	}

	return o, nil
}

// getWordList accepts a file name and returns a pointer to a populated Words object.
func getWordList(name string) (*words.Words, error) {
	w, err := assets.Assets.Open(name)
	if err != nil {
		return nil, err
	}
	buf := &bytes.Buffer{}
	if _, err := buf.ReadFrom(w); err != nil {
		return nil, err
	}
	return words.NewWords(buf)
}
