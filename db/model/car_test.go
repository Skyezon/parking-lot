package model

import (
	"testing"

	"github.com/skyezon/parking-lot/common/errors"
)

func Test_NewCar(t *testing.T) {
	tests := []struct {
		name        string
		color       string
		regisNumber string
		isErr       bool
	}{
		{
			name:        "valid",
			color:       "Red",
			regisNumber: "bk-1234-har",
			isErr:       false,
		},
		{
			name:        "error color",
			color:       "",
			regisNumber: "bk-1234-har",
			isErr:       true,
		},
		{
			name:        "error regis number",
			color:       "Red",
			regisNumber: "wado",
			isErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewCar(tt.color, tt.regisNumber)
			if (err != nil) != tt.isErr {
				t.Errorf(errors.UNIT_TEST_ERR_TEMPLATE, tt.isErr, err)
			}
		})
	}

}

func Test_ValidateRegisNumber(t *testing.T) {
	tests := []struct {
		name        string
		regisNumber string
		isErr       bool
	}{
		{"Valid 1", "AB-1234-CDE", false},
		{"Valid 2", "D-8080-YUK", false},
		{"invalid format 1", "DD-BOBO-YUK", true},
		{"Empty", "", true},
		{"Invalid Format", "ABC-1234-DEF-GHI", true},
		{"Invalid Length", "AB-12345-CDE", true},
		{"Invalid Digit", "AB-1234-123", true},
		{"Incomplete Part", "AB-1234-", true},
		{"Missing Part", "-1234-CDE", true},
		{"Invalid Part", "AB-12-XYZ", true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validateRegisNumber(tc.regisNumber)
			if (err != nil) != tc.isErr {
				t.Errorf(errors.UNIT_TEST_ERR_TEMPLATE,  tc.isErr, err)
			}
		})
	}

}
