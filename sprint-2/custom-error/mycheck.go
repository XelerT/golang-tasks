//go:build !solution

package mycheck

import (
	"bytes"
	"errors"
	"strings"
)

type slice_err []error

func (errs slice_err) Error() string {
	var buf bytes.Buffer
	for i, _ := range errs {
		buf.WriteString(errs[i].Error())
	}
	return buf.String()
}

func (errs *slice_err) add_err(text string) {
	if len(*errs) != 0 {
		*errs = append(*errs, errors.New(";" + text))
	} else {
		*errs = append(*errs, errors.New(text))
	}
}

func (errors slice_err) has_errors() bool {
	if len(errors) != 0 { return true }; return false
}

func MyCheck(input string) error {
	var errs slice_err
	is_number := func(r rune) bool {
		return '0' <= r && r <= '9'
	}
	if strings.ContainsFunc(input, is_number) {
		errs.add_err("found numbers")
	}
	if len(input) >= 20 {
		errs.add_err("line is too long")
	}
	if strings.Count(input, " ") != 2 {
		errs.add_err("no two spaces")
	}

	if errs.has_errors() {
		return errs
	}

	return nil
}
