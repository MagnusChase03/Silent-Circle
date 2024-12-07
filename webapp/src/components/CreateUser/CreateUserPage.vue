<template>
  <div id="wrapper">
    <div class="login-container">
      <div class="icon-chat"></div> <!-- Top-right icon -->
      <div class="icon-person-desktop"></div> <!-- Center-right icon -->

      <div class="login-logo">
        <!-- Logo image -->
      </div>
      <h1>Create User Account</h1>

      <form @submit.prevent="createUser">
        <div class="input-group">
          <input type="text" placeholder="Username" v-model="username" class="icon-username" required />
        </div>
        <div class="input-group">
          <input type="password" placeholder="Password" v-model="password" class="icon-password" required />
        </div>
        <div class="input-group">
          <input type="password" placeholder="Confirm Password" v-model="confirmPassword" class="icon-password" required/>
        </div>
        <button type="submit">Create User</button>
        <div class="have-account">
          <p class="status-bar" :style="{ color: statusBarColor }">{{ statusBarMessage }}</p>
          <p class="">Already have an account:  
            <router-link to="/login" class="login">Log in</router-link>
          </p>
        </div>
      </form>
    </div>
  </div>

</template>
  
  <script>
  // import { j } from 'vite/dist/node/types.d-aGj9QkWt';
  import { ref } from 'vue';
  import { useRouter } from 'vue-router'
  import useAsymmetricKeys from '../../composables/useAsymmetricKeys.js';

  export default {
    name: 'CreateUserPage',
    setup() {
      // data
      const statusBarMessage = ref(''); // For displaying messages
      const statusBarColor = ref(''); // For setting message color
      const username = ref('');
      const password = ref('');
      const confirmPassword = ref('');
      const publicKeyBase64 = ref('');
      const privateKeyBase64 = ref('');
      const router = useRouter();

      function validatePassword() {
        // commented out for easier testing
        // const passwordPattern = /^(?=.*[0-9])(?=.*[!@#$%^&*])[A-Za-z\d!@#$%^&*]{8,}$/;
        // if (!passwordPattern.test(password.value)) {
        //   statusBarMessage.value = "Password must be at least 8 characters long, contain at least one number, and one special character.";
        //   statusBarColor.value = "red";
        //   return false;
        // }

        if (password.value != confirmPassword.value) {
          statusBarMessage.value = "Passwords do not match.";
          statusBarColor.value = "red";
          return false;
        }

        // Clear any previous messages if validation is successful
        statusBarMessage.value = "";
        return true;
      }

      // methods
      async function createUser() {
        // Run validation first
        if (!validatePassword()) {
          return; // Stop if validation fails
        }

        // Generate a new key pair
        const keyPair = await useAsymmetricKeys();
        publicKeyBase64.value = keyPair.publicKeyBase64;
        privateKeyBase64.value = keyPair.privateKeyBase64;

        await fetch(import.meta.env.VITE_API_URL + "/user/create", {
          method: 'POST',
          headers: {
            "Content-Type": "application/x-www-form-urlencoded"
          },
          body: `username=${username.value}&password=${password.value}&publicKey=${encodeURIComponent(publicKeyBase64.value)}`,
        }).then((res) => {
          if(!res.ok){
            statusBarMessage.value = "Cannot create user account, username already exists.";
            statusBarColor.value = "red";

            // If the response is not ok, throw an error
            throw new Error(`Http error! Status: ${res.status}`);
          }
          // Return the response as JSON
          return res.json();
        }).then((data) => {
          if(data.StatusCode==200){
            // log the data to the console
            console.log(data)
            console.log("Public Key Initialized:", publicKeyBase64.value);
            console.log("Private Key Initialized:", privateKeyBase64.value);

            // Save the private key in the local storage
            localStorage.setItem('privateKey', privateKeyBase64.value);

            // Redirect to the home page
            alert("Create account successful");
            router.push('/login');
          }
          // If the response is not ok, throw an error
        }).catch((error) => console.error("Unable to tetch data:",error));
      }
      return {
        statusBarMessage,
        statusBarColor,
        username,
        password,
        confirmPassword,
        createUser
      };
    }
  };
  </script>
  
  <style scoped>
  :root {
    --title-color: white;
    --primary-color: #7b8fcd;
    --secondary-color: #5095cd;
    --border-color: #a9c8cf;
    --text-color: var(--border-color);
    --background-color: #fdfefe;
    --button-hover-color: #3498db;
    --on-focus: white;
  
    --btn-primary-color: #879acd;
    --btn-secondary-color: #3198cd;
    --btn-on-focus-color: #04a3f3;
  }
  
  .login-container {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: center;
    height: 100vh;
    background: linear-gradient(to right, var(--primary-color), var(--secondary-color), var(--primary-color));
    font-family: Arial, sans-serif;
    color: var(--text-color);
    padding-left: 5%;
    position: relative;
  }
  
  /* Top-right chat icon */
  .icon-chat {
    position: absolute;
    top: 5%;
    right: 2%;
    width: 150px;
    height: 150px;
    background-image: url('@/assets/img/chat_icon.png');
    background-size: contain;
    background-repeat: no-repeat;
  }
  
  /* Center-right person icon */
  .icon-person-desktop {
    position: absolute;
    top: 50%;
    left: 35%;
    transform: translateY(-50%);
    width: 500px;
    height: 500px;
    background-image: url('@/assets/img/person_on_deskt.png');
    background-size: contain;
    background-repeat: no-repeat;
  
  }
  
  h1 {
    font-size: 2rem;
    color: var(--title-color);
    font-weight: bold;
    margin-bottom: 20px;
  }
  
  form {
    width: 100%;
    max-width: 300px;
    display: flex;
    flex-direction: column;
  }
  
  .input-group {
    position: relative;
    width: 100%;
    margin-bottom: 15px;
    border: 2px solid var(--border-color);
    border-radius: 20px;
  }
  .input-group:hover{
    background: var(--btn-on-focus-color);
    color: var(--background-color);
    border-color: var(--background-color);
  }
  
  /* Styles for input icons */
  .icon-username, .icon-phone, .icon-email, .icon-password, .icon-profile-picture {
    border: none;
    padding: 10px 10px 10px 50px;
    width: 100%;
    font-size: 1rem;
    outline: none;
    background: none;
    color: var(--text-color);
    background-size: 20px;
    background-repeat: no-repeat;
    background-position: 10px center;
    transition: color 0.3s ease, border-color 0.3s ease;
  }
  
  .input-group:hover .icon-username,
  .input-group:hover .icon-phone,
  .input-group:hover .icon-email,
  .input-group:hover .icon-password,
  .input-group:hover .icon-profile-picture {
    color: var(--on-focus); /* Change text color on hover */
    filter: brightness(0) invert(1); /* Make the icon appear white */
  }
  .icon-username {
    background-image: url('@/assets/img/username_icon.png');
  }
  
  .icon-phone {
    background-image: url('@/assets/img/phone_icon.png');
  }
  
  .icon-email {
    background-image: url('@/assets/img/email_icon.png');
  }
  
  .icon-password {
    background-image: url('@/assets/img/password_icon.png');
  }
  
  .icon-profile-picture {
    background-image: url('@/assets/img/profile_picture_icon.png');
    cursor: pointer;
    display: block;
  }
  
  
  
  input::placeholder {
    color: var(--text-color);
    opacity: 1;
  }
  
  button {
    width: 100%;
    padding: 10px;
    border: none;
    border-radius: 20px;
    font-size: 1rem;
    font-weight: bold;
    color: var(--text-color);
    background: linear-gradient(to right, var(--btn-primary-color), var(--btn-secondary-color), var(--btn-primary-color));
    cursor: pointer;
    transition: background 0.3s;
    border: 2px solid var(--text-color);
  }
  
  button:hover {
    background: var(--btn-on-focus-color);
    color: var(--background-color);
    border-color: var(--background-color);
  }
  
  .have-account{
    margin: 10px;
  }
  .login{
    color: #a9c8cf;
  }
  .login:hover{
    color: var(--on-focus);
    font-weight: bold;
  }
  </style>
  