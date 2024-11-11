<template>
    <div id="wrapper-home">
        <NavBar></NavBar>
        <div id="home-right">
            <!-- Insert you main content here -->
            <h1>Group Chat Page with Group Id: {{ gid }} - Group Name: {{ gname }}</h1>
        </div>
        <SCLogo></SCLogo>
    </div>
</template>

<script>
    import NavBar from '../NavBar.vue';
    import SCLogo from '../SCLogo.vue';

    export default {
        props: ['gid','gname'],
        name: 'GroupChatPage',
        components: {
            NavBar,
            SCLogo
        }
    }

</script>

<style>

</style>