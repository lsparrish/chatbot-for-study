package utils

import (
    "log"
    "os"
    "strings"
    "bufio"
)

func GetApiKey() string {
    file, err := os.Open(".env")
    if err != nil {
        log.Fatal("Error opening .env file")
    }
    defer file.Close()

    variables := make(map[string]string)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
		pair := strings.Split(line, "=")
		if len(pair) < 2 {
    		continue // skip this line
		}
		variables[pair[0]] = pair[1]
    }

    apiKey := variables["OPENAI_API_KEY"]
    return apiKey
}
