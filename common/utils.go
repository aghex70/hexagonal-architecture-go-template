package common

import (
	"bytes"
	"fmt"
	"github.com/aghex70/hexagonal-architecture-go-template/templates"
	"github.com/satori/go.uuid"
	"io"
	"os"
	"strings"
	"text/template"
)

func ScanString(prompt string) (string, error) {
	fmt.Printf("%s: ", prompt)
	var value string
	_, err := fmt.Scan(&value)
	if err != nil {
		return "", err
	}
	if value == "" {
		return "", ProvideValueError
	}
	return value, nil
}

func ScanStringWithDefault(prompt, defaultValue string) (string, error) {
	value, err := ScanString(prompt)
	if err != nil {
		return "", err
	}
	if value == "" {
		return defaultValue, nil
	}

	return value, nil
}

func ScanStringCastBoolean(prompt, defaultValue string) (bool, error) {
	value, err := ScanString(prompt)
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

func ScanMultipleStrings(prompt string) ([]string, error) {
	fmt.Printf("%s: ", prompt)
	var values []string
	for {
		var value string
		if _, err := fmt.Scan(&value); err != nil {
			break
		}
		if value == "" {
			break
		}
		values = append(values, value)
	}

	return values, nil
}

func ProjectExists(path string) (bool, error) {
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

func CreateDirectory(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}
	return nil
}

func GenerateProjectPath(pbp, pn string) string {
	projectPath := fmt.Sprintf("%s%s/", pbp, pn)
	return projectPath
}

func GenerateStubs(pp, side string) error {
	baseDirectories := []string{
		DomainDirectory,
		PortsDirectory,
		ServicesDirectory,
		HandlersDirectory,
		RepositoriesDirectory,
	}

	for _, d := range baseDirectories {
		bd := pp + side + d
		err := CreateDirectory(bd)
		if err != nil {
			return err
		}
	}
	return nil
}

func GenerateFile(basePath, side, path, name, t string, data interface{}) error {
	route := basePath + side + path + strings.ToLower(name) + GolangFileExtension
	f, err := os.Create(route)
	if err != nil {
		return err
	}
	defer f.Close()

	tmpl, err := template.New(GenerateUUID()).Parse(t)
	if err != nil {
		panic(err)
	}

	var b bytes.Buffer
	err = tmpl.Execute(&b, data)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, &b)
	if err != nil {
		panic(err)
	}

	return nil
}

func GenerateMultiPartsFile(basePath, side, path, name, ts, te string, data interface{}) error {
	f, err := os.Create(basePath + side + path + strings.ToLower(name) + GolangFileExtension)
	if err != nil {
		return err
	}
	defer f.Close()

	switch v := data.(type) {
	case []templates.DomainData:
		for _, value := range v {
			tmpl, err := template.New(GenerateUUID()).Parse(ts)
			if err != nil {
				panic(err)
			}

			var b bytes.Buffer
			err = tmpl.Execute(&b, value)
			if err != nil {
				panic(err)
			}

			_, err = io.Copy(f, &b)
			if err != nil {
				panic(err)
			}
			break
		}
	default:
		fmt.Println("unexpected type")
	}

	switch v := data.(type) {
	case []templates.DomainData:
		for _, value := range v {
			tmpl, err := template.New(GenerateUUID()).Parse(te)
			if err != nil {
				panic(err)
			}

			var b bytes.Buffer
			err = tmpl.Execute(&b, value)
			if err != nil {
				panic(err)
			}

			_, err = io.Copy(f, &b)
			if err != nil {
				panic(err)
			}
		}
	default:
		fmt.Println("unexpected type")
	}
	return nil
}

func GenerateUUID() string {
	u := uuid.NewV4()
	return u.String()
}
