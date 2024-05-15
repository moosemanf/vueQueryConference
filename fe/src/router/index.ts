import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import MasterView from '../views/MasterView.vue'
import DetailView from '@/views/DetailView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/master',
      name: 'm',
      redirect: '/master/view'
    },
    {
      path: '/master/view',
      component: MasterView,
      name: 'Master',
      props: true,
      children: [
        {
          // UserProfile will be rendered inside User's <router-view>
          // when /user/:id/profile is matched
          path: ':id',
          component: DetailView,
          name: 'Detail',
          props: true,
        },
      ],
    },
    {
      path: '/dropdown',
      name: 'dropdown',
      component: () => import('../views/DropdownView.vue')
    },
    {
      path: '/sse',
      name: 'sse',
      component: () => import('../views/SseView.vue')
    }
  ]
})

export default router
