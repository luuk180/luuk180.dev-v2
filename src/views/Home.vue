<template>
  <div id="home-page" class="text-gray-700 text-xl">
    <h1 class="text-center">Welcome to my part of the internet!</h1>
    This is going to be my portfolio and I'm very busy building it right now!
  </div>
</template>

<style>
h1 {
    border-right: solid 3px rgba(0,255,0,.75);
    white-space: nowrap;
    overflow: hidden;
    font-size: 28px;
    color: rgba(255,255,255,.70);
}

h1 {
    animation: animated-text 4s steps(29) 1s 1 normal both,
                animated-cursor 600ms steps(29) infinite;
}

@keyframes animated-text{
    from{width: 0;}
    to {width: 472px;}
}

@keyframes animated-cursor{
    from{border-right-color: rgba(0,255,0,.75);}
    to{border-right-color: transparent;}
}
</style>

<script>
    export default {
      name: "home",
    }
</script>