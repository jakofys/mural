package mural

import (
	"errors"
)

func ErrMustBeNumeral(found string) error {
	return errors.New("mural: must be numeral, found: '" + found + "'")
}

func ErrAlreadyExistIdentifier(id string) error {
	return errors.New("mural: lexer identifier declare twice: '" + id + "'")
}

func ErrUnexistedIdentifier(id string) error {
	return errors.New("mural: lexer unidentified identifier: '" + id + "'")
}
