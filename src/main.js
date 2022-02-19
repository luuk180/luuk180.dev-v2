import { createApp } from 'vue';
import router from './router'
import "tailwindcss/tailwind.css"
import App from './App.vue';
import './assets/tailwind.css'

//firebase
// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAnalytics } from "firebase/analytics";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyDyMdYiC8aH1onIWn4LhMM-ei-lI8XydL4",
  authDomain: "luuk180-dev2.firebaseapp.com",
  projectId: "luuk180-dev2",
  storageBucket: "luuk180-dev2.appspot.com",
  messagingSenderId: "1056784744633",
  appId: "1:1056784744633:web:d40795f5c7bfaca7a29b68",
  measurementId: "G-FYP5SFWER3"
};

// Initialize Firebase
const fireApp = initializeApp(firebaseConfig);
const analytics = getAnalytics(fireApp);

const app = initializeApp(App);
app.use(router)
app.mount('#app');