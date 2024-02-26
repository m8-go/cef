package main

import (
	"bytes"
	"embed"
	"encoding/csv"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"unicode"

	"go.m8.ru/cef/internal/gen/naming"
)

//go:embed _template/*
var templateDir embed.FS

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}

func run() error {
	t, err := template.New("templates").
		Funcs(map[string]any{
			"typeMapping":   typeMapping,
			"exported":      exported,
			"unexported":    unexported,
			"add":           add,
			"renameKeyWord": renameKeyWord,
			"pretty":        naming.Pretty,
			"quote":         strconv.Quote,
		}).
		ParseFS(templateDir, "_template/*.tmpl")
	if err != nil {
		return err
	}

	if err := executeHeaderField(t); err != nil {
		return err
	}

	if err := executeExtensionField(t); err != nil {
		return err
	}

	if err := executeMarshalText(t); err != nil {
		return err
	}

	if err := executeUnmarshalText(t); err != nil {
		return err
	}

	return nil
}

type headerField struct {
	FieldName   string
	Typ         string
	Size        string
	Description string
}

func executeHeaderField(t *template.Template) error {
	records, err := csvRecords("header_field.csv")
	if err != nil {
		return err
	}

	var fields []headerField

	for i, v := range records {
		if i == 0 {
			continue
		}

		fields = append(fields, headerField{
			FieldName:   v[0],
			Typ:         v[1],
			Size:        v[2],
			Description: v[3],
		})
	}

	return execute(t, "header_field.tmpl", map[string]any{
		"Fields":    fields,
		"FieldsNum": len(fields),
	}, "header_field_gen.go")
}

type extensionField struct {
	CEFSpecificationVersion string
	CEFKeyName              string
	FullName                string
	DataType                string
	Length                  string
	Meaning                 string
}

func csvRecords(fileName string) ([][]string, error) {
	p, err := filepath.Abs("internal/gen/" + fileName)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func executeExtensionField(t *template.Template) error {
	var fields []extensionField

	records, err := csvRecords("extension_field.csv")
	if err != nil {
		return err
	}

	for i, v := range records {
		if i == 0 {
			continue
		}

		fields = append(fields, extensionField{
			CEFSpecificationVersion: v[0],
			CEFKeyName:              v[1],
			FullName:                v[2],
			DataType:                v[3],
			Length:                  v[4],
			Meaning:                 v[5],
		})
	}

	return execute(t, "extension_field.tmpl", fields, "extension_field_gen.go")
}

func executeMarshalText(t *template.Template) error {
	records, err := csvRecords("extension_field.csv")
	if err != nil {
		return err
	}

	var extensionFields []extensionField

	for i, v := range records {
		if i == 0 {
			continue
		}

		extensionFields = append(extensionFields, extensionField{
			CEFSpecificationVersion: v[0],
			CEFKeyName:              v[1],
			FullName:                v[2],
			DataType:                v[3],
			Length:                  v[4],
			Meaning:                 v[5],
		})
	}

	return execute(t, "marshal_text.tmpl", extensionFields, "marshal_text_gen.go")
}

func executeUnmarshalText(t *template.Template) error {
	records, err := csvRecords("header_field.csv")
	if err != nil {
		return err
	}

	var headerFields []headerField

	for i, v := range records {
		if i == 0 {
			continue
		}

		headerFields = append(headerFields, headerField{
			FieldName:   v[0],
			Typ:         v[1],
			Size:        v[2],
			Description: v[3],
		})
	}

	records, err = csvRecords("extension_field.csv")
	if err != nil {
		return err
	}

	var extensionFields []extensionField

	for i, v := range records {
		if i == 0 {
			continue
		}

		extensionFields = append(extensionFields, extensionField{
			CEFSpecificationVersion: v[0],
			CEFKeyName:              v[1],
			FullName:                v[2],
			DataType:                v[3],
			Length:                  v[4],
			Meaning:                 v[5],
		})
	}

	return execute(t, "unmarshal_text.tmpl", map[string]any{
		"HeaderFields":    headerFields,
		"HeaderFieldsNum": len(headerFields),
		"ExtensionFields": extensionFields,
	}, "unmarshal_text_gen.go")
}

func execute(t *template.Template, templateName string, data any, fileName string) error {
	var b bytes.Buffer

	if err := t.ExecuteTemplate(&b, templateName, data); err != nil {
		return err
	}

	sourceData, err := format.Source(b.Bytes())
	if err != nil {
		return err
	}

	if err := os.WriteFile(fileName, sourceData, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func typeMapping(typ string) string { return _typeMapping[typ] }

var _typeMapping = map[string]string{
	"Numeric":                  "int",
	"AgentSeverityEnumeration": "AgentSeverity",
	"String":                   "string",
	"IP Address":               "net.IP",
	"IPv6 address":             "net.IP",
	"IPv6 Address":             "net.IP",
	"Floating Point":           "float32",
	"Long":                     "int64",
	"Integer":                  "int",
	"Time Stamp":               "string",
	"MAC Address":              "net.HardwareAddr",
	"Double":                   "float64",
}

func exported(str string) string {
	if unicode.IsUpper(rune(str[0])) {
		return str
	}

	var b strings.Builder

	for i, c := range str {
		if i == 0 {
			b.WriteRune(unicode.ToUpper(c))

			continue
		}

		b.WriteRune(c)
	}

	return b.String()
}

func unexported(str string) string {
	var b strings.Builder

	parts := naming.Split(str)

	for i, part := range parts {
		s, ok := naming.Rule(part)
		if ok {
			if i == 0 {
				s = strings.ToLower(s)
			}
			b.WriteString(s)
		} else {
			if i == 0 {
				part = strings.ToLower(part)
			}
			b.WriteString(part)
		}
	}

	return b.String()
}

func add(a, b int) int {
	return a + b
}

func renameKeyWord(s string) string {
	switch s {
	case "type":
		return "typ"

	default:
		return s
	}
}
