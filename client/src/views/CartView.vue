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
  </div>
</template>

<script setup lang="ts">
import { useCartStore } from '../stores/cart'
import { useRouter } from 'vue-router'
import CartItem from '../components/cart/CartItem.vue'

const cartStore = useCartStore()
const router = useRouter()

const formatPrice = (price: number) => {
  return new Intl.NumberFormat('vi-VN').format(price)
}

const handleRemove = (bookId: number) => {
  if (confirm('Bạn có chắc muốn xóa sản phẩm này khỏi giỏ hàng?')) {
    cartStore.removeFromCart(bookId)
  }
}

const handleCheckout = () => {
  alert('Chức năng thanh toán đang được phát triển!')
  // TODO: Implement checkout logic
  // cartStore.clearCart()
  // router.push('/order-success')
}
</script>
