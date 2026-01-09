<template>
  <section class="py-12">
    <div class="max-w-7xl mx-auto px-4">
      <!-- Section Header -->
      <div class="flex items-center justify-between mb-8">
        <h2 class="text-3xl font-bold text-gray-900">{{ title }}</h2>
        <div class="flex space-x-2">
          <button 
            @click="scrollLeft" 
            class="p-2 rounded-full bg-gray-100 hover:bg-gray-200 transition-colors"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
            </svg>
          </button>
          <button 
            @click="scrollRight" 
            class="p-2 rounded-full bg-gray-100 hover:bg-gray-200 transition-colors"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
            </svg>
          </button>
        </div>
      </div>

      <!-- Scrollable Books Container -->
      <div 
        ref="scrollContainer"
        class="flex overflow-x-auto space-x-6 pb-4 scrollbar-hide scroll-smooth"
        style="scrollbar-width: none; -ms-overflow-style: none;"
      >
        <div 
          v-for="book in books" 
          :key="book.id"
          class="flex-shrink-0 w-64"
        >
          <BookCard :book="book" @add-to-cart="$emit('add-to-cart', book)" />
        </div>
      </div>

      <!-- Pagination Dots -->
      <div class="flex justify-center mt-6 space-x-2">
        <div 
          v-for="i in Math.ceil(books.length / 5)" 
          :key="i"
          class="w-2 h-2 rounded-full transition-colors"
          :class="i === 1 ? 'bg-blue-600' : 'bg-gray-300'"
        ></div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import BookCard from '../product/BookCard.vue'
import type { Book } from '../../interfaces/book.interface'

defineProps<{
  title: string
  books: Book[]
}>()

defineEmits(['add-to-cart'])

const scrollContainer = ref<HTMLElement | null>(null)

const scrollLeft = () => {
  if (scrollContainer.value) {
    scrollContainer.value.scrollBy({ left: -300, behavior: 'smooth' })
  }
}

const scrollRight = () => {
  if (scrollContainer.value) {
    scrollContainer.value.scrollBy({ left: 300, behavior: 'smooth' })
  }
}
</script>

<style scoped>
.scrollbar-hide::-webkit-scrollbar {
  display: none;
}
</style>
