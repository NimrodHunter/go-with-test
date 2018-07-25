package main

type Dictionary map[string]string

const (
	ErrCouldNotFoundTheWord = DictionaryErr("could not find the word you were looking for")
	ErrWordAlreadyExist     = DictionaryErr("word already exist in the dictionary")
	ErrWordDoesNotExist     = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrCouldNotFoundTheWord
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrCouldNotFoundTheWord:
		d[word] = definition
	case nil:
		return ErrWordAlreadyExist
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrCouldNotFoundTheWord:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
