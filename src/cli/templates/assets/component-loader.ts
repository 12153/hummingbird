type ComponentDef = (el: HTMLElement, props?: any) => void;

const registry: Record<string, ComponentDef> = {};

export function defineComponent(name: string, init: ComponentDef) {
  registry[name] = init;
}

export function hydrateComponents(scope: HTMLElement | Document = document) {
  scope.querySelectorAll<HTMLElement>("[data-component]").forEach((el) => {
    const name = el.dataset.component!;
    const props = el.dataset.props ? JSON.parse(el.dataset.props) : {};
    if (registry[name]) {
      registry[name](el, props);
    } else {
      console.warn(`Component '${name}' not registered.`);
    }
  });
}
