package main

type Config struct {
	Title             string
	OuterBoxColor     string
	MainBoxColor      string
	CharacterBoxColor string
	DialogBoxColor    string
	DefaultTextColor  string
	BorderWidth       string
}

func (c *Config) Load() {
	ParseToml([]string{"config", "main.toml"}, c)
}
