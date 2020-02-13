package models

import (
	"regexp"
	"testing"
)

var (
	wrappeReg = regexp.MustCompile(`^\[.*]$`)
)

func TestUser_WrappedName(t *testing.T) {
	u := User{Name: "hoge"}
	if !wrappeReg.Match([]byte(u.WrappedName())) {
		t.Errorf("Not content wrapped")
	}
}

func TestUser_Validate(t *testing.T) {
	if err := (&User{Age: 25}).Validate(); err != nil {
		t.Errorf("Not expected value 25")
	}
	if err := (&User{Age: 17}).Validate(); err == nil {
		t.Errorf("Expect error 17")
	}
}
