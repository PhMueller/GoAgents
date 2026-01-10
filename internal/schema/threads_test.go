package schema

/* Unit tests for the Threads schema
We test for the CreateThreadRequest, GetThreadRequest, GetThreadsInfoRequest that invalid inputs are causing a validation error.
*/

import (
	"testing"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func TestRequestObjectsValidateUUID(t *testing.T) {
	/* Test all Thread Schemas that have a requested UUID. */
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
		t.Run(
			tc.name,
			func(t *testing.T) {
				err := bindingValidate.Struct(tc.input)
				if tc.wantErr && err == nil {
					t.Errorf("Expected validation error for %s, got nil", tc.name)
				}
				if !tc.wantErr && err != nil {
					t.Errorf("Did not expect validation error for %s, got: %v", tc.name, err)
				}
			},
		)
	}
}

func TestRequestObjectsValidateRequired(t *testing.T) {
	/* Test all Thread Schemas that have a required field. */

	// do not register the isStringValidUUID validator here, to only test the 'required' tag
	bindingValidate, _ := binding.Validator.Engine().(*validator.Validate)
	emptyString := ""
	intValue := 0

	testCases := []struct {
		name    string
		input   interface{}
		wantErr bool
	}{
		{
			name:    "GetThreadRequest (error: required ID)",
			input:   GetThreadRequest{ID: ""},
			wantErr: true,
		},

		{
			name:    "CreateThreadRequest (no error)",
			input:   CreateThreadRequest{},
			wantErr: false,
		},

		{
			name:    "CreateThreadRequest (no error)",
			input:   CreateThreadRequest{Title: nil},
			wantErr: false,
		},

		{
			name:    "CreateThreadRequest (no error)",
			input:   CreateThreadRequest{Title: &emptyString},
			wantErr: false,
		},

		{
			name:    "GetThreadsInfoRequest (no error)",
			input:   GetThreadsInfoRequest{},
			wantErr: false,
		},

		{
			name:    "GetThreadsInfoRequest (no error)",
			input:   GetThreadsInfoRequest{Cursor: nil, Size: nil},
			wantErr: false,
		},

		{
			name:    "GetThreadsInfoRequest (no error)",
			input:   GetThreadsInfoRequest{Cursor: &emptyString, Size: &intValue},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(
			tc.name,
			func(t *testing.T) {
				err := bindingValidate.Struct(tc.input)
				if tc.wantErr && err == nil {
					t.Errorf("Expected validation error for %s, got nil", tc.name)
				}
				if !tc.wantErr && err != nil {
					t.Errorf("Did not expect validation error for %s, got: %v", tc.name, err)
				}
			},
		)
	}
}
