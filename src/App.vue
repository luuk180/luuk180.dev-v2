<template>
  <div id="app">
      <navbar></navbar>
      <div class="bg-secondary">
          <router-view />
      </div>
  </div>
</template>
<script>
import Navbar from "@/components/Navbar";
export default {
  components: {Navbar},
  mounted() {
    document.title = "luuk180.dev"
  }
}
</script>