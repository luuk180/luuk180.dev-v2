<template>
    <header class="bg-dark p-2 text-white">
        <nav class="navbar navbar-expand-sm bg-dark navbar-dark">
            <div class="container-fluid">
                <a class="navbar-brand"
                    style="-webkit-user-select: none; /* Safari */
                        -moz-user-select: none; /* Firefox */
                        -ms-user-select: none; /* IE10+/Edge */
                        user-select: none; /* Standard */">
                    luuk180
                </a>
                <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#collapsibleNavbar">
                    <span class="navbar-toggler-icon"></span>
                </button>
                <div  class="collapse navbar-collapse" id="collapsibleNavbar">
                    <ul class="navbar-nav">
                        <li>
                            <div v-if="home">
                                <router-link class="nav-link px-3 text-secondary" to="/">Home</router-link>
                            </div>
                            <div v-else>
                                <router-link class="nav-link px-3 text-white" to="/">Home</router-link>
                            </div>
                        </li>
                        <li>
                            <div v-if="projects">
                                <router-link class="nav-link px-3 text-secondary" to="/projects">Projects</router-link>
                            </div>
                            <div v-else>
                                <router-link class="nav-link px-3 text-white" to="/projects">Projects</router-link>
                            </div>
                        </li>
                        <li>
                            <div v-if="cv">
                                <router-link class="nav-link px-3 text-secondary" to="/cv">CV</router-link>
                            </div>
                            <div v-else>
                                <router-link class="nav-link px-3 text-white" to="/cv">CV</router-link>
                            </div>
                        </li>
                    </ul>
                </div>
            </div>
        </nav>
    </header>
</template>

<script>

export default {
  data () {
    return {
      home: false,
      projects: false,
      cv: false,
    }
  },
  watch: {
    '$route' () {
      switch (this.$route.path) {
        case "/":
          this.home = true;
          this.projects = false;
          this.cv = false;
          break;
        case "/projects":
          this.projects = true;
          this.home = false;
          this.cv = false;
          break;
        case "/cv":
          this.cv = true;
          this.home = false;
          this.projects = false;
          break;
      }
    }
  }
}
</script>