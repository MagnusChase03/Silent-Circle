import { createRouter, createWebHistory } from "vue-router";
import HomePage from '../components/Home/HomePage.vue'
import LoginPage from '../components/Login/LoginPage.vue'
import WelcomePage from '../components/Welcome/WelcomePage.vue'
import CreateUserPage from '../components/CreateUser/CreateUserPage.vue'
import NewGroupPage from '../components/NewGroup/NewGroupPage.vue'
import AddMemberPage from '../components/AddMember/AddMemberPage.vue'

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
        }
    ]
})

export default router