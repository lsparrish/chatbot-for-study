package utils

import (
	"bytes"
	"os"
	"io"
	"testing"
)
func TestGetApiKeyWithReaderWriter(t *testing.T) {
	testCases := []struct {
		name           string
		input          io.Reader
		output         io.Writer
		filename       string
		expectedApiKey string
	}{
		{
			name:           "Test with empty file",
			input:          bytes.NewBufferString("testapikey\n"),
			output:         new(bytes.Buffer),
			filename:       "",
			expectedApiKey: "testapikey",
		},
		{
			name:           "Test with non-empty file",
			input:          bytes.NewBuffer(nil),
			output:         new(bytes.Buffer),
			filename:       ".env",
			expectedApiKey: "existingapikey",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.filename != "" {
				os.WriteFile(".env", []byte("OPENAI_API_KEY=existingapikey"), 0666)
				defer os.Remove(".env")
			} else {
				os.Remove(".env")
			}

			apiKey := getApiKeyWithReaderWriter(tc.input, tc.output)

			if apiKey != tc.expectedApiKey {
				t.Errorf("Expected API key %q, but got %q", tc.expectedApiKey, apiKey)
			}
		})
	}
}
