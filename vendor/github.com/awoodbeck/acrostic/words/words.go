package words

import (
	"bytes"
	"errors"
	"io"

	"github.com/awoodbeck/acrostic/tools"
)

// ErrEmptyBuffer is returned when a given buffer is empty.
var ErrEmptyBuffer = errors.New("empty buffer")

// ErrNilBuffer is returned when a given pointer to a bytes buffer is nil.
var ErrNilBuffer = errors.New("nil buffer")

// ErrNoKeys is returned when a key is requested from an empty words map.
var ErrNoKeys = errors.New("word list does not contain any keys")

// ErrNonexistentKey is returned when a nonexistent key is requested from the words map.
var ErrNonexistentKey = errors.New("word list does not contain the specified key")

// NewWords accepts a pointer to a bytes buffer and uses its contents
// to populate a new Words object before returning its pointer.
func NewWords(buf *bytes.Buffer) (*Words, error) {
	w := &Words{words: make(map[byte][]string)}

	return w, w.CompileWords(buf)
}

// Words maintains a word list, facilitates calculation of
// entropy, and can return a random word upon request.
type Words struct {
	words map[byte][]string
	keys  []byte
}

// CompileWords parses the given bytes buffer and populates
// the Words object.
//
// All words are lower cased and any extraneous spaces are trimmed.
func (w *Words) CompileWords(buf *bytes.Buffer) error {
	var k byte

	switch {
	case buf == nil:
		return ErrNilBuffer
	case buf.Len() == 0:
		return ErrEmptyBuffer
	}

Outer:
	for {
		b, err := buf.ReadBytes('\n')
		switch {
		case err != io.EOF && err != nil:
			return err
		case err == io.EOF && len(b) == 0:
			break Outer
		}

		b = bytes.ToLower(bytes.TrimSpace(b))
		if len(b) == 0 {
			continue
		}

		k = b[0]

		if _, ok := w.words[k]; !ok {
			w.words[k] = []string{string(b)}
			w.keys = append(w.keys, k)
			continue
		}

		w.words[k] = append(w.words[k], string(b))
	}

	return nil
}

// Entropy returns the bits of entropy for "n" number of words.
//
// For example, if the word list was 2048 words long and 4 words were
// chosen at random for a passphrase, the entropy would be 44 bits.
//
// This is only taking into consideration the number of words. Capitalizing
// words, or choosing different separators will increase the entropy on paper.
func (w *Words) Entropy(n int) float64 {
	return tools.CalcEntropy(float64(w.WordCount()), float64(n))
}

// GetRandomKey returns a random key from the map.
func (w *Words) GetRandomKey() (byte, error) {
	length := int64(len(w.keys))
	if length == 0 {
		if len(w.words) == 0 {
			return 0, ErrNoKeys
		}
		for k := range w.words {
			w.keys = append(w.keys, k)
		}
		length = int64(len(w.keys))
	}

	n, err := tools.RandomInt64(length)
	if err != nil {
		return 0, err
	}

	return w.keys[n], nil
}

// GetRandomWord returns a random word from the map for the given key.
func (w *Words) GetRandomWord(key byte) (string, error) {
	words, ok := w.words[key]
	if !ok || len(words) == 0 {
		return "", ErrNonexistentKey
	}

	n, err := tools.RandomInt64(int64(len(words)))
	if err != nil {
		return "", err
	}

	return words[n], err
}

// WordCount returns the total number of words in the current object.
func (w *Words) WordCount() int {
	var c int

	for _, v := range w.words {
		c += len(v)
	}

	return c
}
