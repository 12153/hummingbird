# CLI Reference

The Hummingbird CLI provides commands for scaffolding projects, adding components, and managing development workflows.

## Installation

```bash
go install github.com/12153/hummingbird/src/cli/cmd@latest
```

## Commands

### `hummingbird init <project-name>`

Creates a new Hummingbird project with the specified name.

**Usage:**
```bash
hummingbird init my-app
```

**What it does:**
- Creates a new directory with the project name
- Scaffolds the complete project structure
- Generates Go server code with Echo and Templ
- Sets up TypeScript/React frontend with Vite
- Configures Tailwind CSS
- Creates example components and templates

**Generated Structure:**
```
my-app/
├── main.go              # Go server entry point
├── go.mod              # Go module configuration
├── package.json        # Node.js dependencies
├── vite.config.ts      # Vite bundler configuration
├── tailwind.config.js  # Tailwind CSS configuration
├── postcss.config.js   # PostCSS configuration
├── .air.toml           # Air hot-reload configuration
├── .gitignore          # Git ignore patterns
├── app/
│   ├── layout.templ    # Base HTML layout template
│   └── pages/
│       └── index.templ # Homepage template
├── assets/
│   ├── main.ts         # TypeScript entry point
│   ├── component-loader.ts # Component loading system
│   ├── router.ts       # Client-side routing
│   └── components/
│       └── Counter.ts  # Example React component
└── public/             # Static assets directory
```

### `hummingbird add route <name>`

Adds a new route to an existing Hummingbird project.

**Usage:**
```bash
hummingbird add route about
```

**What it does:**
- Creates a new Templ template file in `app/pages/`
- Generates the corresponding Go template code
- Provides a basic route handler example

**Example Output:**
```templ
// app/pages/about.templ
package pages

templ About() {
    <h1>About Page</h1>
    <p>This is the about page.</p>
}
```

**Manual Route Registration:**
You'll need to manually add the route to your `main.go`:

```go
e.GET("/about", func(c echo.Context) error {
    page := pages.About()
    return renderWithLayout(c, page)
})
```

### `hummingbird dev <project-dir>`

Starts the development server with hot reloading.

**Usage:**
```bash
hummingbird dev .
# or
hummingbird dev /path/to/project
```

**What it does:**
- Starts the Go server with Air for hot reloading
- Watches for changes in Go and Templ files
- Automatically rebuilds and restarts the server
- Serves the application on `http://localhost:3000`

**Requirements:**
- [Air](https://github.com/cosmtrek/air) must be installed
- Project must have a valid `.air.toml` configuration

## Configuration Files

### `.air.toml`

Air configuration for hot reloading:

```toml
[build]
  cmd = "templ generate && go build -o ./tmp/main ."
  bin = "tmp/main"
  include_ext = ["go", "templ"]
  exclude_dir = ["tmp", "vendor", "node_modules"]
  
[misc]
  clean_on_exit = true
```

### Development Workflow

1. **Initialize Project:**
   ```bash
   hummingbird init my-app
   cd my-app
   ```

2. **Install Dependencies:**
   ```bash
   go mod tidy
   npm install
   ```

3. **Start Development:**
   ```bash
   hummingbird dev .
   ```

4. **Add New Routes:**
   ```bash
   hummingbird add route contact
   ```

## Common Issues

### Command Not Found

If you get "command not found" error:

1. Ensure Go is installed and `$GOPATH/bin` is in your PATH
2. Verify installation: `go list -m github.com/12153/hummingbird/src/cli/cmd`
3. Reinstall: `go install github.com/12153/hummingbird/src/cli/cmd@latest`

### Air Not Found

If `hummingbird dev` fails:

1. Install Air: `go install github.com/cosmtrek/air@latest`
2. Ensure Air is in your PATH
3. Verify with: `air -v`

### Template Generation Issues

If Templ templates aren't generating:

1. Install Templ CLI: `go install github.com/a-h/templ/cmd/templ@latest`
2. Generate manually: `templ generate`
3. Ensure templates have `.templ` extension

## Environment Variables

- `HUMMINGBIRD_PORT`: Override default port (3000)
- `HUMMINGBIRD_HOST`: Override default host (localhost)
- `NODE_ENV`: Set to "production" for production builds

## Next Steps

- [Getting Started](./getting-started.md) - Create your first project
- [Component Library](./components.md) - Available components
- [API Reference](../api/) - Detailed API documentation