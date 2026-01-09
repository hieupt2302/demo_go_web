<template>
  <section class="bg-gradient-to-br from-blue-50 via-white to-purple-50 py-16">
    <div class="max-w-7xl mx-auto px-4">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-12 items-center">
        <!-- Left Content -->
        <div>
          <h1 class="text-5xl md:text-6xl font-extrabold text-gray-900 leading-tight mb-6">
            Buy and sell your books 
            <span class="text-blue-600">for the best prices</span>
          </h1>
          <p class="text-gray-600 text-lg mb-8 leading-relaxed">
            Find and read more books you'll love, and keep track of the books you want to read. 
            Be part of the world's largest community of book lovers on Goodreads.
          </p>
          <div class="flex space-x-4">
            <button class="bg-blue-600 hover:bg-blue-700 text-white px-8 py-3 rounded-lg font-semibold transition-colors shadow-lg">
              Explore Now
            </button>
          </div>
        </div>

        <!-- Right Image -->
        <div class="relative">
          <div class="relative z-10 transform hover:scale-105 transition-transform duration-300">
            <img 
              src="https://images.unsplash.com/photo-1544947950-fa07a98d237f?w=400&h=600&fit=crop" 
              alt="Featured Book"
              class="rounded-2xl shadow-2xl w-full max-w-md mx-auto"
            />
          </div>
          <!-- Decorative elements -->
          <div class="absolute top-10 right-10 w-72 h-72 bg-blue-200 rounded-full opacity-20 blur-3xl"></div>
          <div class="absolute bottom-10 left-10 w-64 h-64 bg-purple-200 rounded-full opacity-20 blur-3xl"></div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
</script>
