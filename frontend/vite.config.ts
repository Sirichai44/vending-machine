import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'
import path from 'path';

// https://vite.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
    },
  },
  plugins: [react()],
  server: {
    port: 3000,
    proxy: {
      '^/api': {
        target: 'http://localhost:8080',
        changeOrigin: true
      }
    }
  },
})
