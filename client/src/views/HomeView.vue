<template>
  <div class="bg-gray-50">
    <nav class="bg-white shadow-sm py-4 px-8 flex justify-between items-center">
      <div class="text-xl font-bold text-blue-600">BOOKSTORE</div>
      
      <div class="flex items-center gap-4">
        <div v-if="authStore.isAuthenticated" class="flex items-center gap-3">
          <div class="flex flex-col items-end">
            <span class="text-sm font-medium text-gray-700">Xin chào,</span>
            <span class="text-blue-600 font-bold">{{ authStore.user?.fullname }}</span>
          </div>
          <button 
            class="text-xs text-red-500 hover:underline ml-2"
          >
            Đăng xuất
          </button>
        </div>

        <div v-else class="flex gap-2">
          <router-link 
            to="/login" 
            class="px-4 py-2 text-sm font-medium text-blue-600 border border-blue-600 rounded-lg hover:bg-blue-50"
          >
            Đăng nhập
          </router-link>
          <router-link 
            to="/register" 
            class="px-4 py-2 text-sm font-medium text-white bg-blue-600 rounded-lg hover:bg-blue-700"
          >
            Đăng ký
          </router-link>
        </div>
      </div>
    </nav>

    <HeroSection />
    <BookSection title="Best Seller Books" :books="books" @add-to-cart="handleAddToCart" />
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useCartStore } from '../stores/cart'
import { bookApi } from '../api/book'
import type { Book } from '../interfaces/book.interface'
import HeroSection from '../components/home/HeroSection.vue'
import BookSection from '../components/home/BookSection.vue'
import FeaturedSection from '../components/home/FeaturedSection.vue'
import AwardsSection from '../components/home/AwardsSection.vue'
import Toast from '../components/common/Toast.vue'
import { useAuthStore } from '../stores/user'

const cartStore = useCartStore()
const authStore = useAuthStore()
const books = ref<Book[]>([])
const loading = ref(false)
const toastMessage = ref('')
console.log('Auth Store:', authStore)

const fetchBooks = async () => {
  loading.value = true
  try {
    const res = await bookApi.getAll()
    books.value = res.data 
  } catch (error) {
    console.error("Lỗi API:", error)
  } finally {
    loading.value = false
  }
}

const handleAddToCart = (book: Book) => {
  cartStore.addToCart(book)
  toastMessage.value = `Đã thêm "${book.title}" vào giỏ hàng!`
  setTimeout(() => {
    toastMessage.value = ''
  }, 100)
}

onMounted(fetchBooks)
</script>