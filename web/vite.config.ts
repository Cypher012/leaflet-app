import { defineConfig } from 'vite'
import { devtools } from '@tanstack/devtools-vite'
import tsconfigPaths from 'vite-tsconfig-paths'

import { tanstackStart } from '@tanstack/react-start/plugin/vite'
import { cloudflare } from "@cloudflare/vite-plugin";

import viteReact from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/vite'

const config = defineConfig({
  plugins: [
    devtools(),
    cloudflare({ viteEnvironment: { name: "ssr" } }),
    tsconfigPaths({ projects: ['./tsconfig.json'] }),
    tailwindcss(),
    tanstackStart(),
    viteReact(),
  ],
  server: {
    allowedHosts: ['leaflet-dev.com'],
    proxy: {
      '/api': {
        target: 'https://api.leaflet-dev.com',
        changeOrigin: true,
        secure: false,
      },
    },
  },
})

export default config
