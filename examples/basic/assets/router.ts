import { hydrateComponents } from "./component-loader";

document.addEventListener("click", (e) => {
  const link = (e.target as HTMLElement).closest("a");
  if (link?.hasAttribute("data-swap") || link?.hostname === location.hostname) {
    e.preventDefault();
    fetch(link.href, { headers: { "X-Partial": "true" } })
      .then((res) => res.text())
      .then((html) => {
        document.querySelector("main")!.innerHTML = html;
        hydrateComponents();
        history.pushState({}, "", link.href);
      });
  }
});
