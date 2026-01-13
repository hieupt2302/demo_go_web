<template>
  <section class="py-12">
    <div class="max-w-7xl mx-auto px-4">
      <div class="flex items-center justify-between mb-8">
        <div>
          <h2 class="text-3xl font-bold text-gray-900">{{ title }}</h2>
          <div class="h-1 w-20 bg-blue-600 mt-2 rounded-full"></div>
        </div>
        
        <div class="flex space-x-3">
          <button 
            @click="scrollLeft" 
            class="p-2.5 rounded-full bg-white border border-gray-200 shadow-sm hover:bg-gray-50 hover:text-blue-600 transition-all active:scale-95"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
            </svg>
          </button>
          <button 
            @click="scrollRight" 
            class="p-2.5 rounded-full bg-white border border-gray-200 shadow-sm hover:bg-gray-50 hover:text-blue-600 transition-all active:scale-95"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
            </svg>
          </button>
        </div>
      </div>

      <div v-if="loading" class="flex space-x-6 overflow-hidden">
        <div v-for="i in 5" :key="i" class="flex-shrink-0 w-64 animate-pulse">
          <div class="bg-gray-200 h-80 rounded-2xl mb-4"></div>
          <div class="h-4 bg-gray-200 rounded w-3/4 mb-2"></div>
          <div class="h-4 bg-gray-200 rounded w-1/2"></div>
        </div>
      </div>

      <div 
        v-else
        ref="scrollContainer"
        class="flex overflow-x-auto space-x-6 pb-8 scrollbar-hide scroll-smooth"
        style="scrollbar-width: none; -ms-overflow-style: none;"
      >
        <div 
          v-for="book in books" 
          :key="book.id"
          class="flex-shrink-0 w-64 group"
        >
          <BookCard 
            :book="book" 
            @add-to-cart="$emit('add-to-cart', book)" 
          />
        </div>

        <div v-if="books.length === 0" class="w-full py-12 text-center text-gray-500">
          Hiện tại không có sách nào trong danh mục này.
        </div>
      </div>

      <div v-if="books.length > 0" class="flex justify-center mt-2 space-x-2">
        <div 
          v-for="i in Math.ceil(books.length / 4)" 
          :key="i"
          class="w-1.5 h-1.5 rounded-full bg-gray-300 transition-all duration-300"
        ></div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import BookCard from '../product/BookCard.vue'
import type { Book } from '../../interfaces/book.interface'

// Thêm loading vào props để đồng bộ với HomeView
defineProps<{
  title: string
  books: Book[]
  loading?: boolean
}>()

defineEmits(['add-to-cart'])

const scrollContainer = ref<HTMLElement | null>(null)

const scrollLeft = () => {
  if (scrollContainer.value) {
    scrollContainer.value.scrollBy({ left: -320, behavior: 'smooth' })
  }
}

const scrollRight = () => {
  if (scrollContainer.value) {
    scrollContainer.value.scrollBy({ left: 320, behavior: 'smooth' })
  }
}
</script>

<style scoped>
.scrollbar-hide::-webkit-scrollbar {
  display: none;
}
</style>