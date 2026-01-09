import { createApp } from 'vue'
import './style.css'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'
import App from './App.vue'
import router from './router' 

const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)
// localStorage.clear();

createApp(App).use(pinia).use(router).mount('#app')