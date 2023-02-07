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
		HandlersDirectory,
		CommandDirectory,
		HandlersDirectory,
		MigrationsDirectory,
		ConfigDirectory,
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

func GenerateFile(path, extension string, data []templates.FileConfiguration) error {
	f, err := os.Create(path + extension)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, d := range data {
		if d.Repeat {
			for _, entity := range d.RepeatEntities {
				templateContext := d.TemplateContext
				if templateContext.Entity == "" {
					templateContext.Entity = entity
					templateContext.LowerEntity = strings.ToLower(entity)
					templateContext.Initial = string(templateContext.LowerEntity[0])
				}

				err = executeTemplate(f, d.Template, templateContext)
				if err != nil {
					return err
				}
			}
		} else {
			err = executeTemplate(f, d.Template, d.TemplateContext)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func executeTemplate(w io.Writer, ts string, data interface{}) error {
	tmpl, err := template.New(GenerateUUID()).Parse(ts)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, &buf)
	return err
}

func GenerateUUID() string {
	u := uuid.NewV4()
	return u.String()
}
