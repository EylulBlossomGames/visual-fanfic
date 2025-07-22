package main

type Config struct {
	Color1 string
	Color2 string
}

func (c *Config) Read() {
	ParseToml([]string{"config", "main.toml"}, c)
}
