<script setup>
    import MyNav from '../NavBar.vue';
</script>

<template>
    <div class="page-container">    
        <MyNav></MyNav>
    <div class="page-background">
      <div class="header">
        <h1 class="title">Add Member to Group</h1>
      </div>
      <div class="right-side">
        <div class="chat-bubble">
          <div class="user-lookup-box">
            <input
              type="text"
              placeholder="Search for a user..."
              class="user-lookup-input"
            />
          </div>
          <br />
          <button class="add-member-button">
            Add Member to Group
          </button>
        </div>
      </div>
  
      <img
        src="@/assets/img/lookup.gif"
        alt="Animated GIF"
        class="animated-gif"
      />
    </div>
    </div>
  </template>
  
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
    /* Gradiant background */
    /*background: radial-gradient(circle at 50% 50%, #279ACD, #808FCD);*/
    background: none;
    overflow: hidden;
    position: relative;    
  }
  .header {
    position: absolute;
    top: 40px;
    left: 40px;
    display: flex;
    align-items: center;
  }
  .title {
    font-family: cursive;
    font-size: 50px;
    color: #808FCD;
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
  .user-lookup-box {
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
  .user-lookup-input {
    background-color: transparent;
    border: none;
    outline: none;
    padding: 10px;
    font-size: 16px;
    color: white;
    border-radius: 25px;
  }
  .add-member-button {
    display: flex;
    align-items: center;
    justify-content: center;
    background: radial-gradient(circle, #279ACD, #808FCD 150px);
    border: none;
    border-radius: 25px;
    width: 100%;
    max-width: 350px;
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
  </style>