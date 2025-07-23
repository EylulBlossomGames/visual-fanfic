package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// Join root dir with arguments, returning a full path,
// This srep is for compatibility in multiple OS.
func CreatePath(pathParts []string) (finalPath string) {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error accesing to filesystem: %v", err)
	}

	fullPath := append([]string{currentDir}, pathParts...)
	finalPath = filepath.Join(fullPath...)

	return finalPath
}

// Takes a path to TOML file and a data struct
// and fill the struct with TOML content
func ParseToml(pathParts []string, data interface{}) toml.MetaData {
	tomlPath := CreatePath(pathParts)

	file, err := os.ReadFile(tomlPath)
	if err != nil {
		log.Fatalf("Error accesing to TOML file: %v", err)
	}

	metadata, err := toml.Decode(string(file), data)
	if err != nil {
		log.Fatalf("Error decoding TOML file: %v", err)
	}

	// You can ignore this value, but is a "must-be" or it does not compile
	// Assuming the variable ms is (ms *MyStruct) a pointer
	// Right: ParseToml([]string{"myfolder", "example.toml"}, ms)
	// Wrong: _ := ParseToml([]string{"myfolder", "example.toml"}, ms)
	return metadata
}
