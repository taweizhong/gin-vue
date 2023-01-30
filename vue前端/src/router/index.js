import { createRouter, createWebHistory } from 'vue-router'


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/regiester',
      name: 'regiester',
      component: () => import('../views/Regiester/RegiesterUser.vue')
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/Login/LoginUser.vue')
    },
  ]
})

export default router
