<template>
  <div id="home-page" class="container mx-auto py-4 w-full text-white text-l">
    <h1 id="welcome-text" class="w-full">{{welcome}}</h1>
    <p>
        Hey! I'm Luuk, I am currently learning web development and I will be going back to school full-time starting in September.
        On this website you can find some projects I'm working on and you can find more information about me.
    </p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      counter: 0,
      welcome: "",
      welcomeText: "Welcome to my part of the internet!"
    }
  },
  created() {
    this.scrollWelcome();
  },
  methods: {
    scrollWelcome() {
        if (this.counter < this.welcomeText.length) {
          this.welcome += this.welcomeText.charAt(this.counter);
          this.counter++;
          setTimeout(() => this.scrollWelcome(), 500);
        }
    }
  }
}
</script>