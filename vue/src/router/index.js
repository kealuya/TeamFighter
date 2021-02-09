import {createRouter, createWebHistory} from 'vue-router'
import Test from '../views/Test.vue'
import Home2 from '../views/Home2.vue'
import About from '../views/About.vue'

const routes = [
    {
        path: '/',
        name: 'Test',
        component: Test
    },
    {
        path: '/home2',
        name: 'Home2',
        component: Home2
    },
    {
        path: '/about',
        name: 'About',
        // 延迟加载页面在electron中不能使用，待研究
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        // component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
        component: About
    }
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

export default router
