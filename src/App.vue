<template>
  <div id="app">
      <navbar></navbar>
      <div id="container" class="bg-dark">
          <router-view />
      </div>
  </div>
</template>
<script>
import Navbar from "@/components/Navbar";
export default {
  components: {Navbar}
}
</script>