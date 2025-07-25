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

	// Content boxes and their styles
	content := CreateDialogueBoxes(&dialogue.DialogueLines, &cast.Characters)
	castCss := CreateCssForCharacters(&cast.Characters)

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
		CastCss: castCss,
	}

	tmpl.ExecuteTemplate(file, "index.html", myData)

}

func CreateDialogueBoxes(dialogueLines *[]DialogueLine, characterList *[]Character) template.HTML {
	// Creating dialogue boxes
	contentBlocks := []string{}

	chVisibleNames := make(map[string]string)
	for _, character := range *characterList {
		chVisibleNames[character.CodeName] = character.VisibleName
	}

	for _, dl := range *dialogueLines {
		imgSrc := ""
		imgAlt := dl.Cn
		textLine := dl.Text

		narratorClass := ""
		narratorChBoxClass := ""
		if strings.HasPrefix(dl.Cn, "__") {
			narratorClass = "narrator-box"
			narratorChBoxClass = "narrator-character-name-box"
		}

		visibleName := chVisibleNames[dl.Cn]

		dlCode := fmt.Sprintf(
			"<div class='d-line container character-%s'>"+
				"<div class='character-box one-third column %s'><img src='%s' alt='%s'></div>"+
				"<div class='dialog-box two-thirds column'><h5 class='character-name-box %s'>%s:</h5><p>%s</p></div>"+
				"</div>",
			dl.Cn,
			narratorClass,
			imgSrc,
			imgAlt,
			narratorChBoxClass,
			visibleName,
			textLine,
		)

		contentBlocks = append(contentBlocks, dlCode)
	}

	content := template.HTML(strings.Join(contentBlocks, ""))

	return content
}

func CreateCssForCharacters(characterList *[]Character) template.CSS {
	cssBlocks := []string{}

	for _, character := range *characterList {
		fontStyle := ""
		if character.TextStyle == "i" {
			fontStyle = "font-style: italic;"
		} else if character.TextStyle == "b" {
			fontStyle = "font-weight: bold;"
		} else if character.TextStyle == "bi" || character.TextStyle == "ib" {
			fontStyle = "font-weight: bold;font-style: italic;"
		}

		txtColor := ""
		if character.TextColor != "" {
			txtColor = fmt.Sprintf("color: %s;", character.TextColor)
		}

		characterBoxColor := ""
		if character.CharacterBoxColor != "" {
			characterBoxColor = fmt.Sprintf("background-color: %s;", character.CharacterBoxColor)
		}

		dialogBoxColor := ""
		if character.DialogBoxColor != "" {
			dialogBoxColor = fmt.Sprintf("background-color: %s;", character.DialogBoxColor)
		}

		nameColor := ""
		if character.NameColor != "" {
			nameColor = fmt.Sprintf("color: %s;", character.NameColor)
		}

		allStyles := strings.Join([]string{
			fontStyle,
			txtColor,
		}, "")

		characterClassname := fmt.Sprintf(".character-%s", character.CodeName)

		cssPiece := fmt.Sprintf(
			"%s { %s }\n"+
				"%s .character-box { %s }\n"+
				"%s .dialog-box { %s }\n"+
				"%s .character-name-box { %s }\n\n",
			characterClassname,
			allStyles,
			characterClassname,
			characterBoxColor,
			characterClassname,
			dialogBoxColor,
			characterClassname,
			nameColor,
		)

		cssBlocks = append(cssBlocks, cssPiece)
	}

	return template.CSS(strings.Join(cssBlocks, ""))
}
