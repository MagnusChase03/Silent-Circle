<template>
    <div id="wrapper-home">
        <NavBar></NavBar>
        <div id="accept-invite-wrapper">
            <!-- Insert you main content here -->
            <div class="add-member">
                <h1>Accept Ongoing Group Invitations</h1><br />
            </div>

            <div id="lst-invites">
                <ul v-if="invites.length > 0" class="accept-invite">
                    <li v-for="invite in invites" :key="invite.GroupID" class="home-group">
                        <router-link
                            :to="{name: 'group-chat', params: {gid: invite.GroupID, gname: invite.GroupName }}">
                            <div class="contact-group">
                                <button><img src="/src/assets/img/invite-ico.png" alt="" /></button>
                                <h7>{{ invite.GroupName }}</h7>
                            </div>
                        </router-link>
                        <button class="accept-group-invite" @click="acceptInvite(invite.GroupID,invite.GroupName)">Accept</button>
                        <button class="reject-group-invite" @click="rejectInvite(invite.GroupID,invite.GroupName)">Reject</button>
                    </li>
                </ul>
                <div v-else class="no-group">
                    <p style="color:#060658">You currently have no group invitation.</p>
                </div>
            </div>
        </div>
        <SCLogo></SCLogo>
    </div>
</template>

<script>
    import { onMounted, ref } from 'vue';
    import NavBar from '../NavBar.vue';
    import SCLogo from '../SCLogo.vue';
    import useDecryptAsymMsg from '../../composables/useDecryptAsymMsg.js';

    export default {
        name: 'AcceptInvitePage',
        components: {
            NavBar,
            SCLogo
        },
        setup() {
            // data
            const invites = ref([]);
            const privateKey = ref(localStorage.getItem('privateKey')); 
            const encryptedSymmetricKey = ref('');
            const decryptedsymmetricKey = ref('');
            const username = localStorage.getItem('username'); 
            const groupSymmetricKey = ref('');

            // methods
            const acceptInvite = async (gid, gname)=>{
                await fetch(import.meta.env.VITE_API_URL + "/group/invite/accept", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/x-www-form-urlencoded"
                    },
                    body: `group=${gid}`,
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
                            // Retrieve the Group Encrypted Symmetric Key
                            encryptedSymmetricKey.value = data.Data.EncryptedKey;  

                            //decrypt group encrypted symmetric key
                            decryptedsymmetricKey.value = await useDecryptAsymMsg(privateKey.value, encryptedSymmetricKey.value).catch((error) => {
                                console.error("Unable to decrypt symmetric key data:", error);
                                return;
                            });

                            // name the group symmetric key
                            groupSymmetricKey.value = `${username}-${gid}`;

                            // Save the group symmetric key in localStorage with the key name "username-groupid"
                            localStorage.setItem(groupSymmetricKey.value, decryptedsymmetricKey.value);

                            alert(`Accepted invitation to group ${gname} successfully!`);

                            load();
                        }
                    })
                    .catch((error) => console.error("Unable to fetch data:", error));
            };
            
            const rejectInvite = async (gid,gname)=>{
                await fetch(import.meta.env.VITE_API_URL + "/group/invite/reject", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/x-www-form-urlencoded"
                    },
                    body: `group=${gid}`,
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
                            alert(`Recejed invitation to group ${gname}!`);
                        }
                    })
                    .catch((error) => console.error("Unable to fetch data:", error));
            }

            const load = async () => {
                // reset invites
                invites.value = [];
                
                // get ongoing invitation to the user
                await fetch(import.meta.env.VITE_API_URL + "/group/invite/get", {
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
                        // get invites from the response data
                        invites.value = data.Data.Groups;
                    }
                    // If the response is not ok, throw an error
                }).catch((error) => console.error("Unable to tetch data:", error)
                ).catch((error) => console.error("Unable to tetch data:", error));
            }

            onMounted(() => {
                load();
            });

            return { invites, acceptInvite, rejectInvite }
        },

    }
</script>

<style scoped>
#accept-invite-wrapper{
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
}
.add-member h1{
    margin-bottom: 20px;
    margin-left: 20px;
    font-family: "Cookie", cursive;
    color: #516dc1;
    font-size: 2.5em;
}

    #accept-invite-wrapper div#lst-groups {
        overflow-y: auto;
        border: 1px solid #ccc;
        width: 100%;
        display: block;
        padding: 0;
        margin: 0;
    }


        #accept-invite-wrapper ul {
            display: block;
            list-style: none;
            height: 100%
        }
    
        #accept-invite-wrapper ul li {
            display: block;
            border-bottom: #e7e7e7 3px solid;
            height: 25%;
            min-height: 150px;
            display: flex;
            flex-direction: row;
            align-items: center;
        }
    
        #accept-invite-wrapper ul li:last-child {
            border: none;
        }
    
        #accept-invite-wrapper .contact-group {
            color: #5c6689;
            font-size: 25px;
            display: flex;
            flex-direction: row;
            justify-content: flex-start;
            align-items: center;
            margin-left: 20px;
        }
    
        #accept-invite-wrapper button {
            display: inline-block;
            background: none;
            border: none;
        }
    
        #accept-invite-wrapper button img {
            width: 80%;
        }


    
    .home-group{
        justify-content: left;
    }


    .accept-group-invite {
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

    .accept-group-invite:hover {
        background-color: white;
        color: black;
        border: 2px solid #4CAF50;
    }

    .reject-group-invite {
        background-color: #f44336 !important;
        border: none;
        color: white;
        padding: 10px 24px;
        text-align: center;
        text-decoration: none;
        display: inline-block;
        font-size: 16px;
        margin: 4px 2px;
        transition-duration: 0.4s;
        cursor: pointer;
        border-radius: 12px;
    }

    .reject-group-invite:hover {
        background-color: white;
        color: black;
        border: 2px solid #f44336;
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

    div.no-group p {
        font-size: 1.8em;
        color: #CCC;
    }

    #lst-invites {
        margin-left: 20px;
    }

        #lst-invites li:nth-child(odd) {
            background-color: #f2f2f2;
        }
    
    
    
        #lst-invites li:nth-child(even) {
            background-color: #ffffff;
        }

</style>