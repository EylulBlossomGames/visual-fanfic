package main

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Getting template
	templatePath := CreatePath([]string{"src", "index.html"})
	tmpl, err := template.New("Test").ParseFiles(templatePath)
	if err != nil {
		log.Fatalf("Error creating template: %v", err)
	}

	// Creating output folder if does not exist
	outputPath := CreatePath([]string{"output"})
	os.MkdirAll(outputPath, 0755)

	// Output HTML document, containing all of data
	outputHtml := filepath.Join(outputPath, "index.html")
	file, err := os.Create(outputHtml)
	if err != nil {
		log.Fatalf("Error creating output files: %v", err)
	}

	// Data struct in file: 'page.go'
	// These properties are same in template: 'src/index.html'
	myData := Page{
		Title:   "My Amazing Page",
		Content: "Hello......",
	}

	tmpl.ExecuteTemplate(file, "index.html", myData)

	///testing...
	config := Config{}
	config.Read()

}
