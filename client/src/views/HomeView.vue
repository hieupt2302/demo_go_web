<template>
  <div class="bg-gray-50 min-h-screen">
    <HeroSection />

    <BookSection 
      title="Sách Bán Chạy" 
      :books="books" 
      :loading="loading"
      @add-to-cart="handleAddToCart" 
    />

    <FeaturedSection />
    <AwardsSection />

    <Transition name="fade">
      <Toast v-if="toastMessage" :message="toastMessage" />
    </Transition>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useCartStore } from '../stores/cart'
import { bookApi } from '../api/book'
import { userApi } from '../api/user'
import { useAuthStore } from '../stores/user'
import type { Book } from '../interfaces/book.interface'

// Components
import HeroSection from '../components/home/HeroSection.vue'
import BookSection from '../components/home/BookSection.vue'
import FeaturedSection from '../components/home/FeaturedSection.vue'
import AwardsSection from '../components/home/AwardsSection.vue'
import Toast from '../components/common/Toast.vue'

const route = useRoute()
const router = useRouter()
const cartStore = useCartStore()
const authStore = useAuthStore()

const books = ref<Book[]>([])
const loading = ref(false)
const toastMessage = ref('')

/**
 * Xử lý thông tin người dùng từ URL (sau khi đăng nhập/đăng ký thành công)
 */
const fetchUserData = async () => {
  const userId = route.query.uid
  if (userId && !authStore.user) {
    try {
      const res = await userApi.getUserById(Number(userId))
      authStore.setUser(res.data.data)
      
      // Xóa UID khỏi URL để tránh lộ ID người dùng khi copy link
      router.replace({ query: {} })
    } catch (error) {
      console.error("Lỗi lấy thông tin user:", error)
    }
  }
}

/**
 * Lấy danh sách sách từ API
 */
const fetchBooks = async () => {
  loading.value = true
  try {
    const res = await bookApi.getAll()
    // Backend đã sử dụng Preload("Authors").Preload("Categories") nên dữ liệu đã sẵn sàng
    books.value = res.data.data 
  } catch (error) {
    console.error("Lỗi API Books:", error)
  } finally {
    loading.value = false
  }
}

/**
 * Thêm vào giỏ hàng và hiển thị thông báo
 */
const handleAddToCart = (book: Book) => {
  cartStore.addToCart(book)
  toastMessage.value = `Đã thêm "${book.title}" vào giỏ hàng!`
  
  // Reset toast sau 3 giây để đủ thời gian hiển thị
  setTimeout(() => {
    toastMessage.value = ''
  }, 3000)
}

onMounted(async () => {
  // Ưu tiên lấy thông tin user trước để Header hiển thị đúng
  await fetchUserData()
  await fetchBooks()
  const status = sessionStorage.getItem('payment_status')
  const orderId = sessionStorage.getItem('payment_order_id')

  if (status === 'success') {
    // Hiện alert khi người dùng đã thực sự đứng ở trang chủ
    alert(`Chúc mừng! Thanh toán thành công đơn hàng #${orderId}.`)
    
    // Quan trọng: Xóa trạng thái để alert không hiện lại khi F5 trang chủ
    sessionStorage.removeItem('payment_status')
    sessionStorage.removeItem('payment_order_id')
  } else if (status === 'failed') {
    alert("Thanh toán không thành công. Vui lòng kiểm tra lại đơn hàng.")
    sessionStorage.removeItem('payment_status')
  }
})


</script>

<style scoped>
.fade-enter-active, .fade-leave-active {
  transition: opacity 0.5s ease;
}
.fade-enter-from, .fade-leave-to {
  opacity: 0;
}
</style>