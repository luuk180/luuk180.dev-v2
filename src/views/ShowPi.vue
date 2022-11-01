<template>
     <ul class="list-group">
         <div v-if="pis">
             <div v-for="pi in pis" :key="pi.HostName">
                 <li class="list-group-item bg-secondary text-white">{{pi.Hostname}}: {{pi.LocalIp}}</li>
             </div>
         </div>
         <div v-if="!pis">
             <li class="list-group-item bg-secondary text-white">No pi's found!</li>
         </div>
     </ul>
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