import { createRouter, createWebHistory } from "vue-router";
import HomePage from '../components/Home/HomePage.vue'
import LoginPage from '../components/Login/LoginPage.vue'
import WelcomePage from '../components/Welcome/WelcomePage.vue'
import CreateUserPage from '../components/CreateUser/CreateUserPage.vue'
import NewGroupPage from '../components/NewGroup/NewGroupPage.vue'
import AddMemberPage from '../components/AddMember/AddMemberPage.vue'
import GroupChatPage from '../components/GroupChat/GroupChatPage.vue'
import DeleteUserPage from '../components/DeleteUserPage/DeleteUserPage.vue'
import SymmetricDemo from "../components/SymmetricDemo.vue";
import AsymmetricDemo from "../components/AsymmetricDemo.vue";
import InviteUsersPage from "../components/InviteUsers/InviteUsersPage.vue";
import AcceptInvitePage from "../components/AcceptInvite/AcceptInvitePage.vue";

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            component: WelcomePage
        },
        {
            path: '/login',
            component: LoginPage
        },
        {
            path: '/home',
            component: HomePage
        },
        {
            path: '/create-user',
            component: CreateUserPage
        },
        {
            path: '/new-group',
            component: NewGroupPage
        },
        {
            path: '/add-member',
            component: AddMemberPage
        },
        {
            path: '/invite-users',
            component: InviteUsersPage
        },
        {
            path: '/accept-invite',
            component: AcceptInvitePage
        },
        {
            path: '/group-chat/:gid/:gname',
            name: 'group-chat',
            component: GroupChatPage,
            props: true
        },
        {
            path: '/delete-user',
            name: '/delete-user',
            component: DeleteUserPage,
        }, {
            path: '/symmetric-demo',
            name: 'symmetric-demo',
            component: SymmetricDemo

        },
        {
            path: '/asymmetric-demo',
            name: 'asymmetric-demo',
            component: AsymmetricDemo
        },
        {
            path: '/:pathMatch(.*)*',
            redirect: '/'
        }
    ]
})

export default router