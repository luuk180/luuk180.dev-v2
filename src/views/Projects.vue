<template>
    Hello projects!
</template>

<script setup>
import { collection, getDocs } from "firebase/firestore";
import { db } from '@/firebase/init'

const getGH = await getDocs(collection(db, "github"));
getGH.forEach((doc) => {
  console.log(doc.data());
})
</script>