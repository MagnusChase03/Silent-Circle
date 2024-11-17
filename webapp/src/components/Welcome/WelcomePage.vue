<script setup>
</script>

<template>
    <div id="wrapper">
        <div id="welcome-top">
            <div id="div-btn-menu">
                <button><img src="/src/assets/img/menu-btn-welcome.png"></button>
            </div>
        </div>
        <div id="welcome-main">
            <div id="welcome-left"></div>
            <div id="welcome-right">
                <div id="div-logo">
                    <div>
                        <img src="/src/assets/img/logo-welcome.png">
                    </div>
                    <h1>Silent Circle</h1>
                </div>
                <div id="div-login">
                    <div id="div-login-wrap">
                        <button id="btn-create">
                            <router-link to="/create-user">Create Account</router-link>
                        </button>
                        <p id="p-or">or</p>
                        <button id="btn-login">
                            <router-link to="/login">Log Me In</router-link>
                        </button>
                        <p id="p-forgot">Forgot Password?</p>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>
    #btn-login a{
        text-decoration: none;
        color:#FFF;
    }

    #btn-login a:hover{
        text-decoration: none;
        color:#000;
    }

    #btn-create a{
        text-decoration: none;
        color:#000;
    }

    #btn-create a:hover{
        text-decoration: none;
        color:#FFF;
    }
</style>
