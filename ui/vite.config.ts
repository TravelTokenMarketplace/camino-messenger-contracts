/// <reference types="vitest" />
import { copyFileSync } from "node:fs";
import { resolve } from "node:path";
import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// GitHub Pages serves no SPA fallback; copy index.html to 404.html so deep links resolve.
const spaFallback = {
  name: "spa-404-fallback",
  closeBundle() {
    copyFileSync(resolve("dist/index.html"), resolve("dist/404.html"));
  },
};

export default defineConfig({
  plugins: [react(), spaFallback],
  base: "/camino-messenger-contracts/",
  test: {
    globals: true,
    environment: "jsdom",
    setupFiles: ["./src/test-setup.ts"],
  },
});
