<script setup lang="ts">
import { ref } from 'vue'
import { useCartStore } from '../stores/cart'
import { useAuthStore } from '../stores/user'
import { useRouter } from 'vue-router'
import { orderApi } from '../api/order'
import CartItem from '../components/cart/CartItem.vue'
import type { CreateOrderInput, PaymentMethod } from '../interfaces/order.interface'

const cartStore = useCartStore()
const authStore = useAuthStore()
const router = useRouter()

const showCheckoutModal = ref(false)
const shippingAddress = ref('')
const paymentMethod = ref<PaymentMethod>('VNPAY') // Mặc định chọn VNPAY
const loading = ref(false)
const error = ref('')

const formatPrice = (price: number) => {
  return new Intl.NumberFormat('vi-VN').format(price)
}

const handleRemove = (bookId: number) => {
  if (confirm('Bạn có chắc muốn xóa sản phẩm này khỏi giỏ hàng?')) {
    cartStore.removeFromCart(bookId)
  }
}

const handleCheckout = () => {
  if (!authStore.isAuthenticated) {
    alert('Vui lòng đăng nhập để thanh toán!')
    router.push('/login')
    return
  }
  showCheckoutModal.value = true
  shippingAddress.value = authStore.user?.address || ''
}

const submitOrder = async () => {
  if (!shippingAddress.value.trim()) {
    error.value = 'Vui lòng nhập địa chỉ giao hàng'
    return
  }

  loading.value = true
  error.value = ''

  try {
    // Khớp hoàn toàn với CreateOrderInput interface
    const orderData: CreateOrderInput = {
      shipping_address: shippingAddress.value,
      payment_method: paymentMethod.value,
      items: cartStore.items.map(item => ({
        book_id: item.book.id,
        quantity: item.quantity
      }))
    }

    const response = await orderApi.create(orderData)
    
    // bóc tách dữ liệu từ ApiResponse<{ order_id, payment_url }>
    const { order_id, payment_url } = response.data.data

    // Xóa giỏ hàng khi thành công
    cartStore.clearCart()
    showCheckoutModal.value = false

    // Logic điều hướng dựa trên phản hồi Backend
    if (paymentMethod.value === 'BANK_TRANSFER') {
      // Chuyển đến trang hiển thị QR nội bộ
      router.push(`/checkout/bank-info/${order_id}`)
    } else if (payment_url) {
      // Chuyển hướng đến link VNPAY/Momo (URL ngoại sàn)
      window.location.href = payment_url
    } else {
      alert('Đặt hàng thành công! Mã đơn hàng: #' + order_id)
      router.push('/')
    }
  } catch (err: any) {
    error.value = err.response?.data?.message || 'Đặt hàng thất bại. Vui lòng thử lại.'
  } finally {
    loading.value = false
  }
}

const closeModal = () => {
  showCheckoutModal.value = false
  error.value = ''
}
</script>

