<template>
    <div class="login-container">
      <div class="icon-chat"></div> <!-- Top-right icon -->
      <div class="icon-person-desktop"></div> <!-- Center-right icon -->
      
      <div class="login-logo">
        <!-- Logo image -->
      </div>
      <h1>Create User Account</h1>
      
  
      <form @submit.prevent="handleLogin">
        <div class="input-group">
          <input type="text" placeholder="Username" v-model="username" class="icon-username" />
        </div>
        <div class="input-group">
          <input type="text" placeholder="Phone" v-model="phone" class="icon-phone" />
        </div>
        <div class="input-group">
          <input type="text" placeholder="Choose Profile Picture" v-model="profilePicture" class="icon-profile-picture" />
        </div>
        <div class="input-group">
          <input type="email" placeholder="Email" v-model="email" class="icon-email" />
        </div>
        <div class="input-group">
          <input type="password" placeholder="Password" v-model="password" class="icon-password" />
        </div>
        <div class="input-group">
          <input type="password" placeholder="Confirm Password" v-model="confirmPassword" class="icon-password" />
        </div>
        <button type="submit">Create User</button>
        <div class="have-account">
          <p class="status-bar" :style="{ color: statusBarColor }">{{ statusBarMessage }}</p>
          <p class="">Already have an account: <a href="#" class="login">Log in</a></p>
        </div>
      </form>
    </div>
  </template>
  
  <script>
  import { ref } from 'vue';
  import { useRouter } from 'vue-router';
  export default {

    setup() {
        const username = ref('');
        const password = ref('');
        const publicKey = ref(null);
        const router = useRouter();

        async function generateKeyPair() {
            const keyPair = await window.crypto.subtle.generateKey(
                {
                    name: "RSA-OAEP",
                    modulusLength: 2048,
                    publicExponent: new Uint8Array([1, 0, 1]),
                    hash: "SHA-256"
                },
                true,
                ["encrypt", "decrypt"]
            );

            const exportedPublicKey = await window.crypto.subtle.exportKey(
                "spki",
                keyPair.publicKey
            );

            publicKey.value = exportedPublicKey;
        }

        async function createUser() {
            await generateKeyPair();

            const publicKeyBase64 = btoa(String.fromCharCode(...new Uint8Array(publicKey.value)));

            const response = await fetch('/api/create-user', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    username: username.value,
                    password: password.value,
                    publicKey: publicKeyBase64
                })
            });

            if (response.ok) {
                router.push('/login');
            } else {
                console.error('Failed to create user');
            }
        }

        return {
            username,
            password,
            createUser
        };
    },
    data() {
      return {
        username: 'Test name',
        phone: '123-456-7891',
        email: 'test@gmail.com',
        password: 'Testing@12345',
        confirmPassword: 'Testing@12345',
        profilePicture: '/assets/img/chat_icon.png',
        statusBarMessage: '', // For displaying messages
        statusBarColor: '', // For setting message color
        profilePicture: '/assets/img/username_icon.png',
      };
    },
  
    methods: {
      validatePassword() {
        const passwordRegex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$/;
  
        if (!passwordRegex.test(this.password)) {
          this.statusBarMessage = "Password must be at least 8 characters long, contain uppercase, lowercase, number, and a special character.";
          this.statusBarColor = "red";
          return false;
        }
  
        if (this.password !== this.confirmPassword) {
          this.statusBarMessage = "Passwords do not match.";
          this.statusBarColor = "red";
          return false;
        }
  
        // Clear any previous messages if validation is successful
        this.statusBarMessage = "";
        return true;
      },
  
      async handleLogin() {
        // Run validation first
        if (!this.validatePassword()) {
          return; // Stop if validation fails
        }
        
        try {
          const response = await fetch("https://localhost:8081/user/create", {
            method: "POST",
            headers: {
              "Content-Type": "application/json"
            },
            body: JSON.stringify({
              username: this.username,
              password: this.password,
              publicKey: this.publicKey // Ensure this property is defined in your data
            })
          });
          
          const data = await response.json();
          
          if (data.StatusCode === 200) {
            this.statusBarMessage = "User created successfully!";
            this.statusBarColor = "green"; // Set text color to green for success
          } else {
            this.statusBarMessage = "Failed to create user. " + data.Data;
            this.statusBarColor = "red"; // Set text color to red for failure
          }
        } catch (error) {
          console.error("Error creating user:", error);
          this.statusBarMessage = "An error occurred. Please try again.";
          this.statusBarColor = "red";
        }
      }
    }
  };
  </script>
  
  <style>
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
  