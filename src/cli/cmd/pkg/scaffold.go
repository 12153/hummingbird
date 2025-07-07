package pkg

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/12153/hummingbird/cli/templates"
)

func Scaffold(name string) error {
	data := map[string]string{
		"ProjectName": name,
		"Author":      "",
		"Year":        time.Now().Format("2006"),
	}

	return fs.WalkDir(templates.FS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		// Create target path
		relPath := strings.TrimSuffix(path, ".tmpl") // drop .tmpl for output
		targetPath := filepath.Join(name, relPath)

		// Ensure parent dirs exist
		if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
			return err
		}

		dataBytes, err := templates.FS.ReadFile(path)
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, ".templ.tmpl") || strings.HasSuffix(path, ".go.tmpl") {
			tmpl, err := template.New(filepath.Base(path)).Parse(string(dataBytes))
			if err != nil {
				return err
			}

			f, err := os.Create(targetPath)
			if err != nil {
				return err
			}
			defer f.Close()

			return tmpl.Execute(f, data)
		}

		return os.WriteFile(targetPath, dataBytes, 0644)
	})
}
