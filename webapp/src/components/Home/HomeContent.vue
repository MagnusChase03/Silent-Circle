<template>
    <div id="home-right">
        <div id="lst-groups">
            <ul v-if="groups.length > 0">
                <li v-for="group in groups" :key="group.GroupID" class="home-group">
                    <router-link :to="{name: 'group-chat', params: {gid: group.GroupID, gname: group.GroupName }}">
                        <div class="contact-group">
                            <button><img src="/src/assets/img/team-ico.png" alt=""/></button>
                            <h7>{{ group.GroupName }}</h7>
                                </div>
                    </router-link>
                </li>
            </ul>
            <div v-else class="no-group">
                <p style="color:#060658">No groups available. <router-link to="/new-group">Click here</router-link> to create a new group.</p>
            </div>
        </div>  
    </div>
</template>

<script>
import { onMounted, ref } from 'vue';

export default {
    name: 'HomeContent',
    setup() {
        // data
        const groups = ref([]);
        const username = ref(localStorage.getItem('username'));

        
        const load = async () => {
            console.log("Loading groups for user: ", username.value);
            // get groups that an user belongs to
            fetch(import.meta.env.VITE_API_URL + "/user/groups", {
                // Send the username and password to the server
                method: "POST",
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded"
                },
                credentials: "include"
            }).then((res) => {
                // Check if the response is ok
                if (!res.ok) {
                    // If the response is not ok, throw an error
                    throw new Error(`Http error! Status: ${res.status}`);
                }

                // Return the response as JSON
                return res.json();

            }).then((data) => {
                alert("Groups loaded successfully", data);
                if (data.StatusCode == 200) {
                    // get the public key from the response data
                    groups.value = data.Data.Groups;
                    console.log(data);
                }
                // If the response is not ok, throw an error
            }).catch((error) => console.error("Unable to tetch data:", error)
            ).catch((error) => console.error("Unable to tetch data:", error));

        }
        console.log("calling load");
        load();
        //     {
        //         "groupId": "101",
        //         "groupName": "CS4389"
        //     },
        //     {
        //         "groupId": "102",
        //         "groupName": "Work"
        //     },
        //     {
        //         "groupId": "103",
        //         "groupName": "Family"
        //     },
        //     {
        //         "groupId": "104",
        //         "groupName": "Sports"
        //     },
        //     {
        //         "groupId": "105",
        //         "groupName": "Book Club"
        //     },
        //     {
        //         "groupId": "106",
        //         "groupName": "CS4389"
        //     },
        //     {
        //         "groupId": "107",
        //         "groupName": "Work"
        //     },
        //     {
        //         "groupId": "108",
        //         "groupName": "Family"
        //     },
        //     {
        //         "groupId": "109",
        //         "groupName": "Sports"
        //     },
        //     {
        //         "groupId": "1010",
        //         "groupName": "Book Club"
        //     }
        // ]);

        return { groups }
    }
}
</script>

<style scoped>
    #home-right {
        width: 100%;
        height: 100%;
    }
    #home-right div#lst-groups {
        height: 100%;
        overflow-y: auto;
        border: 1px solid #ccc;
        width: 100%;
        display: block;
        padding: 0;
        margin: 0;
    }

    .home-group a{
        text-decoration: none;
    }

    .home-group a:hover{
        cursor: pointer;
    }

    .home-group button:hover{
        cursor: pointer;
    }

    div.no-group {
        width: 100%;
        height: 100%;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-content: space-around;
        align-items: center;
    }

    div.no-group p {
        font-size: 1.8em;
        color: #CCC;
    }
</style>
