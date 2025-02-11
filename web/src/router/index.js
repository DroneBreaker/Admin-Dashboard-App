import { createRouter, createWebHistory } from 'vue-router'

import Login from '../views/Login.vue'
import Home from '../views/Home.vue'

const routes = [
  { path: '/', component: Login },
  { path: '/dashboard', component: Home },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router