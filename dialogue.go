package main

import "sync"

type DialogueLine struct {
	Cn   string // Code Name in short way
	Text string
}

type Dialogue struct {
	DialogueLines []DialogueLine
}

func (d *Dialogue) Load(wg *sync.WaitGroup) {
	defer wg.Done()

	ParseToml([]string{"config", "dialogue.toml"}, d)
}
