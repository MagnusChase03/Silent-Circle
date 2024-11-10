<template>
    <div id="wrapper-home">
        <NavBar></NavBar>

        <!-- Insert your content or any component here -->
        <div style="width: 100%; height: 100%;">
            <h1>This Is New Group Page Again after fix</h1>
        </div>
        <!-- End of main content -->

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