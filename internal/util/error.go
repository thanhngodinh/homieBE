package util

import "errors"

type ErrorDefined error

var ErrorDuplicatePhone ErrorDefined = errors.New("duplicate phone")
var ErrorDuplicateEmail ErrorDefined = errors.New("duplicate email")

var ErrorWrongTypePageIdx ErrorDefined = errors.New("pageIdx must be an integer")
var ErrorWrongTypePageSize ErrorDefined = errors.New("pageSize must be an integer")

var ErrorIdNotMatch ErrorDefined = errors.New("pageIdx must be an integer")
var ErrorCodeNotMatch ErrorDefined = errors.New("pageSize must be an integer")

var ErrorIdEmpty ErrorDefined = errors.New("id is empty")
var ErrorCodeEmpty ErrorDefined = errors.New("code is empty")

func IsDefinedErrorType(err error) bool {
	switch err.(type) {
	case ErrorDefined:
		return true
	default:
		return false
	}
}
