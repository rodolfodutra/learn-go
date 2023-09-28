package main

type Dictionary map[string]string

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("word already exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (dict Dictionary) Search(word string) (string, error) {
	definition, ok := dict[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (dict Dictionary) Add(word, definition string) error {

	_, err := dict.Search(word)

	switch err {
	case ErrNotFound:
		dict[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (dict Dictionary) Update(word, definition string) error {
	_, err := dict.Search(word)

	switch err {
	case ErrNotFound:
		return ErrNotFound
	case nil:
		dict[word] = definition
	default:
		return err
	}

	return nil
}

func (dict Dictionary) Delete(word string) error {
	_, err := dict.Search(word)

	switch err {
	case ErrNotFound:
		return ErrNotFound
	case nil:
		delete(dict, word)
	default:
		return err
	}

	return nil
}
