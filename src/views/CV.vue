<template>
    <div class="container flex justify-center p-2">
        <div class="place-content-center">
            <vue-pdf-embed :source="linkedPath"/>
            <br />
            <a :href=linkedPath class="bg-black text-white p-2 rounded-l text-xl">Download CV</a>
        </div>
    </div>
</template>

<script>
import vuePdfEmbed from 'vue-pdf-embed';

export default {
  data() {
    return {
      linkedPath: String
    }
  },
  created() {this.getPath()},
  methods: {
    getPath() {
      this.linkedPath = "https://files.luuk180.dev/cv.pdf";
    }
  },
  components: {
    vuePdfEmbed
  }
}
</script>