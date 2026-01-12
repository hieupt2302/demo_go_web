<template>
  <div class="bg-gray-50 min-h-screen">
    <HeroSection />

    <BookSection 
      title="Best Seller Books" 
      :books="books" 
      :loading="loading"
      @add-to-cart="handleAddToCart" 
    />

    <FeaturedSection />
    <AwardsSection />

    <Toast v-if="toastMessage" :message="toastMessage" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useCartStore } from '../stores/cart'
import { bookApi } from '../api/book'
import { userApi } from '../api/user'
import type { Book } from '../interfaces/book.interface'
import HeroSection from '../components/home/HeroSection.vue'
import BookSection from '../components/home/BookSection.vue'
import FeaturedSection from '../components/home/FeaturedSection.vue'
import AwardsSection from '../components/home/AwardsSection.vue'
import Toast from '../components/common/Toast.vue'
import { useAuthStore } from '../stores/user'

const route = useRoute()
const router = useRouter()
const cartStore = useCartStore()
const authStore = useAuthStore()
const books = ref<Book[]>([])
const loading = ref(false)
const toastMessage = ref('')

const fetchUserData = async () => {
  const userId = route.query.uid
  console.log(userId)
  
  // Nếu có UID và trong Store chưa có thông tin user
  if (userId && !authStore.user) {
    try {
      const res = await userApi.getUserById(Number(userId))
      // Cập nhật thông tin user vào store (để hiển thị fullname trên Header)
      authStore.setUser(res.data.data)
      
      // Sau khi lấy xong, xóa uid khỏi URL cho đẹp (optional)
      router.replace({ query: {} })
    } catch (error) {
      console.error("Lỗi lấy thông tin user:", error)
    }
  }
}

const fetchBooks = async () => {
  loading.value = true
  try {
    const res = await bookApi.getAll()
    books.value = res.data.data 
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

onMounted(async () => {
  await fetchUserData()
  await fetchBooks()
})
</script>