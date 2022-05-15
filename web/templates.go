package web

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
)

//go:embed src/*.html
//go:embed assets/*
var fs embed.FS

func StaticFiles(stripPrefix string) http.Handler {
	return http.StripPrefix(stripPrefix, http.FileServer(http.FS(fs)))
}

func ParseTemplates() (*template.Template, error) {
	t, err := template.ParseFS(fs, "src/*.html")
	if err != nil {
		return nil, fmt.Errorf("parse src: %w", err)
	}

	return t, nil
}
