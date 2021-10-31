import { createApp } from 'vue';
import router from './router'
import "tailwindcss/tailwind.css"
import App from './App.vue';
import { 
  applyPolyfills,
  defineCustomElements
} from '@aws-amplify/ui-components/loader';

import Amplify from 'aws-amplify';
import awsconfig from './aws-exports';
import './assets/tailwind.css'
Amplify.configure(awsconfig);

applyPolyfills().then(() => {
  defineCustomElements(window);
});

const app = createApp(App);
app.config.isCustomElement = tag => tag.startsWith('amplify-');
app.use(router)
app.mount('#app');