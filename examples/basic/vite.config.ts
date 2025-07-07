import { defineConfig } from "vite";
import path from "path";
import { fileURLToPath } from "url";

// ESM-compatible __dirname
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

export default defineConfig({
  root: "assets",
  build: {
    outDir: "../public",
    emptyOutDir: true,
    rollupOptions: {
      input: {
        main: path.resolve(__dirname, "assets/main.ts"),
      },
      output: {
        entryFileNames: "main.js",
      },
    },
  },
  resolve: {
    alias: {
      "@components": path.resolve(__dirname, "assets/components"),
    },
  },
});
