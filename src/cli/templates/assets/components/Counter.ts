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
