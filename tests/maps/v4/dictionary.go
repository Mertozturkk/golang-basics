package v4

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word] // bu mapin onemli bir ozelligi. ikinci deger basarili olan keylerde boolean deger doner

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, text string) error {
	_, err := d.Search(d[word])
	switch err {
	case ErrNotFound:
		d[word] = text
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
