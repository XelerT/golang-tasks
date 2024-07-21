//go:build !solution

package mycheck

import (
	"bytes"
	"errors"
	"strings"
	"unicode"
)

type slice_err []error

func (errs slice_err) Error() string {
	var buf bytes.Buffer
	for i := range errs {
		buf.WriteString(errs[i].Error())
	}
	return buf.String()
}

func (errs *slice_err) add_err(text string) {
	if len(*errs) != 0 {
		*errs = append(*errs, errors.New(";"+text))
	} else {
		*errs = append(*errs, errors.New(text))
	}
}

func MyCheck(input string) error {
	var errs slice_err

	if strings.ContainsFunc(input, unicode.IsDigit) {
		errs.add_err("found numbers")
	}
	if len(input) >= 20 {
		errs.add_err("line is too long")
	}
	if strings.Count(input, " ") != 2 {
		errs.add_err("no two spaces")
	}

	if len(errs) != 0 {
		return errs
	}

	return nil
}
