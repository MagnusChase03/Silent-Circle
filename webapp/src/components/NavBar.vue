<template>
    <div id="home-left">
        <div id="home-menu-top">
            <div class="div-h-user">
                    Hello
                    <span class="h-user">{{ username }}</span>
            </div>
            <!-- <button><img src="../assets/img/arrow-left-ico.png" alt="Go back"></button> -->
        </div>
        <div id="home-new-group">
            <button>
                <router-link to="/new-group">
                    <img src="../assets/img/add-group-ico.png" alt="New Group">
                </router-link>
            </button>
            <h5>New Group</h5>
        </div>
        <div id="home-invite-user">
            <button>
                <router-link to="/invite-users">
                    <img src="../assets/img/invite-user-ico.png" alt="Invite Users">
                </router-link>
            </button>
            <h5>Invite Users</h5>
        </div>
        <div id="home-btn">
            <button>
                <router-link to="/home">
                    <img src="../assets/img/home-ico.png" alt="Go back Home">
                </router-link>
            </button>
            <h5>Home</h5>
        </div>
        <div id="home-logout">
            <button><img src="../assets/img/log-out-ico.png" alt="Log out" @click.prevent="logOut">&nbsp;</button>
            <h5>Log out</h5>
        </div>
    </div>
</template>

<script>
    import { ref, onMounted } from 'vue';
    import router from '@/router';
    export default {
        name: 'NavBar',
        setup(){
            // data
            const privateKey = ref(localStorage.getItem('privateKey'));
            const publicKey = ref('');

            const username = ref(localStorage.getItem('username'));

            // methods
            const logOut = () => {
                fetch(import.meta.env.VITE_API_URL + "/logout", {
                    // Send the username and password to the server
                    method: "POST",
                    headers: {
                        "Content-Type": "application/x-www-form-urlencoded"
                    },
                    credentials: "include"
                }).then((res) => {
                    if(!res.ok){
                        // If the response is not ok, throw an error
                        throw new Error(`Http error! Status: ${res.status}`);
                    }
                    // Return the response as JSON
                    return res.json();

                }).then((data) => {
                    if(data.StatusCode==200){
                        // Remove the username from the local storage
                        localStorage.removeItem('username');
                        alert("Logout successful");
                        // Redirect to the welcome page
                        router.push('/');
                    }
                    // If the response is not ok, throw an error
                }).catch((error) => console.error("Unable to tetch data:",error));
            }

            // lifecycle hooks
            onMounted(() => {
                // If the username is not set, redirect to the welcome page
                if (!username.value || username.value === 'null') {
                    router.push('/');
                }
            });
            
            // computed
            return { username, logOut}
        }
    }
</script>

<style scoped>
    div#home-menu-top .div-h-user{
        color:#FFF;
        font-size: 1.6em;
        text-align: left;
        height: auto;
        margin:5px
    }

    div#home-menu-top .h-user{
        display: inline-block;
        font-size: 1.7em;
        font-weight: 900;
    }

    div#home-menu-top .btn-log-out-user{
        /* display: block; */
        color: rgb(6, 6, 88);
        text-decoration: underline;
        float: left !important;
        margin:5px
    }

    div#home-menu-top .btn-log-out-user:hover{
        color:#FFF;
        cursor: pointer;
    }

    div#home-logout, div#home-new-group, div#home-invite-user, div#home-btn, div#home-logout button:hover{
        color:#FFF;
        cursor: pointer;
    }

    div#home-menu-top button img {
        width: 24px;
        height: 24px;
    }

    div#home-menu-top{
        position:relative
    }

    div#home-menu-top button{
        position: absolute;
        top: 0px;
        right: 0px;
    }
</style>
