package tests

import (
	"fmt"
	"testing"
)

func TestAuthenticate(t *testing.T) {

	tests := []struct {
		name     string
		input    struct{}
		expected struct{}
	}{}

	for _, test := range tests {
		fmt.Println(test.name)
	}

}
