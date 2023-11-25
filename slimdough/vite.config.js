import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte()],
  server: {
    proxy: {
      '/public': {
        target: 'http://localhost:8000'
      },
      '/api': {
        target: 'http://localhost:8000'
      },
      '/auth': {
        target: 'http://localhost:8000'
      },
    }
  }
})
