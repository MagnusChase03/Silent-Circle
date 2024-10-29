import { createRouter, createWebHistory } from "vue-router";
import HomePage from '../components/MyHome.vue'
import LoginPage from '../components/LoginUser.vue'
import WelcomePage from '../components/MyWelcome.vue'


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
        }
    ]
})

export default router