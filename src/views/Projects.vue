<template>
    <div class="flex-auto">
        <div class="p-2">
            <figure class="gap-10 rounded-xl bg-gray-800 bg-opacity-80 p-2 w-48 float-left" v-for="doc in github" :key="doc.name">
                <div id="title" class="text-center p-1 text-xl bg-gray-600 rounded-xl">
                    <a :href=doc.url>{{doc.name}}</a>
                </div>
                {{doc.description}}
                <br />
                Size: {{doc.diskUsage}} kB
                <br />
                <a :href=doc.homepageUrl>{{doc.homepageUrl}}</a>
            </figure>
        </div>
        <div class="absolute p-2 bottom-0 float-center text-white">You can click the name of the project to go to the repository on GitHub</div>
    </div>
</template>

<script>

export default {
  data() {
    return {
      github: []
    }
  },
  created() {
    this.getGithub()
  },
  methods: {
    async getGithub() {
      const res = await fetch("https://back.luuk180.dev/query");
      this.github = await res.json();
    }
  }
}
</script>