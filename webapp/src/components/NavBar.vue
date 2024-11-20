<template>
    <div id="home-left">
        <div id="home-menu-top">
            <div class="div-h-user">
                    Hello
                    <span class="h-user">{{ username }}</span>
            </div>
            <router-link to="/accept-invite" class="notify-invite">
                <button class="btn-notify" v-show="isVisible">
                    <img src="../assets/img/notify-ico.png" alt="New Invite">
                </button>
            </router-link>
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
        <div id="home-invite-user">
            <button>
                <router-link to="/accept-invite">
                    <img src="../assets/img/invite-ico.png" alt="Accept Invitation">
                </router-link>
            </button>
            <h5>Accept Invitates</h5>
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
    import { ref, onMounted, onUnmounted } from 'vue';
    import router from '@/router';
    export default {
        name: 'NavBar',
        setup(){
            // data
            let socket = null;
            const isVisible = ref(false);
            const privateKey = ref(localStorage.getItem('privateKey'));
            const publicKey = ref('');
            const username = ref(localStorage.getItem('username'));

            const initializeWebSocket = () => {
                var hostname = import.meta.env.VITE_API_URL;
                hostname = hostname.substring(hostname.indexOf("//") + 2);

                const wsUrl = `wss:${hostname}/group/invite/listen`;

                socket = new WebSocket(wsUrl);

                socket.onopen = () => {
                    console.log("WebSocket connection established.");
                };

                socket.onmessage = async (event) => {
                    const messageData = JSON.parse(event.data);
                    isVisible.value = true;
                }

                // Handle connection errors
                socket.onerror = (error) => {
                    console.error("WebSocket error:", error);
                };

                // Handle connection closure
                socket.onclose = () => {
                    console.log("WebSocket connection closed.");
                };
            }


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

                initializeWebSocket();
            });

            onUnmounted(() => {
                if (socket) {
                    socket.close();
                }
            });
            
            // computed
            return { username, logOut, isVisible}
        }
    }
</script>

<style scoped>
    div#home-menu-top .div-h-user{
        color:#FFF;
        font-size: 1.5em;
        text-align: left;
        height: auto;
        margin:5px
    }

    div#home-menu-top .h-user{
        display: inline-block;
        font-size: 1.7em;
        font-weight: 900;
    }

    div#home-menu-top .notify-invite{
        /* display: block; */
        color: rgb(6, 6, 88);
        text-decoration: underline;
        float: left !important;
        margin:5px
    }

    div#home-menu-top .notify-invite .btn-notify:hover{
        cursor: pointer;
    }

    div#home-logout, div#home-new-group, div#home-invite-user, div#home-btn, div#home-logout button:hover{
        color:#FFF;
        cursor: pointer;
    }

    div#home-menu-top button img {
        width: 32px;
        height: 32px;
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
