<template>
  <section class="py-16 bg-gradient-to-r from-yellow-100 to-yellow-200">
    <div class="max-w-7xl mx-auto px-4">
      <div class="flex items-center justify-between">
        <div class="flex-1">
          <h2 class="text-4xl font-bold text-gray-900 mb-4">
            2025 National Book Awards for Fiction Shortlist
          </h2>
          <button class="bg-blue-600 hover:bg-blue-700 text-white px-6 py-2 rounded-lg font-semibold transition-colors">
            Learn More
          </button>
        </div>
        
        <!-- Trophy Illustration -->
        <div class="hidden md:block">
          <div class="relative">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-48 w-48 text-yellow-500" fill="currentColor" viewBox="0 0 24 24">
              <path d="M12 2L9 9H2l6 4.5L5 22l7-5 7 5-3-8.5L22 9h-7z"/>
            </svg>
            <div class="absolute top-0 right-0 w-32 h-32">
              <img src="https://images.unsplash.com/photo-1544947950-fa07a98d237f?w=150&h=200&fit=crop" 
                   class="rounded-lg shadow-xl transform rotate-12" 
                   alt="Award Book" />
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
</script>
