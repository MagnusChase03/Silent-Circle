<template>
  <div id="wrapper-home">
    <!-- NavBar Component -->
    <NavBar></NavBar>
    <div style="width: 100%; height: 100%;">
      <!-- New Group Page Content -->
      <div class="page-container">
        <div class="page-background">
          <div class="header">
            <h1 class="title">Create Group Chat</h1>
          </div>
          <div class="right-side">
            <div class="chat-bubble">
              <div class="group-name-box">
                <input type="text" v-model="groupName" placeholder="Enter Group Name" class="group-name-input"
                  @keyup.enter="createGroup" />
              </div>
              <br />
              <!-- Create group button -->
              <button class="create-group-button" @click="createGroup" :disabled="!groupName.trim()">
                Create Group
              </button>
            </div>
          </div>
          <img src="@/assets/img/create-grp-photo.gif" alt="Animated GIF" class="animated-gif" />
        </div>
      </div>
    </div>
    <!-- End of main content -->
  </div>
</template>

<script>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import NavBar from '../NavBar.vue';
import useSymmetricKey from '../../composables/useSymmetricKey';

export default {
  name: 'NewGroupPage',
  components: {
    NavBar,
  },
  setup() {

    const groupName = ref('');
    const generatedSymmetricKey = ref(''); //symmetric key
    const router = useRouter();
    const createGroup = () => {
      if (!groupName.value.trim()) {
        alert('Please enter a group name');
        return;
      }
      fetch(import.meta.env.VITE_API_URL + "/group/create", {
        method: "POST",
        headers: {
          "Content-Type": "application/x-www-form-urlencoded"
        },
        body: `groupname=${encodeURIComponent(groupName.value)}`,
        credentials: "include"
      })
        .then((res) => {
          if (!res.ok) {
            throw new Error(`Http error! Status: ${res.status}`);
          }
          return res.json();
        })
        .then(async (data) => {
          if (data.StatusCode == 200) {

            console.log(data);

            const groupid = data.Data.GroupID;  // Extracting the GroupID
            const username = localStorage.getItem('username');              
            //generating symmetric key
            generatedSymmetricKey.value = await useSymmetricKey();
            console.log("Symmetric key generated");
            const groupSymmetricKey = `${username}-${groupid}`;

            // Save the symmetric key in localStorage with the key name "username-groupid"
            localStorage.setItem(groupSymmetricKey, generatedSymmetricKey.value);

            alert("Group created successfully!");
            router.push({ path: '/add-member', query: { groupid } });
          }
        })
        .catch((error) => console.error("Unable to fetch data:", error));
    };

    return {
      groupName,
      createGroup,
    };
  }
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

.create-group-button {
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
</style>