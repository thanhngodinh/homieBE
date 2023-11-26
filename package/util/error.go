package util

import "errors"

type ErrorDefined error

var ErrorDuplicatePhone ErrorDefined = errors.New("duplicate phone")
var ErrorDuplicateEmail ErrorDefined = errors.New("duplicate email")

var ErrorWrongTypePageIdx ErrorDefined = errors.New("pageIdx must be an integer")
var ErrorWrongTypePageSize ErrorDefined = errors.New("pageSize must be an integer")

var ErrorIdNotMatch ErrorDefined = errors.New("id param not match with body")
var ErrorCodeNotMatch ErrorDefined = errors.New("code param not match with body")

var ErrorIdEmpty ErrorDefined = errors.New("id is empty")
var ErrorCodeEmpty ErrorDefined = errors.New("code is empty")
var ErrorPostIdEmpty ErrorDefined = errors.New("post id is empty")
var ErrorUserIdEmpty ErrorDefined = errors.New("user id is empty")

const ErrorDuplicateKey = `ERROR: duplicate key value violates unique constraint "user_like_posts_pkey" (SQLSTATE 23505)`

func IsDefinedErrorType(err error) bool {
	switch err.(type) {
	case ErrorDefined:
		return true
	default:
		return false
	}
}
