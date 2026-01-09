<template>
  <div class="bg-gray-50">
    <!-- Hero Section -->
    <HeroSection />

    <!-- Best Seller Books -->
    <BookSection 
      title="Best Seller Books" 
      :books="books" 
      @add-to-cart="handleAddToCart"
    />

    <!-- Featured Section -->
    <FeaturedSection />

    <!-- Awards Section -->
    <AwardsSection />

    <!-- New Releases -->
    <BookSection 
      title="New Releases" 
      :books="books" 
      @add-to-cart="handleAddToCart"
    />

    <!-- Toast Notification -->
    <Toast :message="toastMessage" />
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

const cartStore = useCartStore()
const books = ref<Book[]>([])
const loading = ref(false)
const toastMessage = ref('')

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