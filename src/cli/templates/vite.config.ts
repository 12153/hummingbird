import { defineConfig } from "vite";

export default defineConfig({
  root: "assets",
  build: {
    outDir: "../public",
    emptyOutDir: true,
    rollupOptions: {
      input: {
        main: "assets/main.ts"
      },
      output: {
        entryFileNames: "main.js"
      }
    }
  }
});
