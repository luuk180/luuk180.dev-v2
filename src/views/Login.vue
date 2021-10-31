<template>
    <div id="login">
        <amplify-auth-container>
            <amplify-authenticator>
                <div v-if="authState === 'signedin' && user">
                    <div>Hello, {{user.username}}</div>
                    <router-link to="/admin/dashboard">Dashboard</router-link>
                </div>
                <amplify-sign-out></amplify-sign-out>
            </amplify-authenticator>
        </amplify-auth-container>
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