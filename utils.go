package hexagonal_architecture_go_template

import (
	"fmt"
	"os"
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
		return defaultValue == DefaultActive, nil
	case DefaultActive:
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

func projectExists(path string) (bool, error) {
	f, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	if f.IsDir() {
		return true, nil
	} else {
		return true, ProjectFileError
	}
}

func createProject(path string) error {
	err := os.Mkdir(path, 0755)
	if err != nil {
		return err
	}
	return nil
}

func generateProjectPath(pbp, pn string) string {
	projectPath := fmt.Sprintf("%s/%s/", pbp, pn)
	return projectPath
}
