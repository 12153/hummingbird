# Getting Started with Hummingbird

This guide will help you create your first Hummingbird application.

## Prerequisites

- Go 1.24.4 or later
- Node.js 18+ and npm/yarn
- Basic knowledge of Go and TypeScript/React

## Installation

### Install the CLI

```bash
go install github.com/12153/hummingbird/src/cli/cmd@latest
```

### Verify Installation

```bash
hummingbird --version
```

## Creating Your First Project

### 1. Scaffold a New Project

```bash
hummingbird new my-app
cd my-app
```

This creates a new project with the following structure:

```
my-app/
├── main.go              # Go server entry point
├── go.mod              # Go module file
├── package.json        # Node.js dependencies
├── app/
│   ├── layout.templ    # Base HTML layout
│   └── pages/
│       └── index.templ # Homepage template
├── assets/
│   ├── main.ts         # TypeScript entry point
│   └── components/     # React components
├── public/             # Static assets
└── vite.config.ts      # Vite configuration
```

### 2. Install Dependencies

```bash
# Install Go dependencies
go mod tidy

# Install Node.js dependencies
npm install
```

### 3. Start Development Server

```bash
# Start the development server with hot reload
air
```

Your application will be available at `http://localhost:3000`.

## Understanding the Project Structure

### Go Backend (`main.go`)

The Go server uses Echo framework and serves Templ templates:

```go
func main() {
    e := echo.New()
    e.Static("/assets", "public")
    
    e.GET("/", func(c echo.Context) error {
        page := pages.Index()
        return renderWithLayout(c, page)
    })
    
    e.Logger.Fatal(e.Start(":3000"))
}
```

### Templ Templates (`app/`)

Templates are written in Templ syntax and compiled to Go:

```templ
// app/pages/index.templ
package pages

templ Index() {
    <h1>Welcome to Hummingbird</h1>
    <div id="counter-app"></div>
}
```

### TypeScript Frontend (`assets/`)

React components are loaded dynamically:

```typescript
// assets/main.ts
import { loadComponents } from './component-loader'

document.addEventListener('DOMContentLoaded', () => {
    loadComponents()
})
```

## Adding Your First Component

### 1. Create a React Component

```typescript
// assets/components/Greeting.ts
import React from 'react'

interface GreetingProps {
    name: string
}

export const Greeting: React.FC<GreetingProps> = ({ name }) => {
    return <h2>Hello, {name}!</h2>
}
```

### 2. Update the Template

```templ
// app/pages/index.templ
package pages

templ Index() {
    <h1>Welcome to Hummingbird</h1>
    <div data-component="Greeting" data-props='{"name": "World"}'></div>
}
```

### 3. Register the Component

```typescript
// assets/component-loader.ts
import { Greeting } from './components/Greeting'

export const components = {
    Greeting,
    // ... other components
}
```

## Adding Routes

### 1. Add Route Handler

```go
// main.go
e.GET("/about", func(c echo.Context) error {
    page := pages.About()
    return renderWithLayout(c, page)
})
```

### 2. Create Template

```templ
// app/pages/about.templ
package pages

templ About() {
    <h1>About Us</h1>
    <p>This is the about page.</p>
}
```

## Partial Updates

For dynamic content updates without full page reloads:

```typescript
// Client-side
fetch('/api/data', {
    headers: { 'X-Partial': 'true' }
})
.then(response => response.text())
.then(html => {
    document.getElementById('content').innerHTML = html
})
```

```go
// Server-side
func renderWithLayout(c echo.Context, page templ.Component) error {
    if c.Request().Header.Get("X-Partial") == "true" {
        return page.Render(c.Request().Context(), c.Response())
    }
    layout := app.Layout(page)
    return layout.Render(c.Request().Context(), c.Response())
}
```

## Next Steps

- [CLI Reference](./cli-reference.md) - Learn about CLI commands
- [Component Library](./components.md) - Explore available components
- [API Reference](../api/) - Detailed API documentation
- [Tutorials](../tutorials/) - More advanced examples