package main

type Character struct {
	CodeName    string
	VisibleName string
	Image       string
}

type Characters struct {
	Characters []Character
}

func (c *Characters) Load() {
	ParseToml([]string{"config", "characters.toml"}, c)
}
