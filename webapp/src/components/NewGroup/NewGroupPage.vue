<template>
    <div id="wrapper-home">
        <NavBar></NavBar>
        <div id="home-right">

            <!-- Insert you main content here -->
            <h1>New Group Page</h1>

        </div>
        <SCLogo></SCLogo>
    </div>
</template>

<script>
    import NavBar from '../NavBar.vue';
    import SCLogo from '../SCLogo.vue';

    export default {
        name: 'NewGroupPage',
        components: {
            NavBar,
            SCLogo
        }
    }
</script>
<style>
</style>