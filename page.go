package main

import "html/template"

type Page struct {
	Config  Config
	Content template.HTML
	CastCss template.CSS
}
