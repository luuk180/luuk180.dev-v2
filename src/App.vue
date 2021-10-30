<template>
  <div id="app">
    <div v-if="authState!== 'signedin'">
      <div id="nav" class="bg-fixed w-screen space-x-4 text-center text-gray-800 bg-blue-500 bg-opacity-80 text-xl p-2 h-12">
        <router-link class="bg-white p-1.5 rounded" to="/">Home</router-link>
        <router-link class="bg-white p-1.5 rounded" to="/projects">Projects</router-link>
        <router-link class="bg-white p-1.5 rounded" to="/cv">CV</router-link>
        <router-link class="bg-white p-1.5 rounded" to="/about">About</router-link>
        <router-link class="bg-white p-1.5 rounded right-1" to="/login">Login</router-link>
      </div>
    </div>
    <div v-if="authState == 'signedin' && user">
      <div id="nav" class="bg-fixed w-screen space-x-4 text-center text-gray-800 bg-blue-500 bg-opacity-80 text-xl p-2 h-12">
        <router-link class="bg-white p-1.5 rounded" to="/admin">Home</router-link>
        <router-link class="bg-white p-1.5 rounded" to="/dashboard">Dashboard</router-link>
        <router-link class="bg-white p-1.5 rounded" to='/hours'>Hours</router-link>

      </div>
    </div>
    <div id="container" class="bg-grey-200">
      <router-view/>
    </div>
  </div>
</template>