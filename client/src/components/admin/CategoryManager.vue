<template>
  <section>
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-800">Quản lý Thể loại</h1>
        <p class="text-sm text-gray-500">Tổng cộng: {{ categories.length }} thể loại</p>
      </div>
      <button @click="openAddModal" class="bg-emerald-600 text-white px-5 py-2.5 rounded-xl hover:bg-emerald-700 shadow-lg shadow-emerald-200 transition-all flex items-center gap-2">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
        </svg>
        Thêm thể loại
      </button>
    </div>

    <div class="bg-white rounded-2xl shadow-sm border border-gray-100 overflow-hidden">
      <table class="w-full text-left border-collapse">
        <thead>
          <tr class="bg-gray-50/50 border-b border-gray-100">
            <th class="p-4 font-semibold text-gray-600 w-16">ID</th>
            <th class="p-4 font-semibold text-gray-600">Tên thể loại</th>
            <th class="p-4 font-semibold text-gray-600">Mô tả</th>
            <th class="p-4 font-semibold text-gray-600 text-right">Thao tác</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-50">
          <tr v-for="cat in categories" :key="cat.id" class="hover:bg-emerald-50/30 transition-colors">
            <td class="p-4 text-gray-500">#{{ cat.id }}</td>
            <td class="p-4 font-bold text-gray-900">{{ cat.name }}</td>
            <td class="p-4 text-gray-500 text-sm">{{ cat.description || 'Không có mô tả' }}</td>
            <td class="p-4 text-right space-x-2">
              <button @click="editCategory(cat)" class="p-2 text-amber-600 hover:bg-amber-50 rounded-lg transition-colors font-medium">Sửa</button>
              <button @click="handleDelete(cat.id)" class="p-2 text-red-600 hover:bg-red-50 rounded-lg transition-colors font-medium">Xóa</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-if="showModal" class="fixed inset-0 bg-black/50 backdrop-blur-sm flex items-center justify-center z-[100] p-4">
      <div class="bg-white rounded-3xl w-full max-w-md p-8 shadow-2xl">
        <div class="flex justify-between items-center mb-6">
          <h2 class="text-2xl font-bold text-gray-800">{{ isEdit ? 'Cập nhật thể loại' : 'Thêm thể loại mới' }}</h2>
          <button @click="showModal = false" class="text-gray-400 hover:text-gray-600">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <form @submit.prevent="saveCategory" class="space-y-5">
          <div class="space-y-1">
            <label class="text-sm font-semibold text-gray-700">Tên thể loại *</label>
            <input 
              v-model="form.name" 
              required 
              class="w-full border border-gray-200 rounded-xl px-4 py-2.5 focus:border-emerald-500 focus:ring-2 focus:ring-emerald-200 outline-none transition-all"
              placeholder="Ví dụ: Khoa học viễn tưởng"
            />
          </div>

          <div class="space-y-1">
            <label class="text-sm font-semibold text-gray-700">Mô tả</label>
            <textarea 
              v-model="form.description" 
              rows="4" 
              class="w-full border border-gray-200 rounded-xl px-4 py-2.5 focus:border-emerald-500 focus:ring-2 focus:ring-emerald-200 outline-none transition-all"
              placeholder="Nhập mô tả ngắn gọn..."
            ></textarea>
          </div>

          <div class="flex justify-end gap-3 pt-4">
            <button type="button" @click="showModal = false" class="px-6 py-2.5 text-gray-500 font-medium hover:bg-gray-100 rounded-xl transition-all">Hủy</button>
            <button 
              type="submit" 
              :disabled="loading" 
              class="px-8 py-2.5 bg-emerald-600 text-white font-bold rounded-xl hover:bg-emerald-700 disabled:bg-emerald-300 shadow-lg shadow-emerald-100 transition-all"
            >
              <span v-if="loading">Đang xử lý...</span>
              <span v-else>{{ isEdit ? 'Cập nhật' : 'Thêm mới' }}</span>
            </button>
          </div>
        </form>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { categoryApi } from '../../api/category'
import type { Category } from '../../interfaces/book.interface'

const categories = ref<Category[]>([])
const loading = ref(false)
const showModal = ref(false)
const isEdit = ref(false)

const form = reactive({
  id: 0,
  name: '',
  description: ''
})

const loadCategories = async () => {
  try {
    const res = await categoryApi.getAll()
    categories.value = res.data.data
    console.log(categories.value)
  } catch (err) {
    console.error("Lỗi tải thể loại:", err)
  }
}

const openAddModal = () => {
  isEdit.value = false
  Object.assign(form, { id: 0, name: '', description: '' })
  showModal.value = true
}

const editCategory = (cat: Category) => {
  isEdit.value = true
  Object.assign(form, {
    id: cat.id,
    name: cat.name,
    description: cat.description
  })
  showModal.value = true
}

const saveCategory = async () => {
  loading.value = true
  try {
    if (isEdit.value) {
      await categoryApi.update(form.id, { name: form.name, description: form.description })
    } else {
      await categoryApi.create({ name: form.name, description: form.description })
    }
    showModal.value = false
    await loadCategories()
    alert("Thao tác thành công!")
  } catch (err: any) {
    alert(err.response?.data?.error || "Lỗi xử lý")
  } finally {
    loading.value = false
  }
}

const handleDelete = async (id: number) => {
  if (confirm("Xóa thể loại này có thể ảnh hưởng đến các sách thuộc thể loại này. Bạn chắc chứ?")) {
    try {
      await categoryApi.delete(id)
      await loadCategories()
    } catch (err) {
      alert("Không thể xóa thể loại này.")
    }
  }
}

onMounted(loadCategories)
</script>