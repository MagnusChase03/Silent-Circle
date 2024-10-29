<template>
    <div class="page-background">
       <!-- X CLOSE BUTTON  have to add <button @click class-->
       <button class="close-button">
            ×
        </button>
        <div class="header">
            <h1 class="title">Create New Chat</h1>
        </div>
        <div class="right-side">
            <div class="chat-bubble">
                <div class="group-name-box">
        <input
          type="text"
          v-model="groupName"
          placeholder="Enter Group Name"
          class="group-name-input"
          @keyup.enter="createGroup"
        />
      </div>
      <br>
      <button class="create-group-button" @click="createGroup" :disabled="!groupName.trim()">
        Create Group
      </button>
            </div>
        </div>
      
      <img
         src="@/assets/img/create-grp-photo.gif" 
         alt="Animated GIF" 
         class="animated-gif" 
    />
 
    </div>
</template>

<script>
export default {
    data() {
        return {
            groupName: ''
        }
    },
    methods: {
        goToHome() {
            this.$router.push('/Home');
        },
        createGroup() {
            if (!this.groupName.trim()) {
                alert('Please enter a group name');
                return;
            }
                  
            // API call here
            // this.axios.post('/api/groups', { name: this.groupName })
            
            // After successful creation, you might want to:
            // 1. Navigate to the new group
            // this.$router.push(`/group/${groupId}`);
            // 2. Or clear the input and show a success message
           // this.groupName = '';
            alert('Group created successfully!');
        }
    }
}
</script>

<style scoped>

.page-background{
  display: flex;
  height: 100vh; 
  background: radial-gradient(circle at 50% 50%, #279ACD, #808FCD);
  overflow: hidden; 
  position: relative; 
}
.close-button {
    position: absolute;
    top: 20px;
    left: 20px;
    width: 40px;
    height: 40px;
    border-radius: 50%;
    background: radial-gradient(circle, #279ACD, #808FCD);
    border: none;
    color: white;
    font-size: 24px;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 0;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
    transition: transform 0.2s ease;
    z-index: 10;
}

.header{
    position: absolute; 
    top: 60px; 
    left: 60px; 
    display: flex; 
    align-items: center; 
}
.title {
  font-family: cursive;
  font-size: 50px; 
  color: white; 
  margin: 0; 
}
.right-side {
  display: flex;
  flex-direction: column; 
  align-items: flex-end; 
  width: 600px;
 
}
.chat-bubble{
    position: absolute; 
    top: 10px; 
    right: 50px; 
    width: 600px;
    height: 500px; 
    background-image: url('@/assets/img/white-chat-bubble.png'); 
    background-size: cover; 
    background-repeat: no-repeat; 
}

.group-name-box {

    background: radial-gradient(circle,#d9fffb,#ccebff); 
    width: 100%;
    max-width: 350px;
    border-radius: 25px; 
    padding: 10px;
    margin-bottom: 10px;
    margin-top: 55px;
    margin-left: 130px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2); 
}
.group-name-input {
  background-color: transparent;
  border: none;
  outline: none;
  padding: 10px;
  font-size: 16px;
  color: #333;
  border-radius: 25px; 
}
.group-name-input::placeholder {
    color: #666; 
}
.create-group-button{
  display: flex;
  align-items: center;
  justify-content: center;
  background: radial-gradient(circle, #279ACD, #808FCD 150px);
  border: none;
  border-radius: 25px;
  width: 100%;
  max-width: 350px;
  margin-left:  130px;
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