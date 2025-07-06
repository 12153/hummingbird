// Package templates
package templates

import "embed"

//go:embed app/** assets/** main.go.tmpl vite.config.ts tailwind.config.js postcss.config.js package.json README.md
var FS embed.FS
