package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/BurntSushi/toml"
)

// Join root dir with arguments, returning a full path,
// This srep is for compatibility in multiple OS.
func CreatePath(pathParts []string) (finalPath string) {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error accessing to filesystem: %v", err)
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
		log.Fatalf("Error accessing to TOML file '%s': %v", tomlPath, err)
	}

	metadata, err := toml.Decode(string(file), data)
	if err != nil {
		log.Fatalf("Error decoding TOML file '%s': %v", tomlPath, err)
	}

	// You can ignore this value, but is a "must-be" or it does not compile
	// Assuming the variable ms is (ms *MyStruct) a pointer
	// Right: ParseToml([]string{"myfolder", "example.toml"}, ms)
	// Wrong: _ := ParseToml([]string{"myfolder", "example.toml"}, ms)
	return metadata
}

func CreateDirIfDoesNotExist(pathToDir string) {
	os.MkdirAll(pathToDir, 0755)
}

// Copy a directory to location.
// Copied folder must not be included in destine path.
// Right: './src/assets' to './output/'
// Wrong: './src/assets' to './output/assets'
// Last bar means copy in the folder and do not replace it
func CopyDir(wg *sync.WaitGroup, dir string, destineLocation string) {
	defer wg.Done()

	command := exec.Command("cp", "-R", dir, destineLocation)

	err := command.Run()
	if err != nil {
		log.Fatalf("Error at copying a directory (source: %s, destine: %s): --%v--.", dir, destineLocation, err)
	}
}

// Simply, copy the file to the path.
func CopyFile(wg *sync.WaitGroup, file string, destineLocation string) {
	defer wg.Done()

	command := exec.Command("cp", file, destineLocation)

	err := command.Run()
	if err != nil {
		log.Fatalf("Error at copying a file (source: %s, destine: %s): --%v--.", file, destineLocation, err)
	}
}
