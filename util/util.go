package util

import "github.com/pkg/errors"

func Something() (string,error) {
	return "", errors.New("foo")
}
