package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates/*
var templateFS embed.FS

func main() {
	args := os.Args
	if len(args) < 3 || args[1] != "init" {
		fmt.Println("Usage: hummingbird init <project-name>")
		os.Exit(1)
	}

	project := args[2]
	if err := scaffold(project); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("✅ Project '%s' initialized.\n", project)
}

func scaffold(name string) error {
	dirs := []string{
		"app/pages",
		"app/components",
		"assets",
		"public",
	}

	for _, d := range dirs {
		if err := os.MkdirAll(filepath.Join(name, d), 0755); err != nil {
			return err
		}
	}

	entries, err := templateFS.ReadDir("templates")
	if err != nil {
		return err
	}

	for _, entry := range entries {
		rawName := entry.Name()
		targetPath := filepath.Join(name, strings.TrimSuffix(rawName, ".tmpl"))
		data, err := templateFS.ReadFile("templates/" + rawName)
		if err != nil {
			return err
		}

		if strings.HasSuffix(rawName, ".templ.tmpl") {
			tmpl, err := template.New(rawName).Parse(string(data))
			if err != nil {
				return err
			}
			f, err := os.Create(targetPath)
			if err != nil {
				return err
			}
			defer f.Close()
			tmpl.Execute(f, nil)
		} else {
			err = os.WriteFile(targetPath, data, 0644)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
