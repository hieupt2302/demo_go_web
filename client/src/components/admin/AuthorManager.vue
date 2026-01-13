<template>
  <section>
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-800">Quản lý Tác giả</h1>
        <p class="text-sm text-gray-500">Tổng cộng: {{ authors.length }} tác giả</p>
      </div>
      <button @click="openAddModal" class="bg-indigo-600 text-white px-5 py-2.5 rounded-xl hover:bg-indigo-700 shadow-lg shadow-indigo-200 transition-all flex items-center gap-2">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
        </svg>
        Thêm tác giả
      </button>
    </div>

    <div class="bg-white rounded-2xl shadow-sm border border-gray-100 overflow-hidden">
      <table class="w-full text-left border-collapse">
        <thead>
          <tr class="bg-gray-50/50 border-b border-gray-100">
            <th class="p-4 font-semibold text-gray-600">ID</th>
            <th class="p-4 font-semibold text-gray-600">Tên tác giả</th>
            <th class="p-4 font-semibold text-gray-600">Tiểu sử</th>
            <th class="p-4 font-semibold text-gray-600 text-right">Thao tác</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-50">
          <tr v-for="author in authors" :key="author.id" class="hover:bg-indigo-50/30 transition-colors">
            <td class="p-4 text-gray-400">#{{ author.id }}</td>
            <td class="p-4 font-bold text-gray-900">{{ author.name }}</td>
            <td class="p-4 text-gray-500 text-sm italic max-w-xs truncate">
              {{ author.bio || 'Chưa có thông tin tiểu sử' }}
            </td>
            <td class="p-4 text-right space-x-2">
              <button @click="editAuthor(author)" class="p-2 text-amber-600 hover:bg-amber-50 rounded-lg transition-colors font-medium">Sửa</button>
              <button @click="handleDelete(author.id)" class="p-2 text-red-600 hover:bg-red-50 rounded-lg transition-colors font-medium">Xóa</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-if="showModal" class="fixed inset-0 bg-black/50 backdrop-blur-sm flex items-center justify-center z-[100] p-4">
      <div class="bg-white rounded-3xl w-full max-w-md p-8 shadow-2xl">
        <div class="flex justify-between items-center mb-6">
          <h2 class="text-2xl font-bold text-gray-800">{{ isEdit ? 'Cập nhật tác giả' : 'Thêm tác giả mới' }}</h2>
          <button @click="showModal = false" class="text-gray-400 hover:text-gray-600">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <form @submit.prevent="saveAuthor" class="space-y-5">
          <div class="space-y-1">
            <label class="text-sm font-semibold text-gray-700">Họ và tên *</label>
            <input v-model="form.name" required class="w-full border border-gray-200 rounded-xl px-4 py-2.5 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200 outline-none transition-all" placeholder="Nhập tên tác giả..." />
          </div>

          <div class="space-y-1">
            <label class="text-sm font-semibold text-gray-700">Tiểu sử</label>
            <textarea v-model="form.bio" rows="4" class="w-full border border-gray-200 rounded-xl px-4 py-2.5 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200 outline-none transition-all" placeholder="Thông tin tóm tắt về tác giả..."></textarea>
          </div>

          <div class="flex justify-end gap-3 pt-4">
            <button type="button" @click="showModal = false" class="px-6 py-2.5 text-gray-500 font-medium hover:bg-gray-100 rounded-xl transition-all">Hủy</button>
            <button type="submit" :disabled="loading" class="px-8 py-2.5 bg-indigo-600 text-white font-bold rounded-xl hover:bg-indigo-700 disabled:bg-indigo-300 shadow-lg shadow-indigo-100 transition-all">
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
import { authorApi } from '../../api/author'
import type { Author } from '../../interfaces/book.interface'

const authors = ref<Author[]>([])
const loading = ref(false)
const showModal = ref(false)
const isEdit = ref(false)

const form = reactive({
  id: 0,
  name: '',
  bio: ''
})

const loadAuthors = async () => {
  try {
    const res = await authorApi.getAll()
    authors.value = res.data.data
  } catch (err) {
    console.error("Lỗi tải danh sách tác giả:", err)
  }
}

const openAddModal = () => {
  isEdit.value = false
  Object.assign(form, { id: 0, name: '', bio: '' })
  showModal.value = true
}

const editAuthor = (author: Author) => {
  isEdit.value = true
  Object.assign(form, {
    id: author.id,
    name: author.name,
    bio: author.bio || ''
  })
  showModal.value = true
}

const saveAuthor = async () => {
  loading.value = true
  try {
    if (isEdit.value) {
      await authorApi.update(form.id, { name: form.name, bio: form.bio })
    } else {
      await authorApi.create({ name: form.name, bio: form.bio })
    }
    showModal.value = false
    await loadAuthors()
    alert("Thao tác thành công!")
  } catch (err: any) {
    alert(err.response?.data?.error || "Lỗi xử lý dữ liệu")
  } finally {
    loading.value = false
  }
}

const handleDelete = async (id: number) => {
  if (confirm("Bạn có chắc chắn muốn xóa tác giả này không?")) {
    try {
      await authorApi.delete(id)
      await loadAuthors()
    } catch (err) {
      alert("Không thể xóa tác giả này.")
    }
  }
}

onMounted(loadAuthors)
</script>