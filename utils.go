package hexagonal_architecture_go_template

import (
	"fmt"
	"strings"
)

func scanString(prompt string) (string, error) {
	fmt.Print(fmt.Printf("%s: ", prompt))
	var value string
	_, err := fmt.Scan(&value)
	if err != nil {
		return "", err
	}
	if value == "" {
		return "", ProvideValueError
	}
	fmt.Println()
	return value, nil
}

func scanStringWithDefault(prompt, defaultValue string) (string, error) {
	value, err := scanString(prompt)
	if err != nil {
		return "", err
	}
	if value == "" {
		return defaultValue, nil
	}

	return value, nil
}

func scanStringCastBoolean(prompt, defaultValue string) (bool, error) {
	value, err := scanString(prompt)
	if err != nil {
		return false, err
	}

	switch strings.ToLower(value) {
	case "":
		return defaultValue == DEFAULT_ACTIVE, nil
	case DEFAULT_ACTIVE:
		return true, nil
	default:
		return false, nil
	}
}

func scanMultipleStrings(prompt string) ([]string, error) {
	fmt.Print(fmt.Printf("%s: ", prompt))
	var value string
	var values []string
	for {
		if _, err := fmt.Scan(&value); err != nil {
			break
		}
		values = append(values, value)
	}

	fmt.Println()
	return values, nil
}
