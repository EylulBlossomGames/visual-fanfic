package main

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

// It needs to be public, variables too
type HolaMundo struct {
	Hola  string
	Mundo string
}

type Capitals struct {
	SpainCapital string
}

func TestCreatePath(t *testing.T) {
	thisDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error accessing to filesystem: %v", err)
	}

	myPath := filepath.Join(thisDir, "pablo", "documents", "doc1.txt")
	mySlicePath := []string{"pablo", "documents", "doc1.txt"}

	myCreatedPath := CreatePath(mySlicePath)
	if myPath != myCreatedPath {
		t.Errorf("Expected path and path created by CreatePath() are different.\nEXPECTED:\n%v\nCREATED:\n%v", myPath, myCreatedPath)
	}
}

func TestParseToml(t *testing.T) {
	hm := HolaMundo{}
	path := []string{"testdata", "test1.toml"}
	metadata := ParseToml(path, &hm)

	// expected: hola = "Hola", mundo = "Mundo"
	expectedValues := HolaMundo{
		Hola:  "Hola",
		Mundo: "Mundo",
	}
	if hm.Hola != expectedValues.Hola || hm.Mundo != expectedValues.Mundo {
		t.Errorf("Parsed values from TOML are different:\nEXPECTED:\n%v\nVALUES:\n%v\nMETADATA:\n%v", expectedValues, hm, metadata)
	}

	// Wrong case
	c := Capitals{}
	path2 := []string{"testdata", "test2.toml"}
	metadata2 := ParseToml(path2, &c)

	expectedCapitals := Capitals{
		SpainCapital: "Madrid",
	}
	if expectedCapitals.SpainCapital == c.SpainCapital {
		t.Errorf("Parsed data and expected data must be different.\nEXPECTED:\n%v\nVALUES:\n%v\nMETADATA:\n%v", expectedCapitals, c, metadata2)
	}
}
