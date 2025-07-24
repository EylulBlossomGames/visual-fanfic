package main

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	// Getting template
	templatePath := CreatePath([]string{"src", "index.html"})
	tmpl, err := template.New("main_template").ParseFiles(templatePath)
	if err != nil {
		log.Fatalf("Error creating template: %v", err)
	}

	// Creating output folder if does not exist
	outputPath := CreatePath([]string{"output"})
	CreateDirIfDoesNotExist(outputPath)

	// Heavy tasks in goroutines to saving time
	var wg sync.WaitGroup

	// Copying assets to output folder
	assetsSourcePath := CreatePath([]string{"src", "assets"})
	wg.Add(1)
	go CopyDir(&wg, assetsSourcePath, outputPath)

	faviconPath := CreatePath([]string{"src", "favicon.ico"})
	wg.Add(1)
	go CopyFile(&wg, faviconPath, outputPath)

	// Loading TOML files
	config := Config{}
	wg.Add(1)
	go config.Load(&wg)

	characters := Characters{}
	wg.Add(1)
	go characters.Load(&wg)

	wg.Wait()

	// Output HTML document, containing all of data
	outputHtml := filepath.Join(outputPath, "index.html")
	file, err := os.Create(outputHtml)
	if err != nil {
		log.Fatalf("Error creating output files: %v", err)
	}

	// Data struct in file: 'page.go'
	// These properties are same in template: 'src/index.html'
	myData := Page{
		Title:   config.Title,
		Content: "",
		Css:     "",
		LogoSrc: config.LogoSrc,
	}

	tmpl.ExecuteTemplate(file, "index.html", myData)

}
