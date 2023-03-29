package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func GetApiKey() string {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	return getApiKeyWithReaderWriter(reader, writer)
}

func getApiKeyWithReaderWriter(reader io.Reader, writer io.Writer) string {
	file, err := os.Open(".env")
	if err != nil {
		file, _ = os.Create(".env")
	}
	defer file.Close()
	apiKey := ""
	apiKey = getApiKeyFromScanner(file)
	if apiKey == "" {
		apiKey = getApiKeyFromInput(reader, writer, file)
	}
	return apiKey
}

func getApiKeyFromInput(reader io.Reader, writer io.Writer, file *os.File) string {
	fmt.Fprintln(writer, "There is no .env file yet, so you need to supply an OpenAI API key.")
	fmt.Fprintln(writer, "See: https://platform.openai.com/account/api-keys")
	fmt.Fprint(writer, "Enter OpenAI API key: ")
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		apiKey := strings.TrimSpace(scanner.Text())
		if apiKey != "" {
			_, _ = file.WriteString(fmt.Sprintf("OPENAI_API_KEY=%s", apiKey))
			return apiKey
		}
		fmt.Fprintln(writer, "Error: no API key provided")
		fmt.Fprint(writer, "Enter OpenAI API key: ")
	}
	return ""
}

func getApiKeyFromScanner(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		// Only pay attention to appropriate line
		apiKey := strings.TrimPrefix(scanner.Text(), "OPENAI_API_KEY=")
		if apiKey != scanner.Text() {
			return apiKey
		}
	}
	return ""
}
