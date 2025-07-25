package main

import "sync"

type Character struct {
	CodeName          string
	VisibleName       string
	Image             string
	TextStyle         string
	TextColor         string
	CharacterBoxColor string
	DialogBoxColor    string
	NameColor         string
}

type Cast struct {
	Characters []Character
}

func (c *Cast) Load(wg *sync.WaitGroup) {
	defer wg.Done()

	ParseToml([]string{"config", "characters.toml"}, c)
}
