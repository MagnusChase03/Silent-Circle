<template>
  <div id="wrapper-home">
    <NavBar></NavBar>
    <div class="page-background">
      <!-- Header with Menu Icon and Group Title -->
      <div class="header">
        <h1 class="group-title">{{ gname || "CS 4389 Group" }}</h1>
      </div>

      <!-- Chat Area -->
      <div class="chat-area">
        <!-- Dynamic Chat Messages -->
        <div
          v-for="(message, index) in messages"
          :key="index"
          :class="message.type === 'right' ? 'chat-message-right right' : 'chat-message left'"
        >
          <div v-if="message.type === 'left'" class="chat-message left">
            <img src="@/assets/img/user-blue.png" alt="Avatar" class="avatar" />
            <div class="bubble">
              <p class="username">{{ message.username }}</p>
              <p class="message">{{ message.content }}</p>
            </div>
          </div>
          <div v-else class="chat-message-right right">
            <div class="bubble">
              <p class="message">{{ message.content }}</p>
            </div>
          </div>
        </div>

        <!-- Typing Bar at the Bottom inside Chat Area -->
        <div class="typing-bar">
          <input
            v-model="newMessage"
            type="text"
            placeholder="Type a message..."
            class="message-input"
          />
          <button @click="sendMessage" class="send-button">Send</button>
        </div>

        <!-- Plus Sign Button at the Top Right -->
        <button class="plus-button">
          <router-link to="/add-member">+</router-link>
        </button>

        <!-- Minus Sign Button underneath Plus Button -->
        <button class="minus-button">
          <router-link to="/delete-user">-</router-link>
        </button>
      </div>
    </div>
    <SCLogo></SCLogo>
  </div>
</template>

<script>
import NavBar from "../NavBar.vue";
import SCLogo from "../SCLogo.vue";
import { ref, onMounted, onUnmounted } from "vue";
// Import the encryption and decryption functions
import useDecryptSymMsg from "../../composables/useDecryptSymMsg.js";
import useEncryptSymMsg from "../../composables/useEncryptSymMsg.js";

export default {
  props: ["gid", "gname"],
  name: "GroupChatPage",
  components: {
    NavBar,
    SCLogo,
  },
  setup(props) {
    const messages = ref([]);
    const newMessage = ref("");
    const username = localStorage.getItem("username");
    let socket = null;
    const groupSymmetricKeyTest = `${username}-${props.gid}`;
    const restoredKey = localStorage.getItem(groupSymmetricKeyTest);

    // Initialize WebSocket and set up event handlers
    const initializeWebSocket = () => {
      // const groupId = props.gid || "defaultGroup"; // Use defaultGroup if no gid provided
      var hostname = import.meta.env.VITE_API_URL;
      hostname = hostname.substring(hostname.indexOf("//") + 2);

      const wsUrl = `wss:${hostname}/group/chat?group=${props.gid}`;

      socket = new WebSocket(wsUrl);

      // Handle connection open
      socket.onopen = () => {
        console.log("WebSocket connection established.");
      };
      
      // Handle incoming messages
      socket.onmessage = async (event) => {
        const messageData = JSON.parse(event.data);
        try {
          // Extract the sender's username from the received message
          const msg_from_socket = messageData.Message;
          const sender_uame = msg_from_socket.substring(0, msg_from_socket.indexOf("-"));

          // check if sender is the receiving user
          if(sender_uame === username){
            return;
          }

          // Extract and decode the encrypted message from the received data
          const enccyptedMessage = decodeURIComponent(msg_from_socket.substring(msg_from_socket.indexOf("-") + 1),msg_from_socket.length);
          console.log("Received encrypted message:", enccyptedMessage);

          // Decrypt the received message
          const decryptedMessage = await useDecryptSymMsg(
            restoredKey,
            // messageData.message
            enccyptedMessage
          );
          
          messages.value.push({
            username: messageData.username,
            content: decryptedMessage,
            type: messageData.type, // Should be 'left' or 'right'
          });
        } catch (error) {
          console.error("Decryption failed:", error);
        }
      };

      // Handle connection errors
      socket.onerror = (error) => {
        console.error("WebSocket error:", error);
      };

      // Handle connection closure
      socket.onclose = () => {
        console.log("WebSocket connection closed.");
      };
    };

    // Send the encrypted message
    const sendMessage = async () => {
      if (!newMessage.value.trim()) return; // Avoid sending empty messages

      try {
        // Encrypt the message before sending
        const encryptedMessage = await useEncryptSymMsg(
          restoredKey,
          newMessage.value
        );

        // Send encrypted message to WebSocket
        socket.send(JSON.stringify({
          message: encodeURIComponent(encryptedMessage),
          group: props.gid
        }));

        // Push the message to the local array (for immediate display)
        messages.value.push({
          username: "You",
          content: newMessage.value, // Display original message for local preview
          type: "right",
        });

        // Clear the input field
        newMessage.value = "";
      } catch (error) {
        console.error("Encryption failed:", error);
      }
    };

    // Mount and unmount WebSocket connection
    onMounted(() => {
      initializeWebSocket();
    });

    onUnmounted(() => {
      if (socket) {
        socket.close();
      }
    });

    return {
      messages,
      newMessage,
      sendMessage,
    };
  },
};
</script>


