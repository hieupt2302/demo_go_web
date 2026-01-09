<template>
  <section class="py-16 bg-white">
    <div class="max-w-7xl mx-auto px-4">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-12 items-center">
        <!-- Left: Book Grid Image -->
        <div class="grid grid-cols-3 gap-4">
          <img src="https://images.unsplash.com/photo-1544947950-fa07a98d237f?w=200&h=300&fit=crop" class="rounded-lg shadow-md" alt="Book 1" />
          <img src="https://images.unsplash.com/photo-1512820790803-83ca734da794?w=200&h=300&fit=crop" class="rounded-lg shadow-md" alt="Book 2" />
          <img src="https://images.unsplash.com/photo-1495446815901-a7297e633e8d?w=200&h=300&fit=crop" class="rounded-lg shadow-md" alt="Book 3" />
          <img src="https://images.unsplash.com/photo-1524995997946-a1c2e315a42f?w=200&h=300&fit=crop" class="rounded-lg shadow-md" alt="Book 4" />
          <img src="https://images.unsplash.com/photo-1543002588-bfa74002ed7e?w=200&h=300&fit=crop" class="rounded-lg shadow-md" alt="Book 5" />
          <img src="https://images.unsplash.com/photo-1532012197267-da84d127e765?w=200&h=300&fit=crop" class="rounded-lg shadow-md" alt="Book 6" />
          <img src="https://images.unsplash.com/photo-1481627834876-b7833e8f5570?w=200&h=300&fit=crop" class="rounded-lg shadow-md" alt="Book 7" />
          <img src="https://images.unsplash.com/photo-1519682337058-a94d519337bc?w=200&h=300&fit=crop" class="rounded-lg shadow-md" alt="Book 8" />
          <img src="https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=200&h=300&fit=crop" class="rounded-lg shadow-md" alt="Book 9" />
        </div>

        <!-- Right: Content -->
        <div>
          <h2 class="text-4xl font-bold text-gray-900 mb-4">
            Find Your Favorite <span class="text-blue-600">Book Here!</span>
          </h2>
          <p class="text-gray-600 mb-8 leading-relaxed">
            Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed eu feugiat amet, libero ipsum enim pharetra hac. 
            Urna commodo, lacus ut magna velit eleifend. Amet, quis urna, a eu.
          </p>

          <!-- Stats -->
          <div class="grid grid-cols-3 gap-6 mb-8">
            <div>
              <h3 class="text-3xl font-bold text-gray-900">800+</h3>
              <p class="text-gray-500 text-sm">Book Listing</p>
            </div>
            <div>
              <h3 class="text-3xl font-bold text-gray-900">550+</h3>
              <p class="text-gray-500 text-sm">Register User</p>
            </div>
            <div>
              <h3 class="text-3xl font-bold text-gray-900">1,200+</h3>
              <p class="text-gray-500 text-sm">Books Selling</p>
            </div>
          </div>

          <button class="bg-blue-600 hover:bg-blue-700 text-white px-8 py-3 rounded-lg font-semibold transition-colors shadow-lg">
            Explore Now
          </button>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
</script>
