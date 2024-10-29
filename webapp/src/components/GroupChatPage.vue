<template>
  <div class="page-background">
    <!-- Header with Menu Icon and Group Title -->
    <div class="header">
      <img src="@/assets/img/menu-btn-welcome.png" alt="Menu" class="menu-icon" />
      <h1 class="group-title">CS 4389 Group</h1>
      <img src="@/assets/img/logo-welcome.png" alt="Silent Circle Logo" class="logo" />
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
      <div class="chat-message right">
        <div class="bubble">
          <p class="message">Let's briefly introduce ourselves!</p>
        </div>
      </div>

      <!-- Chat Bubble - Message 4 -->
      <div class="chat-message right">
        <div class="bubble">
          <p class="message">My name is Name! Also a senior!</p>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.page-background {
  display: flex;
  flex-direction: column;
  align-items: center;
  height: 100vh;
  background: radial-gradient(circle at 50% 50%, #279ACD, #808FCD);
  overflow: hidden;
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

.menu-icon {
  width: 30px;
  cursor: pointer;
}

.group-title {
  font-family: cursive;
  font-size: 24px;
  color: white;
  text-align: center;
  flex-grow: 1;
}

.logo {
  width: 50px;
  position: absolute;
  right: 20px;
}

.chat-area {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  width: 100%;
  height: 100%;
  max-width: 1250px;
  max-height: 1000px;
  background-color: white;
  padding: 30px;
  border-radius: 20px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  margin-top: 20px;
  margin-bottom: 20px;
}

.chat-message {
  display: flex;
  align-items: flex-end;
  margin: 10px 0;
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

</style>
