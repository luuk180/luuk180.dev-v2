<template>
    <div id="admin">
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