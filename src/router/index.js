import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import About from '../views/About.vue'
import CV from '../views/CV.vue'
import Projects from '../views/Projects.vue'
import Login from '../views/Login.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/about',
    name: 'About',
    component: About
  },
  {
    path: '/cv',
    name: 'CV',
    component: CV
  },
  {
    path: '/projects',
    name: 'Projects',
    component: Projects
  },
  {
    path: '/login',
    component: Login
  },
  {
    path: '/admin',
    name: 'Admin',
    component: () => import('../views/admin/Admin.vue'),
    children: [
      {path: '/hours', name: 'Hours', component: () => import('../views/admin/Hours.vue')},
      {path: '/dashboard', name: 'Dashboard', component: () => import('../views/admin/Dashboard.vue')},
    ]
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
