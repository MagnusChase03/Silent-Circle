<!--<script setup>
    import MyNav from '../NavBar.vue';
</script> -->
<template>
   <div class="wrapper-home">
     <!-- Nav Bar -->
    <NavBar></NavBar>
    <!-- Insert your content or any component here -->

    <div class="page-background">    
      <div class="header">
          <h1 class="title">Create Group Chat</h1>
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
  </div>
      <br>
    <!-- Create group button -->
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
  <div style="width: 100%; height: 100%;">
    </div>
    <SCLogo></SCLogo>
   </div>
</template>

<script>
    import NavBar from '../NavBar.vue';
    import SCLogo from '../SCLogo.vue';

export default {
  name: 'NewGroupPage',
        components: {
            NavBar,
            SCLogo
        },
    data() {
        return {
            groupName: ''
        }
    },
    methods: {
        goToHome() {
            this.$router.push('/home');
        },
        createGroup() {
            if (!this.groupName.trim()) {
                alert('Please enter a group name');
                return;
            }
                // Navigate to AddMemberPage
                this.$router.push('/add-member');
                //alert('Group created successfully!');
        }
    }
}
</script>
<style>
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
  color: #808FCD;
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
    background-image: url('@/assets/img/bubble-chat-create-group.png'); 
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
/* <template>
    <div id="wrapper-home">
        <NavBar></NavBar>
        <!-- Insert your content or any component here -->
        <div style="width: 100%; height: 100%;">
            <h1>This Is New Group Page Again after fix</h1>
        </div>
        <!-- End of main content -->

        <SCLogo></SCLogo>
    </div>
</template>
<script>
    import NavBar from '../NavBar.vue';
    import SCLogo from '../SCLogo.vue';

    export default {
        name: 'NewGroupPage',
        components: {
            NavBar,
            SCLogo
        }
    }
</script>
<style>
</style>*/