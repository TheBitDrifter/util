package testing_util

import (
	"testing"
)

// TestError struct and error implementation remains unchanged
type TestError struct {
	Message string
}

func (e TestError) Error() string {
	return e.Message
}

func errorFunc() error {
	return TestError{Message: "an error occurred"}
}

func noErrorFunc() error {
	return nil
}

// Custom T to capture errors
type CapturingT struct {
	*testing.T
}

func (c *CapturingT) Errorf(format string, args ...interface{}) {}

func TestCheckError(t *testing.T) {
	tests := []struct {
		name          string
		function      func() error
		expected      []error
		expectSuccess bool
	}{
		{
			name:          "Matching error",
			function:      errorFunc,
			expected:      []error{TestError{Message: "an error occurred"}},
			expectSuccess: true,
		},
		{
			name:          "No error",
			function:      noErrorFunc,
			expected:      []error{TestError{Message: "an error occurred"}},
			expectSuccess: false,
		},
		{
			name:          "Non-matching error",
			function:      errorFunc,
			expected:      []error{TestError{Message: "different error that does not match func"}},
			expectSuccess: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			capturingT := &CapturingT{T: t}
			actualError := tt.function()
			success := CheckError(capturingT, tt.function, actualError, tt.expected...)

			if success != tt.expectSuccess {
				t.Errorf("Expected success %v, got %v", tt.expectSuccess, success)
			}
		})
	}
}
