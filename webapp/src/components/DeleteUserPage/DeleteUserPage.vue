<template>
  <div id="wrapper-home">
    <NavBar></NavBar>
<div class="page-background">
  <!-- Header -->
  <header class="header">
    <!-- Title and search -->
    <h1 class="group-title">CS 4389 Group</h1>
    <div class="search-container">
      <input 
        type="search" 
        placeholder="Search Members in group" 
        v-model="searchQuery"
        class="search-input"
        @input="handleSearch"
      />
      <!-- Search icon image -->
      <img src="@/assets/img/search-ico.png" alt="Search" class="search-icon">
    </div>
  </header>

  <!-- Main Content Layout -->
  <div class="content-layout">
    <!-- Members List -->
    <div class="members-container">
      <h2 class="members-count">
        {{ filteredMembers.length }} member{{ filteredMembers.length !== 1 ? 's' : '' }}
      </h2>
      <div class="members-list">
        <!-- No Results Message when searched-->
        <div v-if="filteredMembers.length === 0" class="no-results">
          No members found matching "{{ searchQuery }}"
        </div>
        
        <div 
          v-for="member in displayedMembers" 
          :key="member.id"
          class="member-item"
          :class="{ 'highlight': searchQuery && member.name.toLowerCase().includes(searchQuery.toLowerCase()) }"
        >
          <div class="member-info">
            <div class="member-avatar" :style="{ backgroundColor: member.avatarColor }">
              <img v-if="member.avatar" :src="member.avatar" :alt="member.name" class="avatar-image" />
              <img v-else :src="icon" alt="Icon" class="avatar-icon" />
            </div>
            <span class="member-name">{{ member.name }}</span>
          </div>
          <button 
            @click="deleteMember(member)" 
            class="delete-button"
            :title="'Delete ' + member.name"
          >
            ×
          </button>
        </div>
        <!-- Show More Button -->
        <button 
          v-if="hasMoreMembers" 
          @click="showMore" 
          class="show-more-button"
        >
          Show More...
        </button>
      </div>
    </div>

    <!-- Delete User Button -->
    <div class="delete-user-container">
      <div class="delete-user-content">
        <button 
          @click="deleteSelectedUser" 
          class="delete-user-button"
          :disabled="!selectedMember"
        >
          Delete User
        </button>
        <!-- Group Icon image -->
        <img src="@/assets/img/groups-ico.png" alt="groups icon" class="groups-icon-image" />
      </div>
    </div>
  </div>
</div>
</div>
</template>

<script >
import icon from '@/assets/img/user-blue-ico.png';
import NavBar from "../NavBar.vue";
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';

export default {
  name: 'Group',
  components: {
    NavBar
  },
  setup() {
    const router = useRouter();
    const selectedMember = ref(null); // Store selected member

    // Initialize data
    const members = ref([
      { id: 1, name: 'User 1', avatar: icon },
      { id: 2, name: 'User 2', avatar: icon },
      { id: 3, name: 'User 3', avatar: icon },
      { id: 4, name: 'User 4', avatar: icon },
      { id: 5, name: 'User 5', avatar: icon },
      { id: 6, name: 'User 6', avatar: icon },
      { id: 7, name: 'User 7', avatar: icon },
      { id: 8, name: 'User 8', avatar: icon },
    ]);

    const searchQuery = ref('');
    const displayLimit = ref(3);
    const searchTimeout = ref(null);

    const filteredMembers = computed(() => {
      if (!searchQuery.value) {
        return members.value;
      }
      const query = searchQuery.value.toLowerCase().trim();
      return members.value.filter(member =>
        member.name.toLowerCase().includes(query)
      );
    });

    const displayedMembers = computed(() => {
      return filteredMembers.value.slice(0, displayLimit.value);
    });

    const hasMoreMembers = computed(() => {
      return filteredMembers.value.length > displayLimit.value;
    });

    const handleSearch = (event) => {
      if (searchTimeout.value) {
        clearTimeout(searchTimeout.value);
      }
      searchTimeout.value = setTimeout(() => {
        searchQuery.value = event.target.value;
        displayLimit.value = 3;
      }, 300);
    };

    const showMore = () => {
      displayLimit.value += 3;
    };

    const deleteMember = (member) => {
      selectedMember.value = member;
    };

    const deleteSelectedUser = async () => {
      if (!selectedMember.value) {
        alert('Please select a member to delete');
        return;
      };

// will adjust fetch method accordingly after other pages are done
      try {
        const response = await fetch(import.meta.env.VITE_API_URL + "/group/ban", {
          method: "POST",
          headers: {
            "Content-Type": "application/x-www-form-urlencoded"
          },
          body: JSON.stringify({
            group: "groupId", 
            username: selectedMember.value.name
          }),
          credentials: "include"
        });

        if (response.ok) {
          const data = await response.json();
          if (data.StatusCode === 200) {
            members.value = members.value.filter(member => member.id !== selectedMember.value.id);
            selectedMember.value = null;
            alert("User deleted successfully");
            router.push('/home');
          } else {
            alert(data.Data || "Failed to delete user");
            router.push('/home');
          }
        } else {
          alert(`Error: ${response.status}`);
          router.push('/home');
        }
      } catch (error) {
        console.error("Unable to delete user:", error);
        alert("Error occurred while trying to delete user.");
      }
    };

    return {
      members,
      searchQuery,
      displayLimit,
      searchTimeout,
      selectedMember,
      filteredMembers,
      displayedMembers,
      hasMoreMembers,
      handleSearch,
      showMore,
      deleteMember,
      deleteSelectedUser,
    };
  }
};
</script>