<template>
  <div class="bg-gray-50 min-h-screen py-12">
    <div class="max-w-7xl mx-auto px-4">
      <!-- Header -->
      <div class="mb-8">
        <h1 class="text-4xl font-extrabold text-gray-900">Giỏ hàng của bạn</h1>
        <p class="text-gray-500 mt-2">{{ cartStore.totalItems }} sản phẩm trong giỏ</p>
      </div>

      <!-- Empty Cart State -->
      <div v-if="cartStore.items.length === 0" class="bg-white rounded-2xl shadow-sm border border-gray-200 p-12 text-center">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-24 w-24 text-gray-300 mx-auto mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
        </svg>
        <h2 class="text-2xl font-bold text-gray-900 mb-2">Giỏ hàng trống</h2>
        <p class="text-gray-500 mb-6">Bạn chưa có sản phẩm nào trong giỏ hàng</p>
        <router-link 
          to="/" 
          class="inline-block bg-blue-600 hover:bg-blue-700 text-white px-8 py-3 rounded-lg font-semibold transition-colors"
        >
          Tiếp tục mua sắm
        </router-link>
      </div>

      <!-- Cart with Items -->
      <div v-else class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <!-- Cart Items -->
        <div class="lg:col-span-2 space-y-4">
          <CartItem 
            v-for="item in cartStore.items" 
            :key="item.book.id"
            :item="item"
            @increment="cartStore.incrementQuantity"
            @decrement="cartStore.decrementQuantity"
            @remove="handleRemove"
          />
        </div>

        <!-- Cart Summary -->
        <div class="lg:col-span-1">
          <div class="bg-white rounded-2xl shadow-sm border border-gray-200 p-6 sticky top-24">
            <h2 class="text-2xl font-bold text-gray-900 mb-6">Tóm tắt đơn hàng</h2>
            
            <div class="space-y-4 mb-6">
              <div class="flex justify-between text-gray-600">
                <span>Tạm tính</span>
                <span class="font-semibold">{{ formatPrice(cartStore.totalPrice) }}đ</span>
              </div>
              <div class="flex justify-between text-gray-600">
                <span>Phí vận chuyển</span>
                <span class="font-semibold">Miễn phí</span>
              </div>
              <div class="border-t border-gray-200 pt-4">
                <div class="flex justify-between text-xl font-bold text-gray-900">
                  <span>Tổng cộng</span>
                  <span class="text-blue-600">{{ formatPrice(cartStore.totalPrice) }}đ</span>
                </div>
              </div>
            </div>

            <button 
              @click="handleCheckout"
              class="w-full bg-blue-600 hover:bg-blue-700 text-white py-4 rounded-lg font-bold text-lg transition-colors shadow-lg hover:shadow-xl mb-4"
            >
              Thanh toán
            </button>

            <router-link 
              to="/" 
              class="block text-center text-blue-600 hover:text-blue-700 font-semibold"
            >
              ← Tiếp tục mua sắm
            </router-link>
          </div>
        </div>
      </div>
    </div>

    <!-- Checkout Modal -->
    <div v-if="showCheckoutModal" class="fixed inset-0 bg-black/60 backdrop-blur-sm flex items-center justify-center z-50 p-4">
      <div class="bg-white rounded-[2rem] shadow-2xl max-w-lg w-full p-8 max-h-[95vh] overflow-y-auto">
        <h2 class="text-2xl font-bold text-gray-900 mb-6">Xác nhận thanh toán</h2>
        
        <div class="mb-6">
          <label class="block text-sm font-bold text-gray-700 mb-2">Địa chỉ giao hàng *</label>
          <textarea
            v-model="shippingAddress"
            rows="3"
            class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-blue-500 outline-none transition-all"
            placeholder="Số nhà, tên đường, phường/xã..."
          ></textarea>
        </div>

        <div class="mb-6">
          <label class="block text-sm font-bold text-gray-700 mb-3">Chọn phương thức thanh toán *</label>
          <div class="grid grid-cols-1 gap-3">
            <label :class="['flex items-center p-4 border-2 rounded-2xl cursor-pointer transition-all', paymentMethod === 'VNPAY' ? 'border-blue-600 bg-blue-50' : 'border-gray-100 hover:bg-gray-50']">
              <input type="radio" value="VNPAY" v-model="paymentMethod" class="hidden" />
              <img src="https://sandbox.vnpayment.vn/paymentv2/Images/brands/logo-vnpay.png" class="h-6 mr-3" alt="VNPAY" />
              <span class="flex-1 font-bold text-gray-700 text-sm">Cổng thanh toán VNPAY</span>
              <div class="w-5 h-5 border-2 rounded-full flex items-center justify-center" :class="paymentMethod === 'VNPAY' ? 'border-blue-600' : 'border-gray-300'">
                <div v-if="paymentMethod === 'VNPAY'" class="w-2.5 h-2.5 bg-blue-600 rounded-full"></div>
              </div>
            </label>

            <label :class="['flex items-center p-4 border-2 rounded-2xl cursor-pointer transition-all', paymentMethod === 'BANK_TRANSFER' ? 'border-blue-600 bg-blue-50' : 'border-gray-100 hover:bg-gray-50']">
              <input type="radio" value="BANK_TRANSFER" v-model="paymentMethod" class="hidden" />
              <div class="w-6 h-6 bg-blue-100 rounded flex items-center justify-center mr-3 text-blue-600">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"/></svg>
              </div>
              <span class="flex-1 font-bold text-gray-700 text-sm">Chuyển khoản VietQR</span>
              <div class="w-5 h-5 border-2 rounded-full flex items-center justify-center" :class="paymentMethod === 'BANK_TRANSFER' ? 'border-blue-600' : 'border-gray-300'">
                <div v-if="paymentMethod === 'BANK_TRANSFER'" class="w-2.5 h-2.5 bg-blue-600 rounded-full"></div>
              </div>
            </label>
          </div>
        </div>

        <div v-if="error" class="mb-4 p-3 bg-red-50 border border-red-200 rounded-xl text-red-600 text-xs font-semibold">
          {{ error }}
        </div>

        <div class="bg-gray-900 rounded-2xl p-5 mb-8 text-white shadow-xl">
          <div class="flex justify-between items-center">
            <span class="text-gray-400 font-medium">Tổng thanh toán:</span>
            <span class="text-2xl font-black text-blue-400">{{ formatPrice(cartStore.totalPrice) }}đ</span>
          </div>
        </div>

        <div class="flex gap-4">
          <button @click="closeModal" :disabled="loading" class="flex-1 py-4 text-gray-500 font-bold hover:bg-gray-100 rounded-2xl transition-all">
            Hủy
          </button>
          <button @click="submitOrder" :disabled="loading" class="flex-1 py-4 bg-blue-600 text-white font-bold rounded-2xl shadow-lg shadow-blue-200 hover:bg-blue-700 active:scale-95 transition-all flex items-center justify-center gap-2">
            <svg v-if="loading" class="animate-spin h-5 w-5 text-white" fill="none" viewBox="0 0 24 24"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path></svg>
            {{ loading ? 'Đang xử lý...' : 'Xác nhận đặt hàng' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
