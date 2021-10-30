<template>
    <div id="dashboard">
        <div v-if="authState !== signedin">You're not allowed to see this page!</div>
        <div v-if="authState === signedin && user">Welcome to the dashboard!</div>
    </div>
</template>