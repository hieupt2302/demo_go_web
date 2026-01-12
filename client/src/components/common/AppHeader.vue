<template>
  <header class="bg-white border-b border-gray-200 sticky top-0 z-50 shadow-sm">
    <div class="max-w-7xl mx-auto px-4 py-4">
      <div class="flex items-center justify-between gap-4">
        <router-link to="/" class="text-2xl font-bold text-gray-900 shrink-0">Books</router-link>

        <!-- SEARCH BAR -->
        <div class="relative flex-1 max-w-lg mx-4 hidden lg:block">
          <div class="flex items-center bg-gray-100 rounded-lg overflow-hidden border border-transparent focus-within:border-blue-500 transition-all">
            <select 
              v-model="searchType" 
              class="bg-gray-200 px-3 py-2 text-sm text-gray-700 outline-none border-r border-gray-300 cursor-pointer"
            >
              <option value="book">Sách</option>
              <option value="author">Tác giả</option>
              <option value="category">Thể loại</option>
            </select>
            
            <input 
              v-model="searchQuery"
              type="text" 
              placeholder="Tìm kiếm nhanh..." 
              class="bg-transparent flex-1 px-4 py-2 text-sm outline-none"
              @focus="showDropdown = true"
            />
            
            <div class="p-2 text-gray-400">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </div>
          </div>

          <!-- DROPDOWN RESULT -->
          <div 
            v-if="showDropdown && (searchResults.length > 0 || isSearching)" 
            v-click-outside="() => showDropdown = false"
            class="absolute top-full left-0 right-0 mt-2 bg-white border border-gray-200 rounded-lg shadow-xl z-[60] max-h-96 overflow-y-auto"
          >
            <div v-if="isSearching" class="p-4 text-center text-gray-500 text-sm italic">
              Đang tìm kiếm...
            </div>
            
            <ul v-else>
              <li 
                v-for="book in searchResults" 
                :key="book.id"
                @click="goToBook(book.id)"
                class="flex items-center gap-4 p-3 hover:bg-blue-50 cursor-pointer border-b border-gray-50 last:border-0 transition-colors"
              >
                <img :src="book.cover_image_url || '/placeholder.jpg'" class="w-10 h-14 object-cover rounded shadow-sm" />
                <div class="flex-1 min-w-0">
                  <p class="text-sm font-bold text-gray-900 truncate">{{ book.title }}</p>
                  <p class="text-xs text-gray-500 truncate">
                    {{ book.author?.name }} | {{ book.category?.name }}
                  </p>
                </div>
                <div class="text-blue-600 font-bold text-sm shrink-0">
                  {{ formatPrice(book.price) }}đ
                </div>
              </li>
            </ul>
          </div>
        </div>

        <!-- RIGHT MENU -->
        <div class="flex items-center space-x-4">
          <nav class="hidden md:flex items-center space-x-8 mr-4">
            <router-link to="/" class="text-gray-700 hover:text-blue-600 transition-colors font-medium text-sm">Trang chủ</router-link>
          </nav>

          <div v-if="authStore.isAuthenticated" class="flex items-center space-x-3">
            <router-link to="/orders" class="text-gray-700 hover:text-blue-600 transition-colors font-medium text-sm">
              Đơn hàng của tôi
            </router-link>
            <span class="text-gray-700 font-medium text-sm">
              Chào, {{ authStore.user?.fullname || 'Thành viên' }}
            </span>
            <button @click="handleLogout" class="text-sm text-red-500 hover:text-red-700 font-medium">
              Đăng xuất
            </button>
          </div>

          <router-link v-else to="/login" class="text-gray-700 hover:text-blue-600 transition-colors font-medium text-sm">
            Đăng nhập
          </router-link>
          
          <router-link to="/cart" class="relative p-2 hover:bg-gray-100 rounded-full transition-colors">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-gray-700" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
            </svg>
            <span v-if="cartStore.totalItems > 0" class="absolute -top-1 -right-1 bg-red-500 text-white text-xs font-bold rounded-full h-5 w-5 flex items-center justify-center">
              {{ cartStore.totalItems }}
            </span>
          </router-link>
        </div>
      </div>
    </div>
  </header>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { useCartStore } from '../../stores/cart'
import { useAuthStore } from '../../stores/user'
import { useRouter } from 'vue-router'
import { bookApi } from '../../api/book'
import type { Book } from '../../interfaces/book.interface'

const cartStore = useCartStore()
const authStore = useAuthStore()
const router = useRouter()

// ===== SEARCH STATE =====
const searchQuery = ref('')
const searchType = ref('book')
const searchResults = ref<Book[]>([])
const isSearching = ref(false)
const showDropdown = ref(false)

const formatPrice = (p: number) => new Intl.NumberFormat('vi-VN').format(p)

// ===== DEBOUNCE TIMER =====
let debounceTimer: any = null

// ===== WATCH SEARCH QUERY WITH DEBOUNCE =====
watch(searchQuery, (newValue) => {
  showDropdown.value = true

  if (debounceTimer) {
    clearTimeout(debounceTimer)
  }

  debounceTimer = setTimeout(async () => {
    if (!newValue.trim()) {
      searchResults.value = []
      isSearching.value = false
      return
    }

    isSearching.value = true
    try {
      const res = await bookApi.search({
        type: searchType.value,
        q: newValue.trim()
      })
      searchResults.value = res.data.data
    } catch (error) {
      console.error("Lỗi search:", error)
      searchResults.value = []
    } finally {
      isSearching.value = false
    }
  }, 300) // ⏱️ 300ms debounce
})

// ===== NAVIGATE TO BOOK =====
const goToBook = (id: number) => {
  showDropdown.value = false
  searchQuery.value = ''
  router.push(`/book/${id}`)
}

// ===== LOGOUT =====
const handleLogout = () => {
  authStore.logout()
  router.push('/login')
}

// ===== CLICK OUTSIDE DIRECTIVE =====
const vClickOutside = {
  mounted(el: any, binding: any) {
    el.clickOutsideEvent = (event: Event) => {
      if (!(el === event.target || el.contains(event.target))) {
        binding.value()
      }
    }
    document.addEventListener('click', el.clickOutsideEvent)
  },
  unmounted(el: any) {
    document.removeEventListener('click', el.clickOutsideEvent)
  }
}
</script>
