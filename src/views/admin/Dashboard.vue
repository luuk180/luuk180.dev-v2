<template>
    <div id="dashboard">
        <div v-if="authState !== 'signedin'">You're not allowed to see this page!</div>
        <div v-if="authState === 'signedin' && user">
            Welcome to the dashboard!
        </div>
    </div>
</template>

<script>
import { onAuthUIStateChange } from '@aws-amplify/ui-components'

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
  }
}
</script>