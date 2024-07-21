//go:build !solution

package testequal

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

func maps_str_equal(lhs, rhs map[string]string) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for key := range lhs {
		if val, ok := rhs[key]; ok {
			if val != rhs[key] {
				fmt.Printf("%s != %s", val, rhs[key])
				return false
			}
			fmt.Println("asdasd")
		} else {
			return false
		}
	}

	return true
}

func slices_equal(lhs, rhs interface{}) bool {
	if reflect.TypeOf(lhs).Kind() != reflect.Slice {
		return false
	}
	rhs_vals := reflect.Indirect(reflect.ValueOf(rhs))
	lhs_vals := reflect.Indirect(reflect.ValueOf(lhs))

	if lhs_vals.Len() != rhs_vals.Len() {
		return false
	}
	if lhs_vals.Cap() != rhs_vals.Cap() {
		return false
	}
	if lhs_vals.Len() == 0 {
		return false
	}
	if lhs_vals.Index(0).Kind() != rhs_vals.Index(0).Kind() {
		return false
	}

	val := reflect.ValueOf(lhs_vals.Index(0))
	if !val.Comparable() {
		return false
	}

	for i := 0; i < rhs_vals.Len(); i++ {
		if lhs_vals.Index(i) != rhs_vals.Index(i) {
			fmt.Printf("%v != %v", lhs_vals.Index(i), rhs_vals.Index(i))
			return false
		}
	}

	return true
}

func equal(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}
	if reflect.TypeOf(expected) != reflect.TypeOf(actual) {
		return false
	}

	switch reflect.ValueOf(expected).Kind() {
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		fallthrough
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		fallthrough
	case reflect.String:
		return expected == actual
	case reflect.Map:
		ex := expected.(map[string]string)
		ac := actual.(map[string]string)
		return maps_str_equal(ex, ac)
	case reflect.Slice:
		return slices_equal(expected, actual)
	default:
		return false
	}
}

func get_msg(msgAndArgs ...interface{}) string {
	n_args := len(msgAndArgs)
	if n_args > 1 {
		return fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	} else if n_args == 1 {
		msg := msgAndArgs[0]
		if str, ok := msgAndArgs[0].(string); ok {
			return str
		}
		return fmt.Sprintf("%+v", msg)
	}
	return ""
}

func get_caller_info() []string {
	var file string
	var line int
	var pc uintptr
	var ok bool

	callers := []string{}
	for i := 0; ; i++ {
		pc, file, line, ok = runtime.Caller(i)
		if !ok {
			break
		}
		foo := runtime.FuncForPC(pc)
		if foo == nil {
			break
		}

		callers = append(callers, fmt.Sprintf("%s:%d", file, line))
	}

	return callers
}

// AssertEqual checks that expected and actual are equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are equal.
func AssertEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	if equal(expected, actual) {
		return true
	}
	t.Errorf("%s %s", get_msg(msgAndArgs...), strings.Join(get_caller_info(), "\n\t\t\t"))

	return false
}

// AssertNotEqual checks that expected and actual are not equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are not equal.
func AssertNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	if !equal(expected, actual) {
		return true
	}
	t.Errorf("%s %s", get_msg(msgAndArgs...), strings.Join(get_caller_info(), "\n\t\t\t"))

	return false
}

// RequireEqual does the same as AssertEqual but fails caller test immediately.
func RequireEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	if !AssertEqual(t, expected, actual, msgAndArgs...) {
		t.FailNow()
	}
}

// RequireNotEqual does the same as AssertNotEqual but fails caller test immediately.
func RequireNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	if !AssertNotEqual(t, expected, actual, msgAndArgs...) {
		t.FailNow()
	}
}
