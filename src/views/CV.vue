<template>
    <div class="container mx-auto">
        <vue-pdf-embed :source=linkedPath></vue-pdf-embed>
    </div>
</template>

<script>
import { ref, getDownloadURL } from 'firebase/storage';
import { storage } from '@/firebase';
import vuePdfEmbed from 'vue-pdf-embed';
const pdfRef = ref(storage, 'Resume-Luuk-Volkering.pdf');

var linkedPath;

export default {
  data() {
    return {
      linkedPath
    }
  },
  created() {this.getPath()},
  methods: {
    getPath() {
      getDownloadURL(pdfRef).then((url) => {
        linkedPath = url;
      })
    }
  },
  components: {
    vuePdfEmbed
  }
}
</script>