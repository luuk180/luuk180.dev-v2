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
Amplify.configure(awsconfig);

applyPolyfills().then(() => {
  defineCustomElements(window);
});

const app = createApp(App);
app.use(router)
app.config.isCustomElement = tag => tag.startsWith('amplify-');
app.mount('#app');