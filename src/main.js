import { createApp } from 'vue';
import router from './router'
import "tailwindcss/tailwind.css"
import App from './App.vue';
import './assets/tailwind.css'
import { firebase } from './firebase';
import { getAnalytics } from "firebase/analytics";

getAnalytics(firebase);
const app = createApp(App);
app.use(router)
app.mount('#app');