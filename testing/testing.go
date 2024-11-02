package testing_util

import (
	"errors"
	"reflect"
	"runtime"
)

type tFace interface {
	Errorf(format string, args ...any)
}

func FuncName(fnc any, t tFace) string {
	fv := reflect.ValueOf(fnc)
	if fv.Kind() != reflect.Func {
		t.Errorf("Provided argument is not a function")
		return ""
	}
	return runtime.FuncForPC(fv.Pointer()).Name()
}

func CheckError(t tFace, fnc any, e error, expected ...error) bool {
	if e == nil {
		return false
	}
	funcName := FuncName(fnc, t)

	var filteredExpected []error
	for _, err := range expected {
		if err != nil {
			filteredExpected = append(filteredExpected, err)
		}
	}
	if len(filteredExpected) == 0 {
		t.Errorf("%s, unexpected error: %v", funcName, e)
		return false
	}
	for _, expErr := range filteredExpected {
		if errors.Is(e, expErr) {
			return true
		}
	}
	t.Errorf("%s, expected one of %v, but got %v", funcName, filteredExpected, e)
	return false
}
