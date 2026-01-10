package schema

/* Unit tests for the Messages schema

We test for the *Request - schema that invalid inputs are causing a validation error.

We do not test the *Response objects as they do not have validation rules.
They are tested indirectly via integration tests and unit test on the handler.
*/

import (
	"testing"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func TestCreateMessageRequest_Validation(t *testing.T) {
	bindingValidate, _ := binding.Validator.Engine().(*validator.Validate)
	_ = bindingValidate.RegisterValidation("isStringValidUUID", IsStringValidUUID)

	testCases := []struct {
		name    string
		input   CreateMessageRequest
		wantErr bool
	}{
		{
			name:    "Empty ThreadID (error: required)",
			input:   CreateMessageRequest{ThreadID: "", Content: "Hello"},
			wantErr: true,
		},
		{
			name:    "Invalid ThreadID (error: not a valid uuid)",
			input:   CreateMessageRequest{ThreadID: "not-a-uuid", Content: "Hello"},
			wantErr: true,
		},
		{
			name:    "Empty Content (error: required)",
			input:   CreateMessageRequest{ThreadID: "123e4567-e89b-12d3-a456-426614174000", Content: ""},
			wantErr: true,
		},
		{
			name:    "Valid Input",
			input:   CreateMessageRequest{ThreadID: "123e4567-e89b-12d3-a456-426614174000", Content: "Hello"},
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

func TestGetMessageRequest_Validation(t *testing.T) {
	bindingValidate, _ := binding.Validator.Engine().(*validator.Validate)
	_ = bindingValidate.RegisterValidation("isStringValidUUID", IsStringValidUUID)

	testCases := []struct {
		name    string
		input   GetMessageRequest
		wantErr bool
	}{
		{
			name:    "Empty ThreadID (error: required)",
			input:   GetMessageRequest{ThreadID: "", MessageID: "123e4567-e89b-12d3-a456-426614174000"},
			wantErr: true,
		},
		{
			name:    "Invalid ThreadID (error: not a valid uuid)",
			input:   GetMessageRequest{ThreadID: "not-a-uuid", MessageID: "123e4567-e89b-12d3-a456-426614174000"},
			wantErr: true,
		},
		{
			name:    "Empty MessageID (error: required)",
			input:   GetMessageRequest{ThreadID: "123e4567-e89b-12d3-a456-426614174000", MessageID: ""},
			wantErr: true,
		},
		{
			name:    "Invalid MessageID (error: not a valid uuid)",
			input:   GetMessageRequest{ThreadID: "123e4567-e89b-12d3-a456-426614174000", MessageID: "not-a-uuid"},
			wantErr: true,
		},
		{
			name:    "Valid Input",
			input:   GetMessageRequest{ThreadID: "123e4567-e89b-12d3-a456-426614174000", MessageID: "123e4567-e89b-12d3-a456-426614174000"},
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

func TestGetMessagesRequest_Validation(t *testing.T) {
	bindingValidate, _ := binding.Validator.Engine().(*validator.Validate)
	_ = bindingValidate.RegisterValidation("isStringValidUUID", IsStringValidUUID)

	testCases := []struct {
		name    string
		input   GetMessagesRequest
		wantErr bool
	}{
		{
			name:    "Empty ThreadID (error: required)",
			input:   GetMessagesRequest{ThreadID: ""},
			wantErr: true,
		},
		{
			name:    "Invalid ThreadID (error: not a valid uuid)",
			input:   GetMessagesRequest{ThreadID: "not-a-uuid"},
			wantErr: true,
		},
		{
			name:    "Valid ThreadID",
			input:   GetMessagesRequest{ThreadID: "123e4567-e89b-12d3-a456-426614174000"},
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
