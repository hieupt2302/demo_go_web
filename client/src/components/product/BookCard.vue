<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-100 overflow-hidden hover:shadow-lg transition-all flex flex-col h-full group">
    <div class="relative h-64 bg-gray-100 overflow-hidden">
      <img 
        :src="book.cover_image_url || 'https://images.unsplash.com/photo-1544947950-fa07a98d237f?w=300&h=400&fit=crop'" 
        class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
        alt="book cover"
      />
      <span v-if="book.category" class="absolute top-3 right-3 bg-blue-600 text-white text-[10px] px-3 py-1 rounded-full uppercase font-bold shadow-md">
        {{ book.category.name }}
      </span>
    </div>

    <div class="p-4 flex flex-col flex-grow">
      <!-- Price Badge -->
      <div class="flex items-center justify-between mb-2">
        <span class="text-xl font-black text-blue-600">{{ formatPrice(book.price) }}đ</span>
        <div class="flex items-center space-x-1">
          <svg v-for="i in 5" :key="i" xmlns="http://www.w3.org/2000/svg" class="h-3 w-3 text-yellow-400" viewBox="0 0 20 20" fill="currentColor">
            <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
          </svg>
        </div>
      </div>

      <h3 class="font-bold text-gray-900 mb-1 line-clamp-2 leading-tight text-sm h-10">
        {{ book.title }}
      </h3>

      <p class="text-xs text-gray-500 mb-3">
        {{ book.author?.name || 'Ẩn danh' }}
      </p>

      <p class="text-gray-400 text-[11px] line-clamp-2 mb-4 flex-grow leading-relaxed">
        {{ book.description || 'Chưa có mô tả' }}
      </p>

      <button 
        @click="$emit('add-to-cart', book)"
        class="w-full bg-gray-900 hover:bg-blue-600 text-white py-2.5 rounded-lg transition-all font-semibold text-sm shadow-md hover:shadow-lg flex items-center justify-center space-x-2"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
        </svg>
        <span>Add To Cart</span>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Book } from '../../interfaces/book.interface'

defineProps<{ book: Book }>()
defineEmits(['add-to-cart'])

const formatPrice = (price: number) => {
  return new Intl.NumberFormat('vi-VN').format(price)
}
</script>
