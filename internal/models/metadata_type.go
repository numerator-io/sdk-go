package models

import (
	"encoding/json"
	"os"
)

// Metadata represents the structure of the manifest.json object
type Metadata struct {
	Name            string `json:"name"`
	ReleaseStrategy string `json:"release_strategy"`
	Version         string `json:"version"`
	Language        string `json:"language"`
}

func GetVersionFromMetadataFile(filePath string) (string, error) {
	// Read the content of the Metadata file
	MetadataJSON, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Parse the JSON data into a Metadata object
	var Metadata Metadata
	err = json.Unmarshal(MetadataJSON, &Metadata)
	if err != nil {
		return "", err
	}

	// Return the version
	return Metadata.Version, nil
}
