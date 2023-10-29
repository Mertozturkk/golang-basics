package v3

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word] // bu mapin onemli bir ozelligi. ikinci deger basarili olan keylerde boolean deger doner

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, text string) {
	d[word] = text
}
