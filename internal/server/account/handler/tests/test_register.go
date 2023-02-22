package tests

import (
	"fmt"
	"github.com/aligang/Gophkeeper/internal/common/account"
	"github.com/aligang/Gophkeeper/internal/server/repository/inmemory"
	"testing"
)

type registerTestInput struct {
	msg           *account.RegisterRequest
	memoryContent *inmemory.Repository
}

type registerTestOutput struct {
	msg    *account.RegisterResponse
	status error
}

func TestRegister(t *testing.T) {
	tests := []struct {
		name     string
		input    registerTestInput
		expected registerTestOutput
	}{
		{
			name: "CORRECT_REGISTRATION",
			input: registerTestInput{
				msg:           nil,
				memoryContent: nil,
			},
			expected: registerTestOutput{
				msg:    nil,
				status: nil,
			},
		},
	}

	for _, test := range tests {
		fmt.Println(test.name)
	}
}
