# API Reference

This section provides detailed API documentation for the Hummingbird framework.

## Core APIs

### Server-Side APIs

#### `renderWithLayout(c echo.Context, page templ.Component) error`

Renders a Templ component with or without layout based on request headers.

**Parameters:**
- `c echo.Context` - Echo context containing request/response
- `page templ.Component` - Templ component to render

**Returns:**
- `error` - Rendering error if any

**Behavior:**
- If `X-Partial: true` header is present, renders only the component
- Otherwise, wraps the component in the base layout

**Example:**
```go
func handler(c echo.Context) error {
    page := pages.Home()
    return renderWithLayout(c, page)
}
```

### Client-Side APIs

#### `defineComponent(name: string, factory: ComponentFactory)`

Registers a new component with the component loader system.

**Parameters:**
- `name: string` - Component name (used in `data-component` attribute)
- `factory: ComponentFactory` - Function that creates the component

**ComponentFactory Type:**
```typescript
type ComponentFactory = (el: HTMLElement, props?: any) => void
```

**Example:**
```typescript
defineComponent("MyComponent", (el, props) => {
    // Component implementation
});
```

#### `loadComponents()`

Scans the DOM for elements with `data-component` attributes and initializes components.

**Usage:**
```typescript
import { loadComponents } from "./component-loader";

document.addEventListener("DOMContentLoaded", () => {
    loadComponents();
});
```

## Framework Modules

### Template System

#### Templ Template Structure

```templ
package pages

templ ComponentName(param Type) {
    <div>
        <!-- Template content -->
    </div>
}
```

**Key Features:**
- Type-safe parameters
- Automatic HTML escaping
- Component composition
- Conditional rendering

#### Template Composition

```templ
// Base layout
templ Layout(content templ.Component) {
    <!DOCTYPE html>
    <html>
        <head>
            <title>My App</title>
        </head>
        <body>
            @content
        </body>
    </html>
}

// Page component
templ HomePage() {
    <h1>Welcome</h1>
}

// Usage
layout := Layout(HomePage())
```

### Component System

#### Component Props Interface

```typescript
interface ComponentProps {
    [key: string]: any;
}
```

#### Component Lifecycle

1. **Registration**: Component is registered with `defineComponent`
2. **Discovery**: DOM is scanned for `data-component` attributes
3. **Props Parsing**: `data-props` attribute is parsed as JSON
4. **Instantiation**: Component factory function is called
5. **Mounting**: Component is attached to the DOM element

### Routing System

#### Server-Side Routing

```go
// Basic route
e.GET("/path", handlerFunction)

// Route with parameters
e.GET("/user/:id", func(c echo.Context) error {
    id := c.Param("id")
    // Handle request
})

// Route with query parameters
e.GET("/search", func(c echo.Context) error {
    query := c.QueryParam("q")
    // Handle search
})
```

#### Client-Side Router

```typescript
// assets/router.ts
interface Route {
    path: string;
    handler: (params: any) => void;
}

class Router {
    private routes: Route[] = [];
    
    addRoute(path: string, handler: (params: any) => void) {
        this.routes.push({ path, handler });
    }
    
    navigate(path: string) {
        // Handle navigation
    }
}
```

## Configuration APIs

### Vite Configuration

```typescript
// vite.config.ts
import { defineConfig } from 'vite';

export default defineConfig({
    build: {
        outDir: 'public',
        rollupOptions: {
            input: 'assets/main.ts'
        }
    },
    server: {
        proxy: {
            '/api': 'http://localhost:3000'
        }
    }
});
```

### Tailwind Configuration

```javascript
// tailwind.config.js
module.exports = {
    content: ['./app/**/*.templ'],
    theme: {
        extend: {
            colors: {
                primary: '#3B82F6'
            }
        }
    },
    plugins: []
};
```

## Utility Functions

### Type Definitions

```typescript
// Component factory type
type ComponentFactory = (el: HTMLElement, props?: any) => void;

// Props type
interface ComponentProps {
    [key: string]: any;
}

// Route handler type
type RouteHandler = (params: Record<string, string>) => void;
```

### Error Handling

```go
// Server-side error handling
func errorHandler(err error, c echo.Context) {
    if he, ok := err.(*echo.HTTPError); ok {
        c.JSON(he.Code, map[string]string{"error": he.Message.(string)})
    } else {
        c.JSON(500, map[string]string{"error": "Internal server error"})
    }
}
```

```typescript
// Client-side error handling
function handleComponentError(componentName: string, error: Error) {
    console.error(`Error in component ${componentName}:`, error);
    // Handle error appropriately
}
```

## Events and Hooks

### Component Events

```typescript
// Custom event dispatch
function dispatchComponentEvent(el: HTMLElement, eventName: string, data?: any) {
    const event = new CustomEvent(eventName, { detail: data });
    el.dispatchEvent(event);
}

// Event listener
el.addEventListener('custom-event', (e) => {
    console.log('Event data:', e.detail);
});
```

### Lifecycle Hooks

```typescript
// Component with lifecycle
defineComponent("MyComponent", (el, props) => {
    // Mount
    console.log("Component mounted");
    
    // Setup cleanup
    const cleanup = () => {
        console.log("Component cleanup");
    };
    
    // Store cleanup function
    (el as any).__cleanup = cleanup;
});
```

## Performance APIs

### Optimization Utilities

```typescript
// Debounce utility
function debounce(func: Function, wait: number) {
    let timeout: NodeJS.Timeout;
    return function executedFunction(...args: any[]) {
        const later = () => {
            clearTimeout(timeout);
            func(...args);
        };
        clearTimeout(timeout);
        timeout = setTimeout(later, wait);
    };
}

// Throttle utility
function throttle(func: Function, limit: number) {
    let inThrottle: boolean;
    return function executedFunction(...args: any[]) {
        if (!inThrottle) {
            func.apply(this, args);
            inThrottle = true;
            setTimeout(() => inThrottle = false, limit);
        }
    };
}
```

## Next Steps

- [Getting Started](../guides/getting-started.md) - Create your first project
- [CLI Reference](../guides/cli-reference.md) - Command-line tools
- [Component Library](../guides/components.md) - Available components
- [Tutorials](../tutorials/) - Step-by-step guides