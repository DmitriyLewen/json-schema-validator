package main

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"os"
)

func main() {
	schemaFilePath := fmt.Sprintf("file://%s", os.Args[1])
	docFilePath := fmt.Sprintf("file://%s", os.Args[2])

	if err := Validate(schemaFilePath, docFilePath); err != nil {
		fmt.Println(err)
	}

}

func Validate(schemaFilePath, docFilePath string) error {
	schemaLoader := gojsonschema.NewReferenceLoader(schemaFilePath)
	documentLoader := gojsonschema.NewReferenceLoader(docFilePath)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return fmt.Errorf("failed to validate: %w", err)
	}

	if result.Valid() {
		return nil
	}

	errSummary := fmt.Sprintf("encountered %d validation errors:", len(result.Errors()))
	for _, verr := range result.Errors() {
		errSummary += fmt.Sprintf("\n  - %s", verr.String())
	}

	return fmt.Errorf(errSummary)
}
