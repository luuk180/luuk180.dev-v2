<template>
 <div v-if="pis">
     <ul class="list-group">
         <div v-for="pi in pis" :key="pi.HostName">
             <li class="list-group-item bg-secondary text-white">{{pi.Hostname}}: {{pi.LocalIp}}</li>
         </div>
     </ul>
 </div>
</template>

<script>
export default {
  name: "ShowPi",

  data() {
    return {
      pis: [],
    }
  },

  created() {
    this.getPis();
  },
  methods: {
    async getPis() {
      const res = await fetch("https://api.luuk180.dev/getpis");
      this.pis = await res.json();
    }
  }
}
</script>

<style scoped>

</style>