package limits

import (
	"testing"
)

// Test that a struct with valid lengths does not generate an error
func TestLimitsCheckValid(t *testing.T) {
	t.Parallel()

	type TestType struct {
		Username string `json:"username" max:"30"`
		Email    string `json:"email" max:"64"`
		Password []byte `json:"password" max:"128"`
	}

	test := TestType{
		Username: "ecnepsnai",
		Email:    "ian@ecn.io",
		Password: []byte("hunter2"),
	}

	if err := Check(&test); err != nil {
		t.Errorf("%s", err.Error())
	}
}

// Test that a struct with a string that is too long generates an error
func TestLimitsCheckInvalidStringMaximum(t *testing.T) {
	t.Parallel()

	type TestType struct {
		Username string `json:"username" max:"30"`
		Email    string `json:"email" max:"64"`
		Password []byte `json:"password" max:"128"`
	}

	test := TestType{
		Username: "this is a really long username, why is it so long? who would want a username this long?",
		Email:    "ian@ecn.io",
		Password: []byte("hunter2"),
	}

	if err := Check(&test); err == nil {
		t.Errorf("No error seen when one expected")
	}
}

// Test that a struct with a slice that has too many elements generates an error
func TestLimitsCheckInvalidSliceMaximum(t *testing.T) {
	t.Parallel()

	type TestType struct {
		Username string `json:"username" max:"30"`
		Email    string `json:"email" max:"64"`
		Password []byte `json:"password" max:"30"`
	}

	test := TestType{
		Username: "ecnepsnai",
		Email:    "ian@ecn.io",
		Password: []byte("this is a really long password, why is it so long? who would want a password this long?"),
	}

	if err := Check(&test); err == nil {
		t.Errorf("No error seen when one expected")
	}
}

// Test that a struct with a string that is too long generates an error
func TestLimitsCheckInvalidStringMinimum(t *testing.T) {
	t.Parallel()

	type TestType struct {
		Username string `json:"username" max:"30"`
		Email    string `json:"email" min:"12"`
		Password []byte `json:"password" max:"128"`
	}

	test := TestType{
		Username: "this is a really long username, why is it so long? who would want a username this long?",
		Password: []byte("hunter2"),
	}

	if err := Check(&test); err == nil {
		t.Errorf("No error seen when one expected")
	}
}

// Test that a struct with a slice that has too many elements generates an error
func TestLimitsCheckInvalidSliceMinimum(t *testing.T) {
	t.Parallel()

	type TestType struct {
		Username string `json:"username" max:"30"`
		Email    string `json:"email" max:"64"`
		Password []byte `json:"password" min:"12"`
	}

	test := TestType{
		Username: "ecnepsnai",
		Email:    "ian@ecn.io",
	}

	if err := Check(&test); err == nil {
		t.Errorf("No error seen when one expected")
	}
}

// Test that you can also test a non-pointer
func TestLimitsCheckValidNotPointer(t *testing.T) {
	t.Parallel()

	type TestType struct {
		Username string `json:"username" max:"30"`
		Email    string `json:"email" max:"64"`
		Password []byte `json:"password" max:"128"`
	}

	test := TestType{
		Username: "ecnepsnai",
		Email:    "ian@ecn.io",
		Password: []byte("hunter2"),
	}

	if err := Check(test); err != nil {
		t.Errorf("%s", err.Error())
	}
}

func TestLimitsStructNoTag(t *testing.T) {
	t.Parallel()

	type TestType struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password []byte `json:"password"`
	}

	test := TestType{
		Username: "ecnepsnai",
		Email:    "ian@ecn.io",
		Password: []byte("hunter2"),
	}

	if err := Check(test); err != nil {
		t.Errorf("%s", err.Error())
	}
}

func TestLimitsEmptyStruct(t *testing.T) {
	t.Parallel()

	type TestType struct{}

	test := TestType{}

	if err := Check(test); err != nil {
		t.Errorf("%s", err.Error())
	}
}
