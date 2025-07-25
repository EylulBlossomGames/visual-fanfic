package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
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

	cast := Cast{}
	wg.Add(1)
	go cast.Load(&wg)

	dialogue := Dialogue{}
	wg.Add(1)
	go dialogue.Load(&wg)

	wg.Wait()

	content := CreateDialogueBoxes(dialogue.DialogueLines)

	// Output HTML document, containing all of data
	outputHtml := filepath.Join(outputPath, "index.html")
	file, err := os.Create(outputHtml)
	if err != nil {
		log.Fatalf("Error creating output files: %v", err)
	}

	// Data struct in file: 'page.go'
	// These properties are same in template: 'src/index.html'
	myData := Page{
		Config:  config,
		Content: content,
	}

	tmpl.ExecuteTemplate(file, "index.html", myData)

}

func CreateDialogueBoxes(dialogueLines []DialogueLine) template.HTML {
	// Creating dialogue boxes
	contentBlocks := []string{}

	for _, dl := range dialogueLines {
		imgSrc := ""
		imgAlt := dl.Cn
		textLine := dl.Text

		narratorClass := ""
		if strings.HasPrefix(dl.Cn, "__") {
			narratorClass = "narrator-box"
		}

		dlCode := fmt.Sprintf(
			"<div class='d-line container'>"+
				"<div class='character-box one-third column %s'><img src='%s' alt='%s'></div>"+
				"<div class='dialog-box two-thirds column'><p>%s</p></div>"+
				"</div>",
			narratorClass,
			imgSrc,
			imgAlt,
			textLine,
		)

		contentBlocks = append(contentBlocks, dlCode)
	}

	content := template.HTML(strings.Join(contentBlocks, ""))

	return content
}