<style scoped>
.page-background {
font-family: cursive;
max-width: auto;
margin: 0 auto;
width: 100%
}

.header {
background: linear-gradient(120deg, #808FCD 10%, #279ACD 100%);
padding: 1rem;
display: flex;
align-items: center;
gap: 1rem;
border-radius: 0px 0px 0 0;
}

.group-title {
color: white;
margin: 0;
flex-grow: 1;
font-size: 1.2rem;
}

.search-container {
position: relative;
flex-grow: 2;
}

.search-input {
width: 100%;
padding: 0.5rem 1rem 0.5rem 2.5rem;
border: none;
border-radius: 20px;
background: rgba(255, 255, 255, 0.9);
}

.search-icon {
position: absolute;
left: 10px;
top: 50%;
transform: translateY(-50%);
width: 20px;
height: 20px;
pointer-events: none;
}

.search-input {
width: 100%;
padding: 0.5rem 1rem 0.5rem 2.5rem;
border: none;
border-radius: 20px;
background: rgba(255, 255, 255, 0.9);
transition: all 0.3s ease;
}

.search-input:focus {
outline: none;
box-shadow: 0 0 0 2px rgba(255, 255, 255, 0.5);
background: white;
}

.content-layout {
display: flex;
padding: 1rem;
gap: 2rem;
}

.members-container {
flex: 1;
background: white;
padding: 1rem;
border-radius: 10px;
}

.members-count {
margin: 0 0 1rem 0;
font-size: 1.1rem;
color: black;
}

.members-list {
display: flex;
flex-direction: column;
gap: 0.5rem;
}

.member-item {
display: flex;
align-items: center;
justify-content: space-between;
padding: 0.5rem;
background: #bee3ed;
border-radius: 5px;
transition: background-color 0.3s ease;
}

.member-item.highlight {
background-color: #a8d8e6;
border: 1px solid #27a6dc;
}

.no-results {
text-align: center;
padding: 1rem;
color: black;
font-style: italic;
background: rgb(214, 212, 212);
border-radius: 5px;
margin: 1rem 0;
}

.member-info {
display: flex;
align-items: center;
gap: 0.5rem;
}

.member-avatar {
width: 50px;
height: 50px;
border-radius: 50%;
position: relative;
}

.avatar-image {
width: 100%;
height: 100%;
border-radius: 50%;
object-fit: cover;
}

.avatar-icon {
width: 20px;
height: 20px;
position: absolute;
top: 50%;
left: 50%;
transform: translate(-50%, -50%);
}

.delete-button {
background: none;
border: none;
font-size: 1.5rem;
cursor: pointer;
color: black;
}

.show-more-button {
background: none;
border: none;
color: black;
cursor: pointer;
padding: 0.5rem;
font-size: 0.9rem;
text-align: left;
}

.delete-user-container {
width: 300px;
display: flex;
align-items: center;
justify-content: center;
}

.delete-user-content {
display: flex;
flex-direction: column;
align-items: center;
gap: 1.5rem;
margin-top: 2rem;
}

.delete-user-button {
background: linear-gradient(120deg, #808FCD 10%, #279ACD 100%);
color: white;
border: none;
padding: 1rem 2rem;
border-radius: 100px;
cursor: pointer;
font-size: 1rem;
width: 200px;
}

.delete-user-button:disabled {
opacity: 0.7;
cursor: not-allowed;
}

.groups-icon-image {
width: 200px;
height: auto;
}
</style>
