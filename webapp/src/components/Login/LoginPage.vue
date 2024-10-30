<template>
    <div id="wrapper">
        <div id="login-logo">
            <div id="div-logo">
            </div>
            <h1>Slient Circle</h1>
        </div>
        <div id="login-main">
            <div id="div-login-form">
                <h1> Welcome Back!</h1>
                <div id="div-uname">
                    <input id="uname" type="text" placeholder="Username" value="root">
                </div>
                <div id="div-pwd">
                    <input id="pwd" type="text" placeholder="Password" value="supersecretpasswordhash">
                </div>
                <div id="div-forgot">
                    <a href="#">Forgot passsword</a>
                </div>
                <div id="div-login">
                    <button id="btn-login" @click.prevent="login">Login</button>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import router from '@/router';

function login() {
    fetch(import.meta.env.VITE_API_URL + "/login", {
        method: "POST",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded"
        },
        body: `username=${uname.value}&password=${pwd.value}`,
        credentials: "include"
    }).then((res) => {
        if(!res.ok){
            throw new Error(`Http error! Status: ${res.status}`);
        }
        return res.json();

    }).then((data) => {
        if(data.StatusCode==200){
            console.log(data)
            router.push('/home');
        }
    }).catch((error) => console.error("Unable to tetch data:",error));
}
</script>

<style scoped>
#div-login-form input {
    width: auto;
    min-width: auto;
    color: #fff;
    font-size: 3vw;
}
</style>
