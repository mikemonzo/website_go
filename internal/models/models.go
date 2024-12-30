package models

import "html/template"

type PageData struct {
	Title         string
	Author        string
	Welcome       string
	ErrorCode     int
	ErrorMessage  string
	HeaderContent template.HTML
	NavbarContent template.HTML
}
