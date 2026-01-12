<template>
  <div class="bg-gray-50 min-h-screen py-12">
    <div class="max-w-7xl mx-auto px-4">
      <div class="mb-8">
        <h1 class="text-4xl font-extrabold text-gray-900">Đơn hàng của tôi</h1>
        <p class="text-gray-500 mt-2">Quản lý và theo dõi đơn hàng của bạn</p>
      </div>

      <div v-if="loading" class="flex justify-center items-center py-20">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
      </div>

      <div v-else-if="!orders || orders.length === 0" class="bg-white rounded-2xl shadow-sm border border-gray-200 p-12 text-center">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-24 w-24 text-gray-300 mx-auto mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
        </svg>
        <h2 class="text-2xl font-bold text-gray-900 mb-2">Chưa có đơn hàng</h2>
        <p class="text-gray-500 mb-6">Bạn chưa thực hiện giao dịch nào.</p>
        <router-link to="/" class="inline-block bg-blue-600 hover:bg-blue-700 text-white px-8 py-3 rounded-lg font-semibold transition-colors">
          Bắt đầu mua sắm
        </router-link>
      </div>

      <div v-else class="space-y-6">
        <div v-for="order in orders" :key="order.id" class="bg-white rounded-2xl shadow-sm border border-gray-200 overflow-hidden">
          <div class="bg-gray-50 px-6 py-4 border-b border-gray-200">
            <div class="flex flex-wrap items-center justify-between gap-4">
              <div>
                <h3 class="text-lg font-bold text-gray-900">Đơn hàng #{{ order.id }}</h3>
                <p class="text-sm text-gray-500">{{ formatDate(order.order_date) }}</p>
              </div>
              <div class="flex items-center gap-4">
                <span :class="getStatusClass(order.status)" class="px-4 py-2 rounded-full text-sm font-semibold">
                  {{ getStatusText(order.status) }}
                </span>
                <span class="text-lg font-bold text-blue-600">{{ formatPrice(order.total_amount) }}đ</span>
              </div>
            </div>
          </div>

          <div class="p-6">
            <div class="space-y-4 mb-4">
              <div v-for="item in order.order_items" :key="item.id" class="flex items-center gap-4 pb-4 border-b border-gray-100 last:border-0">
                <img 
                  :src="item.book?.cover_image_url || '/placeholder-book.jpg'" 
                  :alt="item.book?.title"
                  class="w-16 h-20 object-cover rounded-lg bg-gray-100"
                >
                <div class="flex-1">
                  <h4 class="font-semibold text-gray-900 line-clamp-1">{{ item.book?.title }}</h4>
                  <p class="text-sm text-gray-500">Số lượng: {{ item.quantity }}</p>
                </div>
                <div class="text-right">
                  <p class="font-semibold text-gray-900">{{ formatPrice(item.price_at_purchase) }}đ</p>
                  <p class="text-sm text-gray-500">x{{ item.quantity }}</p>
                </div>
              </div>
            </div>

            <div class="bg-gray-50 rounded-lg p-4 mb-4">
              <h4 class="font-semibold text-gray-900 mb-1 text-sm uppercase tracking-wider">Địa chỉ giao hàng</h4>
              <p class="text-gray-600">{{ order.shipping_address }}</p>
            </div>

            <div class="flex gap-3 justify-end">
              <button 
                v-if="order.status === 'pending'"
                @click="cancelOrder(order.id)"
                class="px-4 py-2 border border-red-300 text-red-600 rounded-lg font-semibold hover:bg-red-50 transition-colors"
              >
                Hủy đơn
              </button>
              <router-link 
                :to="`/order/${order.id}`"
                class="px-4 py-2 bg-white border border-gray-300 text-gray-700 rounded-lg font-semibold hover:bg-gray-50 transition-colors"
              >
                Chi tiết
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useAuthStore } from '../stores/user'
import { orderApi } from '../api/order'
import type { Order } from '../interfaces/order.interface'
import { useRouter } from 'vue-router'

const authStore = useAuthStore()
const router = useRouter()
const orders = ref<Order[]>([])
const loading = ref(false)

const formatPrice = (price: number) => {
  return new Intl.NumberFormat('vi-VN').format(price)
}

const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('vi-VN', {
    year: 'numeric', month: 'long', day: 'numeric',
    hour: '2-digit', minute: '2-digit'
  })
}

const getStatusText = (status: string) => {
  const statusMap: Record<string, string> = {
    pending: 'Chờ xử lý',
    shipped: 'Đang giao',
    delivered: 'Đã giao',
    cancelled: 'Đã hủy'
  }
  return statusMap[status] || status
}

const getStatusClass = (status: string) => {
  const classMap: Record<string, string> = {
    pending: 'bg-yellow-100 text-yellow-800',
    shipped: 'bg-blue-100 text-blue-800',
    delivered: 'bg-green-100 text-green-800',
    cancelled: 'bg-red-100 text-red-800'
  }
  return classMap[status] || 'bg-gray-100 text-gray-800'
}

const fetchOrders = async () => {
  if (!authStore.user) {
    router.push('/login')
    return
  }

  loading.value = true
  try {
    // SỬA TẠI ĐÂY: Đảm bảo Backend có route GET /order/user/:id trả về danh sách
    const response = await orderApi.getByUser(authStore.user.id)
    
    // Kiểm tra nếu data là object đơn lẻ (như JSON bạn gửi) thì bọc vào mảng
    const result = response.data.data
    orders.value = Array.isArray(result) ? result : [result]
    
  } catch (error) {
    console.error('Lỗi khi lấy đơn hàng:', error)
  } finally {
    loading.value = false
  }
}

const cancelOrder = async (orderId: number) => {
  if (!confirm('Bạn có chắc muốn hủy đơn hàng này?')) return

  try {
    await orderApi.update(orderId, { status: 'cancelled' })
    await fetchOrders()
  } catch (error) {
    alert('Không thể hủy đơn hàng.')
  }
}

onMounted(() => {
  fetchOrders()
})
</script>