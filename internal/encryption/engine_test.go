package encryption

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type engineTestInput struct {
	key       string
	inputData string
}

type engineTestOutput struct {
	outputData string
	err        error
}

func TestEncryption(t *testing.T) {
	validEncryptionKey := func() string {
		key := make([]byte, KeySize)
		for idx, _ := range key {
			key[idx] = '0'
		}
		return string(key)
	}()

	referenceData := "test data"

	tests := []struct {
		testName string
		input    engineTestInput
		output   engineTestOutput
	}{
		{
			testName: "valid encryption scenario",
			input: engineTestInput{
				key:       validEncryptionKey,
				inputData: referenceData,
			},
			output: engineTestOutput{
				outputData: "MukBbMZrwJ6gQMij9Lix63Gednt47vIwwT/4qqoWO7+ztL1ifA==",
				err:        nil,
			},
		},
		{
			testName: "invalid key length",
			input: engineTestInput{
				key:       "00",
				inputData: referenceData,
			},
			output: engineTestOutput{
				outputData: "",
				err:        ErrCipherCreation,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			encryptedData, encryptionErr := encrypt(test.input.inputData, test.input.key)
			assert.Equal(t, len(test.output.outputData), len(encryptedData))
			assert.Equal(t, test.output.err, encryptionErr)
			if encryptionErr != nil {
				assert.ErrorIs(t, test.output.err, encryptionErr)
			}
		})
	}
}

func TestDecryption(t *testing.T) {
	validEncryptionKey := func() string {
		key := make([]byte, KeySize)
		for idx, _ := range key {
			key[idx] = '0'
		}
		return string(key)
	}()

	referenceData := "test data"

	tests := []struct {
		testName string
		input    engineTestInput
		output   engineTestOutput
	}{
		{
			testName: "valid decryption scenario",
			input: engineTestInput{
				key:       validEncryptionKey,
				inputData: "MukBbMZrwJ6gQMij9Lix63Gednt47vIwwT/4qqoWO7+ztL1ifA==",
			},
			output: engineTestOutput{
				outputData: referenceData,
				err:        nil,
			},
		},
		{
			testName: "wrong date encoding",
			input: engineTestInput{
				key:       validEncryptionKey,
				inputData: "MDA==",
			},
			output: engineTestOutput{
				outputData: "",
				err:        ErrDataDecoding,
			},
		},
		{
			testName: "invalid key ",
			input: engineTestInput{
				key:       "00",
				inputData: "MDA=",
			},
			output: engineTestOutput{
				outputData: "",
				err:        ErrCipherCreation,
			},
		},
		{
			testName: "failed decryption",
			input: engineTestInput{
				key:       validEncryptionKey,
				inputData: "YVDp2yjfG9qwSuhwh+zMKeEZm/C5MZ6flTJCzDvdDM4HYXAO4g==",
			},
			output: engineTestOutput{
				outputData: "",
				err:        ErrFailedDecryption,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			decryptedData, decryptionErr := decrypt(test.input.inputData, test.input.key)

			assert.Equal(t, test.output.outputData, decryptedData)
			assert.Equal(t, test.output.err, decryptionErr)
			if decryptionErr != nil {
				assert.ErrorIs(t, test.output.err, decryptionErr)
			}
		})
	}
}