<style scoped>
.page-background {
  display: flex;
  flex-direction: column;
  align-items: center;
  background: radial-gradient(circle at 50% 50%, #279ACD, #808FCD);
  overflow: hidden;
  height: 100%;
  width: 100%;
}

.header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
  padding: 20px;
  box-sizing: border-box;
  position: relative;
}


.group-title {
  font-family: cursive;
  font-size: 24px;
  color: white;
  text-align: center;
  flex-grow: 1;
}


.chat-area {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  width: 100%;
  height: 100%;
  max-width: 700px;
  max-height: 1000px;
  background-color: white;
  padding-left: 30px;
  padding-top: 30px;
  padding-bottom: 30px;
  padding-right: 15px;
  border-radius: 20px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  margin-top: 20px;
  margin-bottom: 20px;
  position: relative; /* Keeps the plus and minus buttons inside the chat area */
}

.chat-message {
  display: flex;
  align-items: flex-end;
  margin: 10px 0;
}

.chat-message-right {
  display: flex;
  justify-content: flex-end; /* Aligns the message to the right */
  margin: 10px 0; /* Add margin to separate messages */
  margin-left: 10px;
  width: 90%;
}

.left .bubble {
  background-color: #decff7;
  color: black;
  align-self: flex-start;
  border-radius: 15px 15px 15px 0;
}

.right .bubble {
  background-color: #b79edc;
  color: black;
  align-self: flex-end;
  border-radius: 15px 15px 0 15px;
  max-width: 80%; /* Limits the width of the bubble */
  /*text-align: left;*/
}

.avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  margin-right: 10px;
}

.bubble {
  padding: 10px 15px;
  max-width: 70%;
}

.username {
  font-weight: bold;
  margin-bottom: 5px;
  font-size: 14px;
}

.message {
  font-size: 16px;
}

.typing-bar {
  display: flex;
  align-items: center;
  width: 100%;
  max-width: 705px;
  padding: 10px;
  background-color: #f3f3f3;
  border-top: 1px solid #ddd;
  border-radius: 20px;
  box-shadow: 0 -2px 4px rgba(0, 0, 0, 0.1);
  position: absolute;
  bottom: 0; /* Keeps it at the bottom inside the chat area */
  left: 0;
  padding-left: 20px;
  padding-right: 20px;
}

.message-input {
  flex: 1;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 20px;
  font-size: 16px;
}

.send-button {
  padding: 10px 15px;
  background-color: #279acd;
  color: white;
  border: none;
  border-radius: 20px;
  font-size: 16px;
  margin-left: 10px;
  cursor: pointer;
}

.send-button:hover {
  background-color: #1e7cb8;
}

.plus-button {
  position: absolute;
  top: 20px;
  right: 20px;
  background-color: #279acd;
  color: white;
  border: none;
  border-radius: 50%;
  width: 40px;
  height: 40px;
  font-size: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  /*box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);*/
}

.plus-button:hover {
  background-color: #1e7cb8;
}

/* Minus Button Styling */
.minus-button {
  position: absolute;
  top: 70px; /* Adjust this to position the minus button below the plus button */
  right: 20px;
  background-color: #e74c3c;
  color: white;
  border: none;
  border-radius: 50%;
  width: 40px;
  height: 40px;
  font-size: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  /*box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);*/
}

.minus-button:hover {
  background-color: #c0392b;
}
</style>
