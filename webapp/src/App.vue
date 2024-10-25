<script setup>
// import Navbar from './components/Navbar.vue';
// import LoginPage from './components/user/Login.vue';
import Welcome from './components/Welcome.vue';

import { ref } from 'vue';

let currentPage = ref("login");
</script>

<template>
<!-- <Navbar/>
<LoginPage v-if="currentPage == 'login'" /> -->
<h1>Hello World from App</h1>
<Welcome></Welcome>
</template>

<style scoped>
</style>
