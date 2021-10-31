<template>
  <div id="app">
    <div v-if="authState !== 'signedin'">
      <div id="nav" class="bg-fixed w-screen space-x-4 text-center text-gray-800 bg-blue-500 bg-opacity-80 text-xl p-2 h-12">
        <router-link class="bg-white p-1.5 rounded" to="/">Home</router-link>
        <router-link class="bg-white p-1.5 rounded" to="/projects">Projects</router-link>
        <router-link class="bg-white p-1.5 rounded" to="/cv">CV</router-link>
        <router-link class="bg-white p-1.5 rounded" to="/about">About</router-link>
        <router-link class="bg-white p-1.5 rounded" to="/login">Login</router-link>
      </div>
    </div>
    <div v-if="authState === 'signedin' && user">
      <div id="nav" class="bg-fixed w-screen space-x-4 text-center text-gray-800 bg-blue-500 bg-opacity-80 text-xl p-2 h-12">
        <router-link class="bg-white p-1.5 rounded" to="/admin/dashboard">Dashboard</router-link>
        <router-link class="bg-white p-1.5 rounded" to='/admin/hours'>Hours</router-link>
        <button @click="handleSignOut" class="bg-white p-1.5 rounded w-6 right-4">Sign out</button>
      </div>
    </div>
    <div id="container" class="bg-grey-200">
      <router-view/>
    </div>
  </div>
</template>

<script>
import { onAuthUIStateChange } from '@aws-amplify/ui-components'
import { Auth } from 'aws-amplify'

export default {
  name: 'AuthStateApp',
  created() {
    this.unsubscribeAuth = onAuthUIStateChange((authState, authData) => {
      this.authState = authState;
      this.user = authData;
    })
  },
  data() {
    return {
        user: undefined,
        authState: undefined,
        unsubscribeAuth: undefined,
        formFields: [
            { type: "username" },
            { type: "password" },
            { type: "email" }
        ],
    }
  },
  beforeUnmount() {
    this.unsubscribeAuth();
  },
  methods: {
    async handleSignOut() {
      try {
        await Auth.signOut({global: true})
      }
      catch(error) {
        console.log("error signing out: ", error)
      }
    }
  }
}
</script>