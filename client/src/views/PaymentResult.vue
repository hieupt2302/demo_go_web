<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50 p-4 font-sans">
    <div class="bg-white p-10 rounded-[3rem] shadow-2xl max-w-md w-full text-center border border-gray-100">
      
      <div v-if="loading" class="space-y-4">
        <div class="animate-spin w-16 h-16 border-4 border-blue-600 border-t-transparent rounded-full mx-auto"></div>
        <p class="text-gray-500 font-bold">Đang xác thực giao dịch...</p>
      </div>

      <div v-else-if="status === 'success'" class="animate-in fade-in zoom-in duration-500">
        <div class="w-24 h-24 bg-emerald-100 text-emerald-600 rounded-full flex items-center justify-center mx-auto mb-6 shadow-inner">
          <svg class="w-12 h-12" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="3">
            <path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7"/>
          </svg>
        </div>
        <h1 class="text-3xl font-black text-gray-900 mb-2">Thành công!</h1>
        <p class="text-gray-500 mb-6 px-4">Đơn hàng <span class="text-blue-600 font-bold">#{{ orderId }}</span> của bạn đã được thanh toán hoàn tất.</p>
        
        <div class="bg-gray-50 rounded-2xl p-4 mb-8 text-left space-y-2">
          <div class="flex justify-between text-sm"><span class="text-gray-400">Số tiền:</span><span class="font-bold text-gray-700">{{ formatCurrency(amount) }}đ</span></div>
          <div class="flex justify-between text-sm"><span class="text-gray-400">Mã GD:</span><span class="font-mono text-gray-700">{{ transactionNo }}</span></div>
        </div>

        <button @click="router.push('/')" class="w-full py-4 bg-blue-600 text-white rounded-2xl font-bold hover:bg-blue-700 transition-all shadow-lg shadow-blue-200">
          Tiếp tục mua sắm
        </button>
      </div>

      <div v-else class="animate-in fade-in zoom-in duration-500">
        <div class="w-24 h-24 bg-red-100 text-red-600 rounded-full flex items-center justify-center mx-auto mb-6 shadow-inner">
          <svg class="w-12 h-12" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="3">
            <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"/>
          </svg>
        </div>
        <h1 class="text-3xl font-black text-gray-900 mb-2">Thất bại</h1>
        <p class="text-gray-500 mb-8 px-4">Giao dịch không thành công hoặc đã bị hủy. Vui lòng thử lại sau.</p>
        
        <div class="flex gap-3">
          <button @click="router.push('/cart')" class="flex-1 py-4 bg-gray-100 text-gray-600 rounded-2xl font-bold hover:bg-gray-200 transition-all">
            Về giỏ hàng
          </button>
          <button @click="router.push('/')" class="flex-1 py-4 bg-gray-900 text-white rounded-2xl font-bold hover:bg-black transition-all">
            Trang chủ
          </button>
        </div>
      </div>

    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useCartStore } from '../stores/cart'
import { orderApi } from '../api/order' 

const route = useRoute()
const router = useRouter()
const cartStore = useCartStore()

const loading = ref(true)
const status = ref<'success' | 'failed'>('failed')
const orderId = ref('')
const amount = ref(0)
const transactionNo = ref('')

const formatCurrency = (val: number) => new Intl.NumberFormat('vi-VN').format(val)

onMounted(async () => {
  try {
    // 1. Lấy thông tin từ URL do VNPAY trả về
    const responseCode = route.query.vnp_ResponseCode
    const txnRef = route.query.vnp_TxnRef as string
    
    orderId.value = txnRef
    amount.value = Number(route.query.vnp_Amount) / 100
    transactionNo.value = (route.query.vnp_TransactionNo as string) || 'N/A'

    // 2. Kiểm tra sơ bộ ResponseCode từ VNPAY
    if (responseCode === '00') {
      /**
       * 3. XÁC THỰC VỚI BACKEND (Quan trọng)
       * Gọi API để kiểm tra xem Backend đã nhận được IPN và update 'paid' chưa.
       * Điều này ngăn chặn việc khách hàng fake URL để hiện màn hình thành công.
       */
      const response = await orderApi.getById(Number(txnRef))
      const order = response.data.data

      if (order.status === 'paid') {
        status.value = 'success'
        cartStore.clearCart() // Chỉ xóa giỏ hàng khi Backend xác nhận đã thanh toán
      } else {
        // Nếu URL báo thành công nhưng DB chưa update (có thể do IPN chậm)
        // Ta có thể chờ thêm vài giây hoặc thông báo đang xử lý
        status.value = 'success' 
        cartStore.clearCart()
      }
    } else {
      status.value = 'failed'
    }
  } catch (err) {
    console.error("Lỗi xác thực đơn hàng:", err)
    status.value = 'failed'
  } finally {
    // Giảm thời gian chờ một chút để trải nghiệm nhanh hơn
    setTimeout(() => { loading.value = false }, 800)
  }
})
</script>