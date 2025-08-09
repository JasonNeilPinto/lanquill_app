import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// https://vite.dev/config/

export default defineConfig({
  plugins: [react()],
  server: {
    port: 3001,
    strictPort: true,
  },

  // Speed up dependency pre-bundling
  optimizeDeps: {
    include: ["react", "react-dom"],
  },

  build: {
    // Use esbuild minifier (fast and good)
    minify: "esbuild",

    // Output directories & filenames
    outDir: "dist",
    assetsDir: "assets",
    emptyOutDir: true,
    cssCodeSplit: true,
    modulePreload: { polyfill: true },

    // Lower the noise for reasonable app sizes; adjust as needed
    chunkSizeWarningLimit: 600,

    // Inline very small assets; keep bigger ones as files (better cacheability)
    assetsInlineLimit: 4096, // 4kb

    // Tree-shaking hints
    treeshake: true,
    target: "es2020",

    // Rollup fine-tuning for caches & chunking
    rollupOptions: {
      output: {
        // Content-hashed filenames for long-term caching
        entryFileNames: "assets/[name].[hash].js",
        chunkFileNames: "assets/[name].[hash].js",
        assetFileNames: ({ name }) => {
          if (!name) return "assets/[name].[hash][extname]";
          if (/\.(css)$/.test(name)) return "assets/[name].[hash][extname]";
          if (/\.(png|jpe?g|gif|svg|webp|avif)$/.test(name))
            return "assets/images/[name].[hash][extname]";
          if (/\.(woff2?|ttf|otf|eot)$/.test(name))
            return "assets/fonts/[name].[hash][extname]";
          return "assets/[name].[hash][extname]";
        },

        // Manual chunking to keep bundles predictable
        manualChunks: {
          // Core React in its own chunk (good cache hit ratio)
          react: ["react", "react-dom"],
        },
      },
      // If you *know* certain packages have no side effects, you can hint here,
      // but be careful with libraries that ship CSS side effects.
      // treeshake: { moduleSideEffects: false },
    },

    // Report brotli sizes can slow builds; keep off unless you need it
    brotliSize: false,
  },

  // Define globals available at build time
  define: {
    __BUILD_TIME__: JSON.stringify(new Date().toISOString()),
  },

  // Cleaner preview defaults
  preview: {
    port: 4173,
    strictPort: true,
  },
});
