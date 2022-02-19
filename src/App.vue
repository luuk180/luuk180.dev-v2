<template>
  <div id="app">
    <div id="nav" class="bg-fixed w-screen space-x-4 text-center bg-black bg-opacity-80 text-2xl p-2 h-12">
        <p class="rounded bg-gray-600 text-white float-left w-28">luuk180</p>
        <router-link class="bg-gray-600 text-white p-1 rounded" to="/">Home</router-link>
        <router-link class="bg-gray-600 text-white p-1 rounded" to="/projects">Projects</router-link>
        <router-link class="bg-gray-600 text-white p-1 rounded" to="/cv">CV</router-link>
        <router-link class="bg-gray-600 text-white p-1 rounded" to="/about">About</router-link>
    </div>
    <div id="container" class="absolute inset-x-0 top-12 bottom-0 bg-gray-600 text-white">
      <router-view />
    </div>
  </div>
</template>