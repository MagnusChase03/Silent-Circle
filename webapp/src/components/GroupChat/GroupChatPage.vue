<template>
  <div id="wrapper-home">
    <NavBar></NavBar>
    <div class="page-background">
      <!-- Header with Menu Icon and Group Title -->
      <div class="header">
        <h1 class="group-title">CS 4389 Group</h1>
      </div>

      <!-- Chat Area -->
      <div class="chat-area">
        <!-- Chat Bubble - Message 1 -->
        <div class="chat-message left">
          <img src="@/assets/img/user-blue.png" alt="Avatar" class="avatar" />
          <div class="bubble">
            <p class="username">comet_99</p>
            <p class="message">My name is Temoc, I'm a senior!</p>
          </div>
        </div>

        <!-- Chat Bubble - Message 2 -->
        <div class="chat-message left">
          <img src="@/assets/img/user-blue.png" alt="Avatar" class="avatar" />
          <div class="bubble">
            <p class="username">crane_981</p>
            <p class="message">Hi! I’m Enarc, I’m a junior this year!</p>
          </div>
        </div>

        <!-- Chat Bubble - Message 3 -->
        <div class="chat-message-right right">
          <div class="bubble">
            <p class="message">Let's briefly introduce ourselves!</p>
          </div>
        </div>

        <!-- Chat Bubble - Message 4 -->
        <div class="chat-message-right right">
          <div class="bubble">
            <p class="message">My name is Name! Also a senior!</p>
          </div>
        </div>

        <!-- Typing Bar at the Bottom inside Chat Area -->
        <div class="typing-bar">
          <input
            type="text"
            placeholder="Type a message..."
            class="message-input"
          />
          <button class="send-button">Send</button>
        </div>

        <!-- Plus Sign Button at the Top Right -->
        <button class="plus-button">+</button>

        <!-- Minus Sign Button underneath Plus Button -->
        <button class="minus-button">-</button>
      </div>
    </div>
    <SCLogo></SCLogo>
  </div>
</template>

<script>
import NavBar from "../NavBar.vue";
import SCLogo from "../SCLogo.vue";

export default {
  props: ["gid", "gname"],
  name: "GroupChatPage",
  components: {
    NavBar,
    SCLogo,
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
  margin-left: 400px;
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
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
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
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
}

.minus-button:hover {
  background-color: #c0392b;
}
</style>