<template>
  <div id="app">
    <div id="nav" class="bg-fixed w-screen space-x-4 text-center text-gray-800 bg-blue-500 bg-opacity-80 text-2xl p-2 h-12">
        <router-link class="bg-white p-1.5 rounded" to="/">Home</router-link>
        <router-link class="bg-white p-1.5 rounded" to="/projects">Projects</router-link>
        <router-link class="bg-white p-1.5 rounded" to="/cv">CV</router-link>
        <router-link class="bg-white p-1.5 rounded" to="/about">About</router-link>
    </div>
    <div id="container" class="absolute inset-x-0 top-12 bottom-0 bg-blue-200">
      <router-view />
    </div>
  </div>
</template>