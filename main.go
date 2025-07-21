package main

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Obtaining current directory
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error al acceder al sistema de archivos: %v", err)
	}

	// Getting template
	templatePath := filepath.Join(currentDir, "src", "index.html")
	tmpl, err := template.New("Test").ParseFiles(templatePath)
	if err != nil {
		log.Fatalf("Error al crear la plantilla: %v", err)
	}

	// Creating output folder if does not exist
	outputPath := filepath.Join(currentDir, "output")
	os.MkdirAll(outputPath, 0755)

	// Output HTML document, containing all of data
	outputHtml := filepath.Join(outputPath, "index.html")
	file, err := os.Create(outputHtml)
	if err != nil {
		log.Fatalf("Error al crear el archivo: %v", err)
	}

	// Data struct in file: 'page.go'
	// These properties are same in template: 'src/index.html'
	myData := Page{
		Title:   "My Amazing Page",
		Content: "Hello......",
	}

	tmpl.ExecuteTemplate(file, "index.html", myData)

}
