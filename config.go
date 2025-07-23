package main

type Config struct {
	Title  string
	Color1 string
	Color2 string
}

func (c *Config) Load() {
	ParseToml([]string{"config", "main.toml"}, c)
}
