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
                    <input id="uname" type="text" v-model="username" placeholder="Username" required>
                </div>
                <div id="div-pwd">
                    <input id="pwd" type="text" v-model="password" placeholder="Password" required>
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

<script >
import router from '@/router';
import { ref } from 'vue';
export default{
        name: 'LoginUser',
        setup(){
            // alert('LoginUser component loaded');
            console.log('LoginUser component loaded');
            // data
            const username = ref('root');
            const password = ref('supersecretpasswordhash');
            // methods
            const login = () => {
                fetch(import.meta.env.VITE_API_URL + "/login", {
                    // Send the username and password to the server
                    method: "POST",
                    headers: {
                        "Content-Type": "application/x-www-form-urlencoded"
                    },
                    body: `username=${uname.value}&password=${pwd.value}`,
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
                        console.log(data)
                        // Save the username in the local storage
                        localStorage.setItem('username', username.value);
                        // Redirect to the home page
                        alert("Login successful");
                        router.push('/home');
                    }
                    // If the response is not ok, throw an error
                }).catch((error) => console.error("Unable to tetch data:",error));
            }
            // computed
            return { username, password, login }
            
        }
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
