import { defineComponent } from "../component-loader";

defineComponent("Counter", (el, props) => {
  let count = props.start;

  const button = document.createElement("button");
  button.textContent = `Count: ${count}`;
  button.className = "p-2 bg-blue-500 text-white rounded";

  button.addEventListener("click", () => {
    count++;
    button.textContent = `Count: ${count}`;
  });

  el.appendChild(button);
});
