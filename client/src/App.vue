<script setup lang="ts">
import AppHeader from './components/common/AppHeader.vue'
import AppFooter from './components/common/AppFooter.vue'
import { onMounted } from 'vue'
import { useAuthStore } from './stores/user'
import { userApi } from './api/user'
import { jwtDecode } from 'jwt-decode'

const authStore = useAuthStore()

onMounted(async () => {
  if (authStore.accessToken && !authStore.user) {
    try {
      const decoded: any = jwtDecode(authStore.accessToken)
      const userId = decoded.user_id
      
      const res = await userApi.getUserById(userId)
      authStore.setUser(res.data.data) 
    } catch (error) {
      console.error("Lỗi tự động lấy thông tin người dùng:", error)
      authStore.logout() 
    }
  }
})
</script>

<template>
  <div class="app-container min-h-screen flex flex-col">
    <AppHeader />
    <main class="flex-grow">
      <router-view />
    </main>
    <AppFooter />
  </div>
</template>
