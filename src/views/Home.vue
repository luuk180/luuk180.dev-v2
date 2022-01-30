<template>
  <div id="home-page" class="text-gray-700 text-xl">
    <h1 class="text-center">{{welcomeText}}</h1>
    This is going to be my portfolio and I'm very busy building it right now!
  </div>
</template>

<script>
    export default {
      name: "home",
      props: {
        welcomeText: "",
      },
      methods: {
        animateWelcome() {
          let finalText = "Welcome to luuk180.dev!";
          for(let i = 0; i < finalText.length; i++) {
            this.welcomeText.concat(finalText[i]);
          }
        }
      },
      created() {
        this.animateWelcome();
      }
    }
</script>