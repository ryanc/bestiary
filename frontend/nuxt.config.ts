// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },

  modules: ['@nuxt/ui', '@nuxt/eslint', '@nuxt/image'],

  css: ['~/assets/css/main.css'],

  compatibilityDate: '2025-07-16',

  $development: {
    runtimeConfig: {
      public: {
        apiBase: "http://localhost:8080/api",
      },
      apiBase: "http://backend:3000/api",
    },
  },

  runtimeConfig: {
    public: {
      apiBase: "http://localhost:8080/api",
    },
    apiBase: "http://backend:3000/api",
  },

  vite: {
    server: {
      // allowedHosts: ["frontend"],
      hmr: {
        protocol: "ws",
        host: "localhost",
        port: 3000,
        clientPort: 3000,
      }
    }
  }
})
