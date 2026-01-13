<template>
  <div class="bg-white rounded-xl shadow-sm border border-gray-100 overflow-hidden hover:shadow-lg transition-all flex flex-col h-full group relative">
    
    <router-link :to="`/book/${book.id}`" class="relative h-64 bg-gray-100 overflow-hidden block">
      <img 
        :src="getImageUrl(book.cover_image_url) || 'https://images.unsplash.com/photo-1544947950-fa07a98d237f?w=300&h=400&fit=crop'" 
        class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500"
        alt="book cover"
      />
      <div v-if="book.categories?.length" class="absolute top-2 left-2 flex flex-wrap gap-1 max-w-[80%]">
        <span 
          v-for="cat in book.categories.slice(0, 2)" :key="cat.id"
          class="bg-blue-600/90 backdrop-blur-sm text-white text-[9px] px-2 py-0.5 rounded shadow-sm uppercase font-bold"
        >
          {{ cat.name }}
        </span>
      </div>
    </router-link>

    <div class="p-4 flex flex-col flex-grow">
      <div class="flex items-center justify-between mb-2">
        <span class="text-xl font-black text-blue-600">{{ formatPrice(book.price) }}đ</span>
      </div>

      <router-link :to="`/book/${book.id}`" class="block group-hover:text-blue-600 transition-colors">
        <h3 class="font-bold text-gray-900 mb-1 line-clamp-2 leading-tight text-sm h-10 italic group-hover:not-italic">
          {{ book.title }}
        </h3>
      </router-link>

      <p class="text-[11px] text-gray-500 mb-2 truncate">
        <span class="font-medium text-gray-400">Tác giả:</span> 
        {{ book.authors?.map(a => a.name).join(', ') || 'Đang cập nhật' }}
      </p>

      <p class="text-gray-400 text-[11px] line-clamp-2 mb-4 flex-grow leading-relaxed">
        {{ book.description || 'Chưa có mô tả...' }}
      </p>

      <button 
        @click.stop="$emit('add-to-cart', book)"
        class="w-full bg-gray-900 hover:bg-blue-600 text-white py-2.5 rounded-lg transition-all font-semibold text-sm shadow-md hover:shadow-lg flex items-center justify-center space-x-2 active:scale-95 z-10"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
        </svg>
        <span>Thêm vào giỏ</span>
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

const getImageUrl = (path: string) => {
  if (!path) return '/placeholder-book.jpg'
  if (path.startsWith('http')) return path
  return `${import.meta.env.VITE_API_BASE_URL}/${path}`.replace('/api/', '/').replace(/\\/g, '/')
}
</script>