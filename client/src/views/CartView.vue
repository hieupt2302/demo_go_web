

<script setup lang="ts">
import { ref } from 'vue'
import { useCartStore } from '../stores/cart'
import { useAuthStore } from '../stores/user'
import { useRouter } from 'vue-router'
import { orderApi } from '../api/order'
import CartItem from '../components/cart/CartItem.vue'

const cartStore = useCartStore()
const authStore = useAuthStore()
const router = useRouter()

const showCheckoutModal = ref(false)
const shippingAddress = ref('')
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
  // Lấy địa chỉ từ user nếu có
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
    const orderData = {
      user_id: authStore.user!.id,
      total_amount: cartStore.totalPrice,
      status: 'pending',
      shipping_address: shippingAddress.value,
      order_items: cartStore.items.map(item => ({
        book_id: item.book.id,
        quantity: item.quantity,
        price: item.book.price
      }))
    }

    const response = await orderApi.create(orderData)
    
    if (response.data) {
      // Xóa giỏ hàng sau khi đặt hàng thành công
      cartStore.clearCart()
      
      // Chuyển đến trang chi tiết đơn hàng hoặc trang thành công
      alert('Đặt hàng thành công! Mã đơn hàng: #' + response.data.data.id)
      router.push('/')
    }
  } catch (err: any) {
    error.value = err.response?.data?.error || 'Đặt hàng thất bại. Vui lòng thử lại.'
    console.error('Lỗi đặt hàng:', err)
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
    <div v-if="showCheckoutModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
      <div class="bg-white rounded-2xl shadow-2xl max-w-md w-full p-8">
        <h2 class="text-2xl font-bold text-gray-900 mb-6">Thông tin giao hàng</h2>
        
        <div class="mb-6">
          <label class="block text-sm font-medium text-gray-700 mb-2">
            Địa chỉ giao hàng <span class="text-red-500">*</span>
          </label>
          <textarea
            v-model="shippingAddress"
            rows="4"
            class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="Nhập địa chỉ đầy đủ của bạn..."
          ></textarea>
        </div>

        <div v-if="error" class="mb-4 p-3 bg-red-50 border border-red-200 rounded-lg">
          <p class="text-red-600 text-sm">{{ error }}</p>
        </div>

        <div class="bg-gray-50 rounded-lg p-4 mb-6">
          <div class="flex justify-between text-sm text-gray-600 mb-2">
            <span>Tổng tiền hàng:</span>
            <span class="font-semibold">{{ formatPrice(cartStore.totalPrice) }}đ</span>
          </div>
          <div class="flex justify-between text-lg font-bold text-gray-900">
            <span>Tổng thanh toán:</span>
            <span class="text-blue-600">{{ formatPrice(cartStore.totalPrice) }}đ</span>
          </div>
        </div>

        <div class="flex gap-3">
          <button
            @click="closeModal"
            :disabled="loading"
            class="flex-1 px-6 py-3 border border-gray-300 text-gray-700 rounded-lg font-semibold hover:bg-gray-50 transition-colors disabled:opacity-50"
          >
            Hủy
          </button>
          <button
            @click="submitOrder"
            :disabled="loading"
            class="flex-1 px-6 py-3 bg-blue-600 text-white rounded-lg font-semibold hover:bg-blue-700 transition-colors disabled:opacity-50"
          >
            {{ loading ? 'Đang xử lý...' : 'Đặt hàng' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
