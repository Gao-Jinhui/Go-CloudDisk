package utils

import "github.com/pkg/errors"

var (
	ErrWrongPassword  = errors.New("wrong password")
	ErrEmptyName      = errors.New("empty name")
	ErrEmptyPhone     = errors.New("empty phone")
	ErrEmptyEmail     = errors.New("empty email")
	ErrRecordNotFound = errors.New("record not found")
)
