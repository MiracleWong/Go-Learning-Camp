package main

import (
"errors"
"testing"
)

type errorString string

func (es *errorString) Error() string {
	return string(*es)
}

func New(str string) error {
	ss := errorString(str)
	return &ss
}

var ErrName = New("EOF")
var ErrStruct = errors.New("EOF")

func TestErrorString(t *testing.T) {
	if ErrName == New("EOF") {
		t.Log("ErrName")
	}
	if ErrStruct == errors.New("EOF") {
		t.Log("ErrStruct")
	}
}
