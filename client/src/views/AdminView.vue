<template>
  <div class="min-h-screen bg-gray-50 flex">
    <aside class="w-64 bg-gray-900 text-white p-6 shrink-0">
      <div class="flex items-center gap-3 mb-10">
        <div class="bg-blue-600 p-2 rounded-lg font-bold">BK</div>
        <h2 class="text-xl font-bold">Admin Console</h2>
      </div>
      
      <nav class="space-y-2">
        <button 
          v-for="tab in menuItems" :key="tab.id"
          @click="currentTab = tab.id"
          :class="['w-full flex items-center gap-3 px-4 py-3 rounded-xl transition-all', 
                  currentTab === tab.id ? 'bg-blue-600 text-white shadow-lg' : 'text-gray-400 hover:bg-gray-800']"
        >
          <span>{{ tab.icon }}</span>
          <span class="font-medium">{{ tab.label }}</span>
        </button>
      </nav>
    </aside>

    <main class="flex-1 p-8 overflow-y-auto">
      <div class="max-w-6xl mx-auto">
        <BookManager v-if="currentTab === 'books'" />
        <CategoryManager v-if="currentTab === 'categories'" />
        <AuthorManager v-if="currentTab === 'authors'" />
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import BookManager from '../components/admin/BookManager.vue'
import CategoryManager from '../components/admin/CategoryManager.vue'
import AuthorManager from '../components/admin/AuthorManager.vue'

const currentTab = ref('books')

const menuItems = [
  { id: 'books', label: 'Quản lý Sách', icon: '📚' },
  { id: 'categories', label: 'Thể loại', icon: '🏷️' },
  { id: 'authors', label: 'Tác giả', icon: '✍️' }
]
</script>