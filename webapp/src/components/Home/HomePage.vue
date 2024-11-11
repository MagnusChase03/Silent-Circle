<script>
    import NavBar from '../NavBar.vue';
    import HomeContent from './HomeContent.vue';
    import SCLogo from '../SCLogo.vue';

    export default {
        name: 'MyHome',
        components: {
            NavBar,
            SCLogo,
            HomeContent
        }
    }
</script>

<template>
    <div id="wrapper-home">
        <NavBar></NavBar>
        <HomeContent></HomeContent>
        <SCLogo></SCLogo>
    </div>
</template>

<style scoped>
</style>
