<template>
    <div id="home-right">
        <div id="lst-groups">
            <ul v-if="groups.length > 0">
                <li v-for="group in groups" :key="group.groupId" class="home-group">
                    <router-link :to="{name: 'group-chat', params: {gid: group.groupId, gname: group.groupName }}">
                        <div class="contact-group">
                            <button><img src="/src/assets/img/team-ico.png" alt=""/></button>
                            <h7>{{ group.groupName }}</h7>
                                </div>
                    </router-link>
                </li>
            </ul>
            <div v-else class="no-group">
                <p>No groups available. Please create a new group.</p>
            </div>
        </div>  
    </div>
</template>

<script>
import { ref } from 'vue';

export default {
    name: 'HomeContent',
    setup() {
        const groups = ref([]);
        // const groups = ref([
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
