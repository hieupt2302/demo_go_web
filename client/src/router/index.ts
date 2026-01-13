import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '../stores/user'
import LoginView from '../views/LoginView.vue'
import RegisterView from '../views/RegisterView.vue'
import HomeView from '../views/HomeView.vue'
import CartView from '../views/CartView.vue'
import OrdersView from '../views/OrdersView.vue'
import AdminView from '../views/AdminView.vue'
import BookDetail from '../components/product/BookDetail.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: LoginView
    },
    {
      path: '/register',
      name: 'register',
      component: RegisterView
    },
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/cart',
      name: 'cart',
      component: CartView
    },
    {
      path: '/orders',
      name: 'orders',
      component: OrdersView
    },
    {
      path: '/admin',
      name: 'admin',
      component: AdminView,
      meta: { requiresAdmin: true }
    },
    {
      path: '/book/:id',
      name: 'book',
      component: BookDetail
    }
  ]
})

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()

  const publicPages = ['login', 'register', 'home']
  const authRequired = !publicPages.includes(to.name as string)

  if (authRequired && !authStore.accessToken) {
    return next({ name: 'login' })
  }

  // Nếu đã đăng nhập mà cố tình vào lại trang login/register
  if (authStore.isAuthenticated && (to.name === 'login' || to.name === 'register')) {
    return next({ name: 'home' })
  }

  // Kiểm tra quyền truy cập
  if (to.meta?.requiresAdmin && !authStore.isAdmin) {
    return next({ name: 'home' })
  }

  next()
})

export default router