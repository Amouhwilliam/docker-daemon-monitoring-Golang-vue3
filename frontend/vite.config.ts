import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig(() => {
  const config = {
    plugins: [
      vue(),
    ],
    server: {
      host: '0.0.0.0',
      port: 5100,
      proxy: {}
    },
    publicDir: false,
    base: ''
  };

  // Install proxy rules based on PROXY env variable
  const proxyEnv = process.env.PROXY || '';
  const proxyMatches = proxyEnv.matchAll(/(\S+)\s*->\s*(\S+)/g);
  for (let match of proxyMatches) {
    const path = match[1];
    const target = match[2];
    console.log(`Adding proxy ${path} -> ${target}`);
    config.server.proxy[path] = {
      target,
      rewrite: p => p.replace(new RegExp(`^${path}`), '')
    };
  }

  return config;
});

