import {createRouter, createWebHashHistory} from 'vue-router'
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
    /*
    假如我有一个子路由地址为child。
    如果不启用Hash模式，在开发模式下没啥问题，http://localhost:9080/child，
    但是在生产模式下，file://${__dirname}/index.html/child却是无法匹配的一条路径。
    因此在electron下，vue-router请不要使用history模式，而使用默认的hash模式。
    那么上面的问题就迎刃而解，变为file://${__dirname}/index.html#child即可。
     */
    history: createWebHashHistory(process.env.BASE_URL),
    routes
})

export default router
