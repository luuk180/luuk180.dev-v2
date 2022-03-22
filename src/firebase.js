import { initializeApp } from "firebase/app";
import {getFirestore} from "firebase/firestore";
import { getStorage } from 'firebase/storage';

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
export const firebase = initializeApp(firebaseConfig);
export const db = getFirestore();
export const storage = getStorage();