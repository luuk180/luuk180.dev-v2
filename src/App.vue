<template>
  <div id="app">
      <navbar></navbar>
      <div id="container" class="absolute inset-x-0 top-12 bottom-0 bg-gray-600 text-white">
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