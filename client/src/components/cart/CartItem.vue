<template>
  <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-4 flex items-center space-x-4 hover:shadow-md transition-shadow">
    <!-- Book Image -->
    <div class="flex-shrink-0 w-24 h-32 bg-gray-100 rounded-lg overflow-hidden">
      <router-link :to="`/book/${item.book.id}`">
        <img 
          :src="getImageUrl(item.book.cover_image_url) || 'https://images.unsplash.com/photo-1544947950-fa07a98d237f?w=150&h=200&fit=crop'" 
          class="w-full h-full object-cover hover:scale-110 transition-transform duration-300"
          alt="book cover"
        />
      </router-link>
    </div>

    <!-- Book Info -->
    <div class="flex-grow">
      <router-link :to="`/book/${item.book.id}`" class="hover:text-blue-600">
        <h3 class="font-bold text-gray-900 mb-1 line-clamp-1">{{ item.book.title }}</h3>
      </router-link>
      <p class="text-sm text-gray-500 mb-2">{{ item.book.authors?.map(a => a.name).join(', ') || 'Ẩn danh' }}</p>
      <p class="text-lg font-bold text-blue-600">{{ formatPrice(item.book.price) }}đ</p>
    </div>

    <!-- Quantity Controls -->
    <div class="flex items-center space-x-3">
      <button 
        @click="$emit('decrement', item.book.id)"
        class="w-8 h-8 rounded-full bg-gray-100 hover:bg-gray-200 flex items-center justify-center transition-colors"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 12H4" />
        </svg>
      </button>
      
      <span class="font-semibold text-gray-900 w-8 text-center">{{ item.quantity }}</span>
      
      <button 
        @click="$emit('increment', item.book.id)"
        class="w-8 h-8 rounded-full bg-gray-100 hover:bg-gray-200 flex items-center justify-center transition-colors"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
        </svg>
      </button>
    </div>

    <!-- Subtotal -->
    <div class="text-right">
      <p class="text-sm text-gray-500 mb-1">Subtotal</p>
      <p class="text-lg font-bold text-gray-900">{{ formatPrice(item.book.price * item.quantity) }}đ</p>
    </div>

    <!-- Remove Button -->
    <button 
      @click="$emit('remove', item.book.id)"
      class="p-2 text-red-500 hover:bg-red-50 rounded-lg transition-colors"
    >
      <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
      </svg>
    </button>
  </div>
</template>

<script setup lang="ts">
import type { CartItem } from '../../interfaces/cartItem.interfaces'

defineProps<{ item: CartItem }>()
defineEmits(['increment', 'decrement', 'remove'])

const formatPrice = (price: number) => {
  return new Intl.NumberFormat('vi-VN').format(price)
}
const getImageUrl = (path: string) => {
  if (!path) return '/placeholder-book.jpg'
  if (path.startsWith('http')) return path
  return `${import.meta.env.VITE_API_BASE_URL}/${path}`.replace('/api/', '/').replace(/\\/g, '/')
}
</script>
