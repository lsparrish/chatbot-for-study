package utils

import (
    "bufio"
    "log"
    "os"
    "strings"
)

func GetApiKey() string {
    file, err := os.Open(".env")
    if err == nil {
        defer file.Close()
        return getApiKeyFromScanner(bufio.NewScanner(file))
    } else {
        log.Print(err)
    }

    log.Print("Enter OpenAI API key: ")
    apiKey := getApiKeyFromScanner(bufio.NewScanner(os.Stdin))
    if apiKey != "" {
        if err := os.WriteFile(".env", []byte("OPENAI_API_KEY="+apiKey), 0600); err != nil {
            log.Fatal(err)
        }
    }
    return apiKey
    
}

func getApiKeyFromScanner(scanner *bufio.Scanner) string {
    for scanner.Scan() {
        if apiKey := strings.TrimPrefix(scanner.Text(), "OPENAI_API_KEY="); apiKey != scanner.Text() {
            return apiKey
        }
    }
    return ""
}
