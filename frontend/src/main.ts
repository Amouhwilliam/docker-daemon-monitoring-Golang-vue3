
import {createApp} from 'vue';
import {createRouter, createWebHistory} from 'vue-router';
import App from './components/App.vue';
import DockerContainerView from './components/DockerContainerView.vue';
import { VueQueryPlugin } from '@tanstack/vue-query'
import PrimeVue from 'primevue/config';
import './style.css'
import 'primevue/resources/themes/aura-light-green/theme.css'


const app = createApp(App);
app.use(VueQueryPlugin)
app.use(PrimeVue);

// Create router
const router = createRouter({
    history: createWebHistory('/'),
    routes: []
});

// Register routes
router.addRoute({
    path: '/containers',
    name: 'containers',
    component: DockerContainerView
});

// Redirect root to /info
router.addRoute({
    path: '/',
    name: 'home',
    redirect: '/containers'
});

// Bootstrap application
app.use(router);
app.mount('#app');
