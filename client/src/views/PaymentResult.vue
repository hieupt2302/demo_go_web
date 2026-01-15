<script setup lang="ts">
import { onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useCartStore } from '../stores/cart'
import { orderApi } from '../api/order'

const route = useRoute()
const router = useRouter()
const cartStore = useCartStore()

onMounted(async () => {
  const vnpParams = route.query
  console.log("Dữ liệu từ VNPAY:", vnpParams)

  if (!vnpParams.vnp_ResponseCode) {
    router.push('/')
    return
  }

  try {
    // 1. Gọi Backend và truyền toàn bộ vnpParams qua query string
    const response = await orderApi.vnpay(vnpParams)
    
    if (response.data) {
      cartStore.clearCart()

      // 2. Lưu trạng thái vào sessionStorage
      sessionStorage.setItem('payment_status', 'success')
      sessionStorage.setItem('payment_order_id', vnpParams.vnp_TxnRef as string)
      
      // 3. Về Home ngay lập tức
      router.push('/')
    }
  } catch (err: any) {
    console.error("Lỗi xác thực:", err)
    sessionStorage.setItem('payment_status', 'failed')
    router.push('/')
  }
})
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-white text-gray-400">
  </div>
</template>