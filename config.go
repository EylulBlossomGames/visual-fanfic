package main

import "sync"

type Config struct {
	Title             string
	Author            string
	FooterBoxColor    string
	FooterFontColor   string
	FooterText        string
	HeaderColor       string
	HeaderFontColor   string
	TitleFont         string
	Font              string
	LogoSrc           string
	OuterBoxColor     string
	MainBoxColor      string
	CharacterBoxColor string
	DialogBoxColor    string
	DefaultTextColor  string
	BorderWidth       string
}

func (c *Config) Load(wg *sync.WaitGroup) {
	defer wg.Done()

	ParseToml([]string{"config", "main.toml"}, c)
}
