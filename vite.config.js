import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    port: 3001, // ← here!
    strictPort: true, // optional: error if 3001 is unavailable
  },
});
