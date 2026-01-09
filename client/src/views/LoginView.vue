<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-100 py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-full space-y-8 bg-white p-10 rounded-xl shadow-lg">
      <h2 class="text-center text-3xl font-extrabold text-gray-900">Đăng nhập vào BookStore</h2>
      
      <form class="mt-8 space-y-6" @submit.prevent="handleLogin">
        <BaseInput 
          label="Email" 
          type="email" 
          v-model="loginForm.email" 
          placeholder="email@example.com" 
        />
        
        <BaseInput 
          label="Mật khẩu" 
          type="password" 
          v-model="loginForm.password" 
          placeholder="********" 
        />

        <div v-if="error" class="text-red-500 text-sm text-center">
          {{ error }}
        </div>

        <button
          type="submit"
          :disabled="loading"
          class="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none disabled:bg-blue-300"
        >
          {{ loading ? 'Đang xử lý...' : 'Đăng nhập' }}
        </button>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { authApi } from '../api/auth'
import { useAuthStore } from '../stores/user'
import BaseInput from '../components/common/BaseInput.vue'

const router = useRouter()
const authStore = useAuthStore()

const loginForm = reactive({
  email: '',
  password: ''
})

const loading = ref(false)
const error = ref('')

const handleLogin = async () => {
  loading.value = true
  error.value = ''
  
  try {
    const res = await authApi.login(loginForm)
    console.log('Dữ liệu trả về từ API:', res)
    const { access_token, refresh_token, user } = res.data
    
    authStore.setTokens(access_token, refresh_token)
    if (user) authStore.setUser(user)

    console.log('Đăng nhập thành công:', user)
    
    router.push('/')
  } catch (err: any) {
    error.value = err.response?.data?.error || 'Đăng nhập thất bại. Vui lòng thử lại.'
  } finally {
    loading.value = false
  }
}
</script>