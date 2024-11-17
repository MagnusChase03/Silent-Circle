<template>
    <div id="wrapper-home">
        <NavBar></NavBar>
        <div id="home-right">
            <!-- Insert you main content here -->
            <div class="add-member">
                <h1>Invite Users Page</h1><br/>
                <label for="txt-username">Username:</label>
                <input type="text" id="txt-username" name="txt-username" v-model="invitee_uname" required>
                <!-- <button @click="addMember">Add Member</button> -->
            </div>
            <div id="lst-groups">
                <ul>
                    <li v-for="group in groups" :key="group.GroupID" class="home-group">
                            <router-link :to="{name: 'group-chat', params: {gid: group.GroupID, gname: group.GroupName }}">
                                <div class="contact-group">
                                        <button><img src="/src/assets/img/team-ico.png" alt=""/></button>
                                        <h7>{{ group.GroupName }}</h7>
                                </div>
                            </router-link>
                            <a><button class="send-group-invite" @click="sendInvite(group.GroupID)">Send Group Invitation</button></a>
                    </li>
                </ul>
            </div>
        </div>
        <SCLogo></SCLogo>
    </div>
</template>

<script>
    import { onMounted, ref } from 'vue';
    import NavBar from '../NavBar.vue';
    import SCLogo from '../SCLogo.vue';
    import useEncryptAsymMsg from '../../composables/useEncryptAsymMsg.js';

    export default {
        name: 'AddMemberPage',
        components: {
            NavBar,
            SCLogo
        },
        setup(){
            // data
            const groups = ref([]);
            const invitee_uname = ref('');
            const publicKey = ref('');

            const sendInvite = async (gid) => {
                // get the public key of the user
                if (!invitee_uname.value) {
                    alert("Enter invitee username.");
                    return;
                }

                // get the public key of the invitee user
                await getInviteePublicKey().catch((error) => {
                    console.error("Unable to fetch public key:", error);
                    return;
                });


                // retrieve the group symmetric key
                const groupSymmetricKey = getGroupSymmetricKey(gid);
                const encryptedGroupSymmetricKey = await useEncryptAsymMsg(publicKey.value,groupSymmetricKey).catch((error) => {
                    console.error("Unable to encrypt data:", error);
                    return;
                });

                console.log("Encrypted Group Symmetric Key:", encryptedGroupSymmetricKey);

                // send the group invitation through the API
                await fetch(import.meta.env.VITE_API_URL + "/group/invite", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/x-www-form-urlencoded"
                    },
                    body: `username=${invitee_uname.value}&key=${encodeURIComponent(encryptedGroupSymmetricKey)}&group=${gid}`,
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
                            // Retrieve the Group Encrypted Symmetric Key
                            alert("Group invitation sent successfully.");
                        }
                    })
                    .catch((error) => console.error("Unable to fetch data:", error));
            }

            // methods
            function getGroupSymmetricKey(groupid) {
                const invitor_uname = localStorage.getItem('username');
                const groupSymmetricKey = localStorage.getItem(`${invitor_uname}-${groupid}`);
                if (groupSymmetricKey) {
                    console.log("Group symmetric key retrieved:", groupSymmetricKey); // Debug statement
                    return groupSymmetricKey;
                } else {
                    console.error("Group symmetric key not found.");
                }
            }

            async function getInviteePublicKey() {
                try {
                    const response = await fetch(import.meta.env.VITE_API_URL + "/user/get", {
                        method: "POST",
                        headers: {
                            "Content-Type": "application/x-www-form-urlencoded",
                        },
                        body: `username=${invitee_uname.value}`,
                        credentials: "include",
                    });

                    if (!response.ok) {
                        throw new Error(`Http error! Status: ${response.status}`);
                    }

                    const data = await response.json();

                    if (data.StatusCode === 200) {
                        publicKey.value = data.Data.PublicKey;
                        console.log("Public Key Retrieved:", publicKey.value); // Debug statement
                    } else {
                        throw new Error('Failed to fetch public key');
                    }
                } catch (error) {
                    console.error("Unable to fetch public key:", error);
                    alert("Unable to fetch public key. Please try again.");
                }
            }

            const load = async () => {
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
                    if (data.StatusCode == 200) {
                        // get the public key from the response data
                        groups.value = data.Data.Groups;
                        console.log(data);
                    }
                    // If the response is not ok, throw an error
                }).catch((error) => console.error("Unable to tetch data:", error)
                ).catch((error) => console.error("Unable to tetch data:", error));
            }

            // lifecycle hooks
            onMounted(() => {
                load();
            });


            return { groups, sendInvite, invitee_uname }
        }
    }
</script>

<style scoped>
   #home-right {
        width: 100%;
        height: 100%;
    }
    #home-right div#lst-groups {
        height: 85%;
        overflow-y: auto;
        border: 1px solid #ccc;
        width: 100%;
        display: block;
        padding: 0;
        margin: 0;
    }
    
    .home-group{
        justify-content: space-between;
    }

    .home-group .send-group-invite{
        margin-right: 100px;
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

    #home-right .add-member {
        width: 100%;
        height: 15%;
        display: block;
        padding: 0;
        margin: 0;
    }

    .send-group-invite {
        background-color: #4CAF50 !important;
        border: none;
        color: white;
        padding: 10px 24px;
        text-align: center;
        text-decoration: none;
        display: inline-block;
        font-size: 16px;
        margin: 4px 20px;
        transition-duration: 0.4s;
        cursor: pointer;
        border-radius: 12px;
    }

    .send-group-invite:hover {
        background-color: white;
        color: black;
        border: 2px solid #4CAF50;
    }

</style>
