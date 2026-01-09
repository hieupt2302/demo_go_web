<template>
  <header class="bg-white border-b border-gray-200 sticky top-0 z-50 shadow-sm">
    <div class="max-w-7xl mx-auto px-4 py-4">
      <div class="flex items-center justify-between">
        <!-- Logo -->
        <router-link to="/" class="text-2xl font-bold text-gray-900">
          Books
        </router-link>

        <!-- Navigation -->
        <nav class="hidden md:flex items-center space-x-8">
          <router-link to="/" class="text-gray-700 hover:text-blue-600 transition-colors font-medium">
            Trang chủ
          </router-link>
          <a href="#" class="text-gray-700 hover:text-blue-600 transition-colors font-medium">
            Về chúng tôi
          </a>
          <a href="#" class="text-gray-700 hover:text-blue-600 transition-colors font-medium">
            Liên hệ
          </a>
          <a href="#" class="text-gray-700 hover:text-blue-600 transition-colors font-medium">
            Blog
          </a>
        </nav>

        <!-- Right side actions -->
        <div class="flex items-center space-x-4">
          <router-link to="/login" class="text-gray-700 hover:text-blue-600 transition-colors font-medium">
            Đăng nhập
          </router-link>
          
          <!-- Cart Icon with Badge -->
          <router-link 
            to="/cart" 
            class="relative p-2 hover:bg-gray-100 rounded-full transition-colors"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-gray-700" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
            </svg>
            <span 
              v-if="cartStore.totalItems > 0" 
              class="absolute -top-1 -right-1 bg-red-500 text-white text-xs font-bold rounded-full h-5 w-5 flex items-center justify-center"
            >
              {{ cartStore.totalItems }}
            </span>
          </router-link>
        </div>
      </div>
    </div>
  </header>
</template>

<script setup lang="ts">
import { useCartStore } from '../../stores/cart'

const cartStore = useCartStore()
</script>
