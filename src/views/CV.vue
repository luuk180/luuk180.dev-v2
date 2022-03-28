<template>
    <div class="container mx-auto flex justify-center p-2">
        <div class="">
            <vue-pdf-embed :source=linkedPath />
        </div>
        <br />
        <a :href=linkedPath class="bg-black text-white p-2 rounded-l text-xl">Download CV</a>
    </div>
</template>

<script>
import { ref, getDownloadURL } from 'firebase/storage';
import { storage } from '@/firebase';
import vuePdfEmbed from 'vue-pdf-embed';
const pdfRef = ref(storage, 'Resume-Luuk-Volkering.pdf');

export default {
  data() {
    return {
      linkedPath: String
    }
  },
  created() {this.getPath()},
  methods: {
    getPath() {
      getDownloadURL(pdfRef).then((url) => {
        this.linkedPath = url;
      })
    }
  },
  components: {
    vuePdfEmbed
  }
}
</script>