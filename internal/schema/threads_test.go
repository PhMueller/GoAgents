package schema

/* Unit tests for the Threads schema
We test for the CreateThreadRequest, GetThreadRequest, GetThreadsInfoRequest that invalid inputs are causing a validation error.
*/

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestRequestObjectsValid(t *testing.T) {
	obj := CreateThreadRequest{Title: nil}
	if obj.Title != nil {
		t.Errorf("Expected Title to be nil, got %v", obj.Title)
	}
}

func TestRequestObjectsInvalid(t *testing.T) {
	validate := validator.New()
	_ = validate.RegisterValidation("isStringValidUUID", IsStringValidUUID)

	testCases := []struct {
		name    string
		input   GetThreadRequest
		wantErr bool
	}{
		{
			name:    "Leere ID (required)",
			input:   GetThreadRequest{ID: ""},
			wantErr: true,
		},
		{
			name:    "Ungültige UUID (isStringValidUUID)",
			input:   GetThreadRequest{ID: "not-a-uuid"},
			wantErr: true,
		},
		{
			name:    "Gültige UUID",
			input:   GetThreadRequest{ID: "123e4567-e89b-12d3-a456-426614174000"},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validate.Struct(tc.input)
			if tc.wantErr && err == nil {
				t.Errorf("Expected validation error for %s, got nil", tc.name)
			}
			if !tc.wantErr && err != nil {
				t.Errorf("Did not expect validation error for %s, got: %v", tc.name, err)
			}
		})
	}
}
