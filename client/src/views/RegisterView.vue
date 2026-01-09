<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-100 py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-full space-y-8 bg-white p-10 rounded-xl shadow-lg">
      <div>
        <h2 class="text-center text-3xl font-extrabold text-gray-900">Tạo tài khoản mới</h2>
        <p class="mt-2 text-center text-sm text-gray-600">
          Hoặc
          <router-link to="/login" class="font-medium text-blue-600 hover:text-blue-500">
            đăng nhập nếu đã có tài khoản
          </router-link>
        </p>
      </div>

      <form class="mt-8 space-y-4" @submit.prevent="handleRegister">
        <BaseInput label="Họ và tên" v-model="form.fullname" placeholder="Nguyễn Văn A" required />
        
        <BaseInput label="Email" type="email" v-model="form.email" placeholder="email@example.com" required />
        
        <div class="grid grid-cols-2 gap-4">
          <BaseInput label="Số điện thoại" v-model="form.phone" placeholder="090..." />
          <BaseInput label="Vai trò" v-model="form.role" placeholder="customer" />
        </div>

        <BaseInput label="Địa chỉ" v-model="form.address" placeholder="Hà Nội, Việt Nam" />

        <BaseInput label="Mật khẩu" type="password" v-model="form.password" placeholder="********" required />
        
        <BaseInput label="Xác nhận mật khẩu" type="password" v-model="confirmPassword" placeholder="********" required />

        <div v-if="error" class="text-red-500 text-sm text-center font-medium">
          {{ error }}
        </div>

        <button
          type="submit"
          :disabled="loading"
          class="w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-green-600 hover:bg-green-700 focus:outline-none disabled:bg-green-300"
        >
          {{ loading ? 'Đang xử lý...' : 'Đăng ký ngay' }}
        </button>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { authApi } from '../api/auth'
import BaseInput from '../components/common/BaseInput.vue'

const router = useRouter()
const loading = ref(false)
const error = ref('')
const confirmPassword = ref('')

const form = reactive({
  fullname: '',
  email: '',
  password: '',
  phone: '',
  address: '',
  role: 'customer'
})

const handleRegister = async () => {
  if (form.password !== confirmPassword.value) {
    error.value = 'Mật khẩu xác nhận không khớp!'
    return
  }

  loading.value = true
  error.value = ''

  try {
    await authApi.register(form)
    
    alert('Đăng ký thành công! Hãy đăng nhập.')
    router.push('/login')
  } catch (err: any) {
    error.value = err.response?.data?.error || 'Đăng ký thất bại. Vui lòng thử lại.'
  } finally {
    loading.value = false
  }
}
</script>