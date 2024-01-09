package id

import (
	"errors"
	"testing"
)

func TestPadPrefix(t *testing.T) {
	type test struct {
		Name     string
		Prefix   string
		Expected string
	}

	tests := []test{
		{
			Name:     "Pad",
			Prefix:   "PAD",
			Expected: "PAD00",
		},
		{
			Name:     "No Pad",
			Prefix:   "NOPAD",
			Expected: "NOPAD",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(tt *testing.T) {
			pad := padPrefix(test.Prefix)
			if pad != test.Expected {
				tt.Errorf("expected %s but got %s", test.Expected, pad)
			}
		})
	}
}

func TestNew(t *testing.T) {
	_ = RegisterType("PAD", "TEST_PAD")
	id, err := NewFromSequence("PAD", 1)
	if err != nil {
		t.Error(err)
	}
	if valid, err := IsValid(id); !valid || err != nil {
		t.Errorf("expected a valid id but got %v", id)
	}
}

func TestIsValid(t *testing.T) {
	type test struct {
		Name  string
		Id    Id
		Error error
	}

	tests := []test{
		{
			Name:  "Valid",
			Id:    Id("TES00b4069096cc152c66"),
			Error: nil,
		},
		{
			Name:  "Invalid length",
			Id:    Id("TES00b4069096cc152"),
			Error: ErrBadIdLength,
		},
		{
			Name:  "Unregistered Prefix",
			Id:    Id("DURRRb4069096cc152c66"),
			Error: ErrPrefixNotFound,
		},
	}

	if err := RegisterType("TES", "TEST_PREFIX"); err != nil {
		t.Errorf("got error %v", err)
		t.FailNow()
	}

	for _, test := range tests {
		t.Run(test.Name, func(tt *testing.T) {
			if _, err := IsValid(test.Id); !errors.Is(err, test.Error) {
				tt.Errorf("expected error %T but got %T", test.Error, err)
			}
		})
	}
}
