# Hummingbird Framework Documentation

Hummingbird is a full-stack web framework that combines the power of Go with modern frontend technologies. It uses Templ for server-side templating and provides a rich TypeScript/React component library for interactive UI elements.

## Architecture

Hummingbird follows a hybrid approach:

- **Backend**: Go with Echo web framework and Templ for HTML templating
- **Frontend**: TypeScript/React components with Vite for bundling
- **Styling**: Tailwind CSS with component-based architecture
- **CLI**: Go-based scaffolding tool for project generation

## Key Features

- **Server-Side Rendering**: Fast initial page loads with Templ templates
- **Partial Updates**: HTMX-style partial page updates for dynamic content
- **Component Library**: Pre-built React components for common UI patterns
- **TypeScript Support**: Full type safety across the stack
- **Hot Reloading**: Development server with live reload
- **CLI Tools**: Scaffolding and code generation utilities

## Project Structure

```
hummingbird/
├── src/
│   ├── cli/           # CLI tool for scaffolding
│   ├── components/    # React component library
│   ├── hooks/         # Custom React hooks
│   ├── themes/        # Theme definitions
│   ├── types/         # TypeScript type definitions
│   └── utils/         # Utility functions
├── docs/              # Documentation
├── examples/          # Example applications
└── tests/             # Test suites
```

## Getting Started

1. **Installation**: Install the Hummingbird CLI
2. **Project Creation**: Use the CLI to scaffold a new project
3. **Development**: Start the development server
4. **Component Usage**: Import and use components in your templates

## Documentation Structure

- **[Getting Started](./guides/getting-started.md)** - Installation and first project
- **[CLI Reference](./guides/cli-reference.md)** - Command-line tool documentation
- **[Component Library](./guides/components.md)** - Available components and usage
- **[API Reference](./api/)** - Detailed API documentation
- **[Tutorials](./tutorials/)** - Step-by-step guides
- **[Examples](../examples/)** - Working example projects

## Contributing

See the main repository README for contribution guidelines.