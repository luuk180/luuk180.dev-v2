<template>
    <div class="container">
        <vue-pdf-embed :source="linkedPath"/>
        <br />
        <div class="d-flex justify-content-center">
            <a role="button" class="btn btn-primary" :href=linkedPath>Download CV</a>
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