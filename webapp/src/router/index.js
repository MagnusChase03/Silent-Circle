import { createRouter, createWebHistory } from "vue-router";
import HomePage from '../components/Home/HomePage.vue'
import LoginPage from '../components/Login/LoginPage.vue'
import WelcomePage from '../components/Welcome/WelcomePage.vue'
import CreateUserPage from '../components/CreateUser/CreateUserPage.vue'


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
        }
    ]
})

export default router