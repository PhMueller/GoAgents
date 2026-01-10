package schema

/* Unit tests for the Threads schema

We test for the *Request - schema that invalid inputs are causing a validation error.

We do not test the *Response objects as they do not have validation rules.
They are tested indirectly via integration tests and unit test on the handler.
*/

import (
	"testing"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func TestCreateThreadRequest_Validation(t *testing.T) {
	bindingValidate, _ := binding.Validator.Engine().(*validator.Validate)
	// CreateThreadRequest hat keine required- oder UUID-Felder, nur optionale Title
	testCases := []struct {
		name    string
		input   CreateThreadRequest
		wantErr bool
	}{
		{
			name:    "Empty Title (no error)",
			input:   CreateThreadRequest{},
			wantErr: false,
		},
		{
			name:    "Nil Title (no error)",
			input:   CreateThreadRequest{Title: nil},
			wantErr: false,
		},
		{
			name:    "Non-empty Title (no error)",
			input:   CreateThreadRequest{Title: ptrString("")},
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := bindingValidate.Struct(tc.input)
			if tc.wantErr && err == nil {
				t.Errorf("Expected validation error for %s, got nil", tc.name)
			}
			if !tc.wantErr && err != nil {
				t.Errorf("Did not expect validation error for %s, got: %v", tc.name, err)
			}
		})
	}
}

func TestGetThreadRequest_Validation(t *testing.T) {
	bindingValidate, _ := binding.Validator.Engine().(*validator.Validate)
	_ = bindingValidate.RegisterValidation("isStringValidUUID", IsStringValidUUID)

	testCases := []struct {
		name    string
		input   GetThreadRequest
		wantErr bool
	}{
		{
			name:    "Empty ID (error: required)",
			input:   GetThreadRequest{ID: ""},
			wantErr: true,
		},
		{
			name:    "Invalid ID (error: not a valid uuid)",
			input:   GetThreadRequest{ID: "not-a-uuid"},
			wantErr: true,
		},
		{
			name:    "Valid ID",
			input:   GetThreadRequest{ID: "123e4567-e89b-12d3-a456-426614174000"},
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := bindingValidate.Struct(tc.input)
			if tc.wantErr && err == nil {
				t.Errorf("Expected validation error for %s, got nil", tc.name)
			}
			if !tc.wantErr && err != nil {
				t.Errorf("Did not expect validation error for %s, got: %v", tc.name, err)
			}
		})
	}
}

func TestGetThreadsInfoRequest_Validation(t *testing.T) {
	bindingValidate, _ := binding.Validator.Engine().(*validator.Validate)
	intValue := 0

	testCases := []struct {
		name    string
		input   GetThreadsInfoRequest
		wantErr bool
	}{
		{
			name:    "Empty struct (no error)",
			input:   GetThreadsInfoRequest{},
			wantErr: false,
		},
		{
			name:    "Nil Cursor and Size (no error)",
			input:   GetThreadsInfoRequest{Cursor: nil, Size: nil},
			wantErr: false,
		},
		{
			name:    "Non-empty Cursor and Size (no error)",
			input:   GetThreadsInfoRequest{Cursor: ptrString(""), Size: &intValue},
			wantErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := bindingValidate.Struct(tc.input)
			if tc.wantErr && err == nil {
				t.Errorf("Expected validation error for %s, got nil", tc.name)
			}
			if !tc.wantErr && err != nil {
				t.Errorf("Did not expect validation error for %s, got: %v", tc.name, err)
			}
		})
	}
}

func ptrString(s string) *string { return &s }
