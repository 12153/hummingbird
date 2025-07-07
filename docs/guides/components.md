# Component Library

Hummingbird provides a comprehensive component library organized into categories for different use cases. Components are built with TypeScript and designed to work seamlessly with the Templ templating system.

## Component System

### Component Definition

Components are defined using the `defineComponent` function:

```typescript
import { defineComponent } from "../component-loader";

defineComponent("MyComponent", (el) => {
    // Component logic here
    // el is the DOM element where the component will be mounted
});
```

### Component Usage in Templates

Components are used in Templ templates with data attributes:

```templ
<div data-component="MyComponent" data-props='{"key": "value"}'></div>
```

## Available Component Categories

### Base Components

Foundation components for building interfaces.

**Location:** `src/components/base/`

Basic building blocks like buttons, inputs, and containers.

### Data Display Components

Components for presenting data and information.

**Location:** `src/components/data-display/`

Tables, lists, cards, and other data visualization components.

### Feedback Components

Components for user feedback and notifications.

**Location:** `src/components/feedback/`

Alerts, toasts, loading indicators, and progress bars.

### Form Components

Interactive form elements and validation.

**Location:** `src/components/forms/`

Form inputs, validation, and form management utilities.

### Layout Components

Structural components for page layout.

**Location:** `src/components/layout/`

Grids, containers, sidebars, and layout utilities.

### Navigation Components

Components for site navigation.

**Location:** `src/components/navigation/`

Menus, breadcrumbs, pagination, and navigation bars.

### Overlay Components

Modal and overlay components.

**Location:** `src/components/overlays/`

Modals, tooltips, dropdowns, and popup components.

## Example Components

### Counter Component

A simple interactive counter component:

```typescript
// assets/components/Counter.ts
import { defineComponent } from "../component-loader";

defineComponent("Counter", (el) => {
    let count = 0;

    const button = document.createElement("button");
    button.textContent = "Count: 0";
    button.className = "p-2 bg-blue-500 text-white rounded";

    button.addEventListener("click", () => {
        count++;
        button.textContent = `Count: ${count}`;
    });

    el.appendChild(button);
});
```

Usage in template:

```templ
<div data-component="Counter"></div>
```

### Button Component

A reusable button component with props:

```typescript
// assets/components/Button.ts
import { defineComponent } from "../component-loader";

interface ButtonProps {
    text: string;
    variant?: "primary" | "secondary";
    onClick?: string;
}

defineComponent("Button", (el, props: ButtonProps) => {
    const button = document.createElement("button");
    button.textContent = props.text;
    
    const baseClasses = "px-4 py-2 rounded font-medium";
    const variantClasses = {
        primary: "bg-blue-500 text-white hover:bg-blue-600",
        secondary: "bg-gray-200 text-gray-800 hover:bg-gray-300"
    };
    
    button.className = `${baseClasses} ${variantClasses[props.variant || "primary"]}`;
    
    if (props.onClick) {
        button.addEventListener("click", () => {
            // Execute callback or custom logic
            eval(props.onClick);
        });
    }
    
    el.appendChild(button);
});
```

Usage in template:

```templ
<div data-component="Button" data-props='{"text": "Click me", "variant": "primary"}'></div>
```

### Modal Component

A modal overlay component:

```typescript
// assets/components/Modal.ts
import { defineComponent } from "../component-loader";

interface ModalProps {
    title: string;
    content: string;
    isOpen?: boolean;
}

defineComponent("Modal", (el, props: ModalProps) => {
    const modal = document.createElement("div");
    modal.className = "fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50";
    modal.style.display = props.isOpen ? "flex" : "none";
    
    const modalContent = document.createElement("div");
    modalContent.className = "bg-white rounded-lg p-6 max-w-md w-full mx-4";
    
    const title = document.createElement("h2");
    title.textContent = props.title;
    title.className = "text-xl font-bold mb-4";
    
    const content = document.createElement("p");
    content.textContent = props.content;
    content.className = "text-gray-600 mb-4";
    
    const closeButton = document.createElement("button");
    closeButton.textContent = "Close";
    closeButton.className = "px-4 py-2 bg-gray-500 text-white rounded";
    closeButton.addEventListener("click", () => {
        modal.style.display = "none";
    });
    
    modalContent.appendChild(title);
    modalContent.appendChild(content);
    modalContent.appendChild(closeButton);
    modal.appendChild(modalContent);
    
    document.body.appendChild(modal);
    
    // Store reference for external control
    (el as any).modal = modal;
});
```

## Component Loader System

### Component Registration

Components are automatically loaded when the page loads:

```typescript
// assets/component-loader.ts
const components = new Map<string, (el: HTMLElement, props?: any) => void>();

export function defineComponent(name: string, factory: (el: HTMLElement, props?: any) => void) {
    components.set(name, factory);
}

export function loadComponents() {
    document.querySelectorAll("[data-component]").forEach((el) => {
        const componentName = el.getAttribute("data-component");
        const propsAttr = el.getAttribute("data-props");
        
        if (componentName && components.has(componentName)) {
            const props = propsAttr ? JSON.parse(propsAttr) : {};
            components.get(componentName)!(el as HTMLElement, props);
        }
    });
}
```

### Automatic Loading

Components are loaded on DOM ready:

```typescript
// assets/main.ts
import { loadComponents } from "./component-loader";

document.addEventListener("DOMContentLoaded", () => {
    loadComponents();
});
```

## Best Practices

### Component Organization

1. **Single Responsibility**: Each component should have a single, clear purpose
2. **Consistent Naming**: Use PascalCase for component names
3. **Type Safety**: Always define TypeScript interfaces for props
4. **Accessibility**: Include proper ARIA attributes and keyboard navigation

### Performance

1. **Lazy Loading**: Only load components when needed
2. **Memory Management**: Clean up event listeners and timers
3. **Efficient Updates**: Use efficient DOM manipulation techniques
4. **Caching**: Cache expensive calculations

### Styling

1. **Tailwind CSS**: Use Tailwind classes for consistent styling
2. **Component Variants**: Support different visual variants
3. **Responsive Design**: Ensure components work on all screen sizes
4. **Dark Mode**: Consider dark mode compatibility

## Creating Custom Components

### 1. Define the Component

```typescript
// assets/components/MyComponent.ts
import { defineComponent } from "../component-loader";

interface MyComponentProps {
    title: string;
    items: string[];
}

defineComponent("MyComponent", (el, props: MyComponentProps) => {
    // Component implementation
});
```

### 2. Use in Templates

```templ
<div data-component="MyComponent" data-props='{"title": "My List", "items": ["Item 1", "Item 2"]}'></div>
```

### 3. Import in Main

```typescript
// assets/main.ts
import "./components/MyComponent";
```

## Next Steps

- [Getting Started](./getting-started.md) - Create your first project
- [API Reference](../api/) - Detailed API documentation
- [Tutorials](../tutorials/) - Advanced component examples