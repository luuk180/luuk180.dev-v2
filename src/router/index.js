import { createRouter, createWebHistory } from 'vue-router'
import { Auth } from 'aws-amplify'
import Home from '../views/Home.vue'
import About from '../views/About.vue'
import CV from '../views/CV.vue'
import Projects from '../views/Projects.vue'
import Login from '../views/Login.vue'

function checkAuth(to, from) {
  Auth.currentAuthenticatedUser({
    bypassCache: false  // Optional, By default is false. If set to true, this call will send a request to Cognito to get the latest user data
  }).then(user => console.log(user))
  .catch(err => {
    console.log(err);
    router.push(from)
  })
}

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
    name: 'Login',
    component: Login
  },
  {
    path: '/admin',
    name: 'Admin',
    component: () => import('../views/admin/Dashboard.vue'),
    beforeEnter: [checkAuth],
    children: [
      {path: '/hours', name: 'Hours', component: () => import('../views/admin/Hours.vue')},
    ]
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
