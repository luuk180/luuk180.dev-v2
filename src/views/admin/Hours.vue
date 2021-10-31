<template>
    <div id="hours">
        <div v-if="authState !== 'signedin'">You are not allowed to see this page!</div>
        <div v-if="authState === 'signedin' && user">
            Welcome to the hours page!
        </div>
    </div>
</template>