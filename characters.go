package main

import "sync"

type Character struct {
	CodeName    string
	VisibleName string
	Image       string
}

type Characters struct {
	Characters []Character
}

func (c *Characters) Load(wg *sync.WaitGroup) {
	defer wg.Done()

	ParseToml([]string{"config", "characters.toml"}, c)
}
