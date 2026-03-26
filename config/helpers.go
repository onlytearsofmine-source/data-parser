package helpers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ReadJSONFile reads a JSON file and unmarshals it into the given interface.
func ReadJSONFile(filePath string, v interface{}) error {
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(fileData, v); err != nil {
		return err
	}

	return nil
}

// WriteJSONFile writes the given data as JSON to the specified file.
func WriteJSONFile(filePath string, data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, jsonData, 0644)
}

// FileExists checks if a file exists and is not a directory.
func FileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return !info.IsDir()
}

// GetFileExtension returns the extension of the file.
func GetFileExtension(filePath string) string {
	return filepath.Ext(filePath)
}