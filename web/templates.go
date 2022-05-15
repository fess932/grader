package web

import (
	"embed"
	"fmt"
	"html/template"
)

//go:embed *.html
var fs embed.FS

func ParseTemplates() (*template.Template, error) {
	t, err := template.ParseFS(fs, "*.html")
	if err != nil {
		return nil, fmt.Errorf("parse templates: %w", err)
	}

	return t, nil
}
