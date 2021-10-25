<template>
  <div id="nav" class="bg-fixed w-screen space-x-4 text-center text-gray-800 bg-blue-800 bg-opacity-80 text-xl p-2" style="background-image: ">
    <router-link class="bg-yellow-400 p-1.5 rounded" to="/">Home</router-link>
    <router-link class="bg-yellow-400 p-1.5 rounded" to="/projects">Projects</router-link>
    <router-link class="bg-yellow-400 p-1.5 rounded" to="/cv">CV</router-link>
    <router-link class="bg-yellow-400 p-1.5 rounded" to="/about">About</router-link>
  </div>
  <router-view/>
</template>