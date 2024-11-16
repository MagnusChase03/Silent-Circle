<template>
    <div id="wrapper-home">
        <NavBar></NavBar>
        <div id="home-right">
            <!-- Insert you main content here -->
            <div class="add-member">
                <h1>Accept Ongoing Group Invitations</h1><br/>

            </div>
            <div id="lst-groups">
                <ul>
                    <li v-for="group in groups" :key="group.groupId" class="home-group">
                            <router-link :to="{name: 'group-chat', params: {gid: group.groupId, gname: group.groupName }}">
                                <div class="contact-group">
                                        <button><img src="/src/assets/img/team-ico.png" alt=""/></button>
                                        <h7>{{ group.groupName }}</h7>
                                </div>
                            </router-link>
                            <a><button class="send-group-invite">Send Group Invitation</button></a>
                    </li>
                </ul>
            </div>
        </div>
        <SCLogo></SCLogo>
    </div>
</template>

<script>
    import { ref } from 'vue';
    import NavBar from '../NavBar.vue';
    import SCLogo from '../SCLogo.vue';

    export default {
        name: 'AcceptInvitePage',
        components: {
            NavBar,
            SCLogo
        },
        setup(){
            const groups = ref([
                {
                    "groupId": "101",
                    "groupName": "CS4389"
                },
                {
                    "groupId": "102",
                    "groupName": "Work"
                },
                {
                    "groupId": "103",
                    "groupName": "Family"
                },
                {
                    "groupId": "104",
                    "groupName": "Sports"
                },
                {
                    "groupId": "105",
                    "groupName": "Book Club"
                },
                {
                    "groupId": "106",
                    "groupName": "CS4389"
                },
                {
                    "groupId": "107",
                    "groupName": "Work"
                },
                {
                    "groupId": "108",
                    "groupName": "Family"
                },
                {
                    "groupId": "109",
                    "groupName": "Sports"
                },
                {
                    "groupId": "1010",
                    "groupName": "Book Club"
                }
            ]);

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
</style>