<script setup lang="ts">
import AppHeader from './components/common/AppHeader.vue'
import AppFooter from './components/common/AppFooter.vue'
</script>

<template>
  <div class="app-container min-h-screen flex flex-col">
    <AppHeader />
    <main class="flex-grow">
      <router-view />
    </main>
    <AppFooter />
  </div>
</template>