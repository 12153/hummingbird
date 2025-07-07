package pkg 

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func AddRoute(name, projectRoot string) error {
	// Normalize route path
	segments := strings.Split(name, "/")
	fileName := segments[len(segments)-1]
	if !strings.HasSuffix(fileName, ".templ") {
		fileName += ".templ"
	}

	routePath := filepath.Join(append([]string{projectRoot, "app", "pages"}, segments...)...)
	if !strings.HasSuffix(routePath, ".templ") {
		routePath += ".templ"
	}

	if err := os.MkdirAll(filepath.Dir(routePath), 0755); err != nil {
		return err
	}

	// Inject name into template
	data := map[string]string{
		"FunctionName": routeFuncName(segments),
	}

	f, err := os.Create(routePath)
	if err != nil {
		return err
	}
	defer f.Close()

	const routeTemplate = `templ {{ .FunctionName }}() {
  <div>
    <h1>{{ .FunctionName }}</h1>
    <p>This is the '{{ .FunctionName }}' route.</p>
  </div>
}`

	tmpl, err := template.New("route").Parse(routeTemplate)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(f, data); err != nil {
		return err
	}

	fmt.Printf("✅ Created route: %s\n", routePath)
	return nil
}

// Generate function name from path: blog/[slug] → BlogSlug
func routeFuncName(segments []string) string {
	builder := strings.Builder{}
	for _, s := range segments {
		s = strings.TrimSuffix(s, ".templ")
		s = strings.Trim(s, "[]")
		s = strings.Title(s)
		builder.WriteString(s)
	}
	return builder.String()
}
