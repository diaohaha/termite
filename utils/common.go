package utils

import (
	"github.com/pingcap/errors"
	"unicode"
)

func inArray(need interface{}, needArr []interface{}) bool {
	for _, v := range needArr {
		if need == v {
			return true
		}
	}
	return false
}

func CheckNaminConventions(name string) error {
	err := errors.New("命名错误")
	if name == "" {
		return err
	}
	runeStr := []rune(name)
	for _, v := range runeStr {
		if unicode.IsDigit(v) {
			continue
		}
		if unicode.IsLetter(v) {
			continue
		}
		if v == rune('_') {
			continue
		}
		return err
	}
	return nil
}
