<template>
    <div v-if="!isFetching" class="flex-auto">
        <div class="p-2">
            <div class="card-group">
            <div class="card" style="width: 18rem;" v-for="doc in github" :key="doc.name">
                <div class="card-header">
                    <div class="card-title">
                        <a class="card-link text-primary" :href=doc.url>{{doc.name}}</a>
                    </div>
                </div>
                <p class="card-text">{{doc.description}}</p>
                <p class="card-text">Size: {{doc.diskUsage}} kB</p>
                <a v-if="doc.homepageUrl" class="card-link card-footer" :href=doc.homepageUrl>{{doc.homepageUrl}}</a>
            </div>
            </div>
        </div>
        <div class="bottom-10 align-content-left text-white">You can click the name of the project to go to the repository on GitHub</div>
    </div>
    <div v-if="isFetching">
        <p class="text-xl-center">Page is loading...</p>
    </div>
</template>

<script>

export default {
  data() {
    return {
      github: [],
      isFetching: true,
    }
  },
  created() {
    this.getGithub()
  },
  methods: {
    async getGithub() {
      const res = await fetch("https://api.luuk180.dev/query");
      this.github = await res.json();
      this.isFetching = false;
    }
  }
}
</script>