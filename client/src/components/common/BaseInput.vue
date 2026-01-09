<template>
  <div class="mb-4">
    <label v-if="label" class="block text-gray-700 text-sm font-bold mb-2">{{ label }}</label>
    <input
      :type="type"
      :value="modelValue"
      @input="$emit('update:modelValue', ($event.target as HTMLInputElement).value)"
      class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:ring-2 focus:ring-blue-500"
      :placeholder="placeholder"
    />
  </div>
</template>

<script setup lang="ts">
defineProps(['label', 'modelValue', 'type', 'placeholder'])
defineEmits(['update:modelValue'])
</script>