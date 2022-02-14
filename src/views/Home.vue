<template>
  <div id="home-page" class="container mx-auto py-4 w-full text-gray-700 text-l">
    <h1 id="welcome-text" class="w-full">Welcome to my part of the internet!</h1>
    This is going to be my portfolio and I'm very busy building it right now!
  </div>
</template>

<style>
#welcome-text {
    border-right: solid 3px rgba(191, 219, 254, 1);
    white-space: nowrap;
    overflow: hidden;
    font-size: 28px;
    color: rgba(55, 65, 81, 1);
}

#welcome-text {
    animation: animated-text 4s steps(29) 1s 1 normal both,
                animated-cursor 600ms steps(29) infinite;
}

@keyframes animated-text{
    from{width: 0;}
    to {width: 472px;}
}

@keyframes animated-cursor{
    from{border-right-color: rgba(55,65,81,1);}
    to{border-right-color: transparent;}
}
</style>

<script>
    export default {
      name: "home",
    }
</script>