<template>
    <div class="flex-auto w-screen">
        <figure class="rounded-xl bg-gray-800 bg-opacity-80 p-2 w-48 float-left" v-for="doc in github" :key="doc.name">
            <div id="title" class="text-center p-1 text-xl bg-gray-600 rounded-xl">
                {{doc.name}}
            </div>
            {{doc.url}}
            <br />
            {{doc.description}}
            <br />
            Size: {{doc.diskUsage}} kB
            <br />
            <a :href=doc.homepageUrl>{{doc.homepageUrl}}</a>
        </figure>
    </div>
</template>

<script>
import { collection, getDocs } from "firebase/firestore";
import { db } from '@/firebase'

export default {
  data() {
    return {
      github: []
    }
  },
  created() {
    this.getGithub()
  },
  methods: {
    async getGithub() {
      const getGH = await getDocs(collection(db, "github"));
      getGH.forEach((doc) => {
        this.github.push(doc.data());
      })
    }
  }
}
</script>