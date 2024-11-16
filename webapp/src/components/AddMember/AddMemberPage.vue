<template>
  <div id="wrapper-home">
    <!-- NavBar Component -->
    <NavBar></NavBar>
    <div style="width: 100%; height: 100%;">
      <!-- Add Member Page Content -->
      <div class="page-container">
        <div class="page-background">
          <div class="header">
            <h1 class="title">Invite User</h1>
          </div>
          <div class="right-side">
            <div class="chat-bubble">
              <div class="group-name-box">
                <input type="text" v-model="username" placeholder="Enter username" class="user-name-input" @keyup.enter="invite" />
              </div>
              <br />
              <!-- Invite user button -->
              <button class="invite-button" @click="invite" :disabled="!username.trim()">
                Invite user
              </button>
              <!-- Link to go back to homepage -->
              <router-link to="/home" class="back-home-link">Go back to homepage</router-link>
            </div>
          </div>
          <img src="@/assets/img/create-grp-photo.gif" alt="Animated GIF" class="animated-gif" />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import NavBar from '../NavBar.vue';
import useEncryptAsymMsg from '../../composables/useEncryptAsymMsg.js';  // Import the encryption function
import { ref } from 'vue';

export default {
  name: 'AddMemberPage',
  components: {
    NavBar,
  },
  data() {
    return {
      username: '',  // Username input for the invite
      publicKey: ref(''),  // Public key of the invitee
      privateKey: ref(localStorage.getItem('privateKey')),  // Get invitor's private key from localStorage
      groupid: this.$route.query.groupid || '101',  // Get the group ID from the query params
      groupSymmetricKey: ref(''), // Group symmetric key (to be retrieved from localStorage)
    };
  },
  created() {
    this.loadGroupSymmetricKey(); // Load the group symmetric key when the component is created
  },
  methods: {
    async loadGroupSymmetricKey() {
      const username = localStorage.getItem('username');
      const groupSymmetricKey = localStorage.getItem(`${username}-${this.groupid}`);
      if (groupSymmetricKey) {
        this.groupSymmetricKey = groupSymmetricKey;
        console.log("Group symmetric key retrieved:", this.groupSymmetricKey); // Debug statement
      } else {
        console.error("Group symmetric key not found.");
        alert("Unable to retrieve group symmetric key.");
      }
    },

    async fetchPublicKey() {
      try {
        const response = await fetch(import.meta.env.VITE_API_URL + "/user/get", {
          method: "POST",
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
          body: `username=${this.username}`,
          credentials: "include",
        });

        if (!response.ok) {
          throw new Error(`Http error! Status: ${response.status}`);
        }

        const data = await response.json();

        if (data.StatusCode === 200) {
          this.publicKey = data.Data.PublicKey;
          console.log("Public Key Retrieved:", this.publicKey); // Debug statement
        } else {
          throw new Error('Failed to fetch public key');
        }
      } catch (error) {
        console.error("Unable to fetch public key:", error);
        alert("Unable to fetch public key. Please try again.");
      }
    },

    async invite() {
      if (!this.username.trim()) {
        alert('Please enter a username.');
        return;
      }

      // Fetch the public key of the invitee
      await this.fetchPublicKey();

      if (!this.publicKey) {
        alert("Could not retrieve the invitee's public key.");
        return;
      }

      // Encrypt the group symmetric key using the invitee's public key
      const encryptedGroupSymmetricKey = await this.encryptGroupKey(this.publicKey);

      if (!encryptedGroupSymmetricKey) {
        alert("Failed to encrypt the group symmetric key.");
        return;
      }

      // Send the group invite
      this.sendGroupInvite(encryptedGroupSymmetricKey);
    },

    async encryptGroupKey(publicKey) {
      try {
        console.log("Encrypting group symmetric key:", this.groupSymmetricKey); // Debug statement
        // Encrypt the group symmetric key using the invitee's public key
        const encryptedMessage = await useEncryptAsymMsg(publicKey, this.groupSymmetricKey);
        console.log("Encrypted Group Symmetric Key:", encryptedMessage); // Debug statement
        return encryptedMessage;  // Return the encrypted group symmetric key
      } catch (error) {
        console.error("Error encrypting group symmetric key:", error);
        return null;
      }
    },

    async sendGroupInvite(encryptedGroupSymmetricKey) {
      try {
        const response = await fetch(import.meta.env.VITE_API_URL + "/group/invite", {
          method: "POST",
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
          body: `username=${this.username}&key=${encodeURIComponent(encryptedGroupSymmetricKey)}&group=${this.groupid}`,
          credentials: "include",
        });

        if (!response.ok) {
          throw new Error(`Http error! Status: ${response.status}`);
        }

        const data = await response.json();

        if (data.StatusCode === 200) {
          alert("Group invitation sent successfully!");
          this.$router.push('/home');  // Navigate back to the homepage
        } else {
          alert("Failed to send group invitation.");
        }
      } catch (error) {
        console.error("Error sending group invite:", error);
        alert("An error occurred while sending the group invitation.");
      }
    }
  },
};
</script>

<style scoped>
  .page-container {
    display: flex;
    height: 100vh;
  }
  .page-background {
    flex-grow: 1;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    background: none;
    position: relative;
    overflow: hidden;
  }

  .header {
    position: absolute;
    top: 30px;  
    left: 30px; 
    display: flex;
    align-items: center;
  }
  .title {
    font-family: cursive;
    font-size: 50px;
    color: #808fcd;
    margin: 0;
  }
  .right-side {
    display: flex;
    flex-direction: column;
    align-items: flex-end;
    width: 600px;
  }
  .chat-bubble {
    position: absolute;
    top: 10px;
    right: 50px;
    width: 600px;
    height: 500px;
    background-image: url('@/assets/img/bubble-chat-create-group.png');
    background-size: cover;
    background-repeat: no-repeat;
  }

  .group-name-box {
    background: radial-gradient(circle, #d9fffb, #ccebff);
    width: 100%;
    max-width: 350px;
    border-radius: 25px;
    padding: 10px;
    margin-bottom: 10px;
    margin-top: 55px;
    margin-left: 130px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  }
  .user-name-input {
    background-color: transparent;
    border: none;
    outline: none;
    padding: 10px;
    font-size: 16px;
    color: #333;
    border-radius: 25px;
  }
  .user-name-input::placeholder {
    color: #666;
  }
  .invite-button {
    display: flex;
    align-items: center;
    justify-content: center;
    background: radial-gradient(circle, #279acd, #808fcd 150px);
    border: none;
    border-radius: 25px;
    width: 100%; 
    max-width: 370px;
    margin-left: 130px;
    padding: 20px; 
    cursor: pointer;
    font-size: 16px;
    text-align: center;
    color: white;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  }
  .animated-gif {
    position: absolute;
    bottom: 10px;
    left: 50%;
    transform: translateX(-50%);
    width: 400px;
    height: auto;
  }

  /* Style for the "Go back to homepage" link */
  .back-home-link {
    display: block;
    margin-top: 20px;
    font-size: 16px;
    color: white;
    text-align: center;
    text-decoration: underline;
    cursor: pointer;
  }

  .back-home-link:hover {
    color: #808fcd;
  }
</style>
