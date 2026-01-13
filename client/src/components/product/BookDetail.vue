<template>
  <div class="bg-gray-50 min-h-screen py-12">
    <div class="max-w-7xl mx-auto px-4">
      <button @click="router.back()" class="flex items-center text-gray-600 hover:text-blue-600 mb-8 transition-colors group">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2 group-hover:-translate-x-1 transition-transform" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
        </svg>
        Quay lại
      </button>

      <div v-if="loading" class="flex justify-center py-20">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
      </div>

      <div v-else-if="book" class="bg-white rounded-3xl shadow-xl overflow-hidden">
        <div class="flex flex-col md:flex-row">
          <div class="md:w-2/5 bg-gray-100 p-8 flex justify-center items-center">
            <img 
              :src="getImageUrl(book.cover_image_url)" 
              class="w-full max-w-sm rounded-2xl shadow-2xl hover:scale-105 transition-transform duration-500" 
              alt="book cover"
            />
          </div>

          <div class="md:w-3/5 p-8 md:p-12">
            <div class="flex flex-wrap gap-2 mb-4">
              <span 
                v-for="cat in book.categories" :key="cat.id"
                class="bg-blue-50 text-blue-600 text-xs font-bold px-3 py-1.5 rounded-full border border-blue-100 uppercase tracking-wider"
              >
                {{ cat.name }}
              </span>
            </div>

            <h1 class="text-4xl font-black text-gray-900 mb-4 leading-tight">{{ book.title }}</h1>
            
            <div class="flex items-center mb-6 space-x-4">
              <span class="text-3xl font-bold text-blue-600">{{ formatPrice(book.price) }}đ</span>
              <span v-if="book.stock_quantity > 0" class="text-sm text-green-600 bg-green-50 px-3 py-1 rounded-lg font-medium">
                Còn hàng ({{ book.stock_quantity }})
              </span>
              <span v-else class="text-sm text-red-600 bg-red-50 px-3 py-1 rounded-lg font-medium">Hết hàng</span>
            </div>

            <div class="mb-8">
              <h3 class="text-sm font-bold text-gray-400 uppercase tracking-widest mb-3">Tác giả</h3>
              <div class="flex flex-wrap gap-4">
                <div v-for="author in book.authors" :key="author.id" class="flex items-center group cursor-pointer">
                  <div class="w-10 h-10 rounded-full bg-blue-100 flex items-center justify-center text-blue-600 font-bold mr-3 group-hover:bg-blue-600 group-hover:text-white transition-colors">
                    {{ author.name.charAt(0) }}
                  </div>
                  <div>
                    <p class="text-gray-900 font-bold leading-none">{{ author.name }}</p>
                    <p class="text-xs text-gray-500 mt-1">Xem tác phẩm khác</p>
                  </div>
                </div>
              </div>
            </div>

            <div class="grid grid-cols-2 gap-6 mb-8 py-6 border-y border-gray-100">
              <div>
                <p class="text-xs text-gray-400 uppercase font-bold mb-1">ISBN</p>
                <p class="text-gray-900 font-medium">{{ book.isbn || 'N/A' }}</p>
              </div>
              <div>
                <p class="text-xs text-gray-400 uppercase font-bold mb-1">Nhà xuất bản</p>
                <p class="text-gray-900 font-medium">{{ book.publisher || 'N/A' }}</p>
              </div>
            </div>

            <div class="mb-8">
              <h3 class="text-sm font-bold text-gray-400 uppercase tracking-widest mb-3">Mô tả nội dung</h3>
              <p class="text-gray-600 leading-relaxed whitespace-pre-line">{{ book.description || 'Đang cập nhật nội dung...' }}</p>
            </div>

            <button 
              @click="handleAddToCart"
              :disabled="book.stock_quantity === 0"
              class="w-full md:w-auto px-12 py-4 bg-gray-900 text-white font-bold rounded-2xl hover:bg-blue-600 disabled:bg-gray-300 shadow-xl transition-all active:scale-95 flex items-center justify-center gap-3"
            >
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
              </svg>
              Thêm vào giỏ hàng
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { bookApi } from '../../api/book'
import { useCartStore } from '../../stores/cart'
import type { Book } from '../../interfaces/book.interface'

const route = useRoute()
const router = useRouter()
const cartStore = useCartStore()

const book = ref<Book | null>(null)
const loading = ref(true)

const formatPrice = (p: number) => new Intl.NumberFormat('vi-VN').format(p)

const getImageUrl = (path: string) => {
  if (!path) return '/placeholder-book.jpg'
  if (path.startsWith('http')) return path
  return `${import.meta.env.VITE_API_BASE_URL}/${path}`.replace('/api/', '/').replace(/\\/g, '/')
}

const fetchBookDetail = async () => {
  const id = route.params.id
  loading.value = true
  try {
    const res = await bookApi.getById(Number(id))
    book.value = res.data.data
  } catch (error) {
    console.error("Lỗi lấy chi tiết sách:", error)
  } finally {
    loading.value = false
  }
}

const handleAddToCart = () => {
  if (book.value) {
    cartStore.addToCart(book.value)
    alert(`Đã thêm "${book.value.title}" vào giỏ hàng!`)
  }
}

onMounted(fetchBookDetail)
</script>