<template>
  <section>
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-2xl font-bold text-gray-800">Quản lý Kho Sách</h1>
        <p class="text-sm text-gray-500">Tổng cộng: {{ books.length }} cuốn sách</p>
      </div>
      <button @click="openAddModal" class="bg-blue-600 text-white px-5 py-2.5 rounded-xl hover:bg-blue-700 shadow-lg shadow-blue-200 transition-all flex items-center gap-2">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
        </svg>
        Thêm sách mới
      </button>
    </div>

    <div class="bg-white rounded-2xl shadow-sm border border-gray-100 overflow-hidden">
      <table class="w-full text-left border-collapse">
        <thead>
          <tr class="bg-gray-50/50 border-b border-gray-100">
            <th class="p-4 font-semibold text-gray-600">Ảnh</th>
            <th class="p-4 font-semibold text-gray-600">Thông tin sách</th>
            <th class="p-4 font-semibold text-gray-600">Giá bán</th>
            <th class="p-4 font-semibold text-gray-600">Tồn kho</th>
            <th class="p-4 font-semibold text-gray-600 text-right">Thao tác</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-50">
          <tr v-for="book in books" :key="book.id" class="hover:bg-blue-50/30 transition-colors">
            <td class="p-4">
              <img 
                :src="getImageUrl(book.cover_image_url)" 
                class="w-12 h-16 object-cover rounded-lg shadow-sm bg-gray-100" 
                @error="handleImageError"
              />
            </td>
            <td class="p-4">
              <div class="font-bold text-gray-900">{{ book.title }}</div>
              <div class="text-xs text-gray-400 mt-1">ISBN: {{ book.isbn || 'N/A' }}</div>
            </td>
            <td class="p-4 text-blue-600 font-bold">{{ formatPrice(book.price) }}đ</td>
            <td class="p-4">
              <span :class="['px-2 py-1 rounded-md text-xs font-medium', book.stock_quantity > 10 ? 'bg-green-100 text-green-700' : 'bg-red-100 text-red-700']">
                {{ book.stock_quantity }} cuốn
              </span>
            </td>
            <td class="p-4 text-right space-x-2">
              <button @click="editBook(book)" class="p-2 text-amber-600 hover:bg-amber-50 rounded-lg transition-colors font-medium">Sửa</button>
              <button @click="handleDelete(book.id)" class="p-2 text-red-600 hover:bg-red-50 rounded-lg transition-colors font-medium">Xóa</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-if="showModal" class="fixed inset-0 bg-black/50 backdrop-blur-sm flex items-center justify-center z-[100] p-4">
      <div class="bg-white rounded-3xl w-full max-w-2xl max-h-[90vh] overflow-y-auto p-8 shadow-2xl">
        <div class="flex justify-between items-center mb-6">
          <h2 class="text-2xl font-bold text-gray-800">{{ isEdit ? 'Cập nhật sách' : 'Tạo sách mới' }}</h2>
          <button @click="showModal = false" class="text-gray-400 hover:text-gray-600 transition-colors">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <form @submit.prevent="saveBook" class="space-y-5">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-5">
            <div class="space-y-1">
              <label class="text-sm font-semibold text-gray-700">Tiêu đề sách *</label>
              <input v-model="form.title" required class="w-full border border-gray-200 rounded-xl px-4 py-2.5 focus:border-blue-500 focus:ring-2 focus:ring-blue-200 outline-none transition-all" />
            </div>
            <div class="space-y-1">
              <label class="text-sm font-semibold text-gray-700">Mã ISBN</label>
              <input v-model="form.isbn" class="w-full border border-gray-200 rounded-xl px-4 py-2.5 focus:border-blue-500 focus:ring-2 focus:ring-blue-200 outline-none transition-all" placeholder="978-..." />
            </div>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-5">
            <div class="space-y-1">
              <label class="text-sm font-semibold text-gray-700">Giá bán (VNĐ) *</label>
              <input type="number" v-model="form.price" required class="w-full border border-gray-200 rounded-xl px-4 py-2.5 outline-none focus:border-blue-500" />
            </div>
            <div class="space-y-1">
              <label class="text-sm font-semibold text-gray-700">Số lượng tồn kho *</label>
              <input type="number" v-model="form.stock_quantity" required class="w-full border border-gray-200 rounded-xl px-4 py-2.5 outline-none focus:border-blue-500" />
            </div>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-5">
            <div class="space-y-1">
              <label class="text-sm font-semibold text-gray-700">Ngày xuất bản</label>
              <input type="date" v-model="form.publish_date" class="w-full border border-gray-200 rounded-xl px-4 py-2.5 outline-none focus:border-blue-500 focus:ring-2 focus:ring-blue-200 transition-all" />
            </div>
            <div class="space-y-1">
              <label class="text-sm font-semibold text-gray-700">Thể loại</label>
              <select v-model="form.category_id" class="w-full border border-gray-200 rounded-xl px-4 py-2.5 outline-none focus:border-blue-500">
                <option v-for="c in categories" :key="c.id" :value="c.id">{{ c.name }}</option>
              </select>
            </div>
          </div>

          <div class="space-y-1">
            <label class="text-sm font-semibold text-gray-700">Tác giả</label>
            <select v-model="form.author_id" class="w-full border border-gray-200 rounded-xl px-4 py-2.5 outline-none focus:border-blue-500">
              <option v-for="a in authors" :key="a.id" :value="a.id">{{ a.name }}</option>
            </select>
          </div>

          <div class="space-y-1">
            <label class="text-sm font-semibold text-gray-700">Ảnh bìa {{ isEdit ? '(Để trống nếu giữ ảnh cũ)' : '*' }}</label>
            <div class="mt-1 flex justify-center px-6 pt-5 pb-6 border-2 border-gray-300 border-dashed rounded-xl hover:border-blue-400 transition-colors">
              <div class="space-y-1 text-center">
                <svg class="mx-auto h-12 w-12 text-gray-400" stroke="currentColor" fill="none" viewBox="0 0 48 48">
                  <path d="M28 8H12a4 4 0 00-4 4v20m32-12v8m0 0v8a4 4 0 01-4 4H12a4 4 0 01-4-4v-4m32-4l-3.172-3.172a4 4 0 00-5.656 0L28 28M8 32l9.172-9.172a4 4 0 015.656 0L28 28m0 0l4 4m4-24h8m-4-4v8m-12 4h.02" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
                </svg>
                <div class="flex text-sm text-gray-600">
                  <input type="file" @change="onFileChange" accept="image/*" class="w-full text-sm text-gray-500 file:mr-4 file:py-2 file:px-4 file:rounded-full file:border-0 file:text-sm file:font-semibold file:bg-blue-50 file:text-blue-700 hover:file:bg-blue-100" />
                </div>
              </div>
            </div>
          </div>

          <div class="space-y-1">
            <label class="text-sm font-semibold text-gray-700">Mô tả sách</label>
            <textarea v-model="form.description" rows="3" class="w-full border border-gray-200 rounded-xl px-4 py-2.5 outline-none focus:border-blue-500" placeholder="Thông tin tóm tắt về sách..."></textarea>
          </div>

          <div class="flex justify-end gap-3 pt-4">
            <button type="button" @click="showModal = false" class="px-6 py-2.5 text-gray-500 font-medium hover:bg-gray-100 rounded-xl transition-all">Hủy</button>
            <button type="submit" :disabled="loading" class="px-8 py-2.5 bg-blue-600 text-white font-bold rounded-xl hover:bg-blue-700 disabled:bg-blue-300 shadow-lg shadow-blue-100 transition-all">
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
import { bookApi } from '../../api/book'
import { categoryApi } from '../../api/category'
import { authorApi } from '../../api/author'
import type { Book, Category, Author } from '../../interfaces/book.interface'

const books = ref<Book[]>([])
const categories = ref<Category[]>([])
const authors = ref<Author[]>([])
const loading = ref(false)
const showModal = ref(false)
const isEdit = ref(false)
const selectedFile = ref<File | null>(null)

const form = reactive({
  id: 0,
  title: '',
  isbn: '',
  price: 0,
  stock_quantity: 0,
  description: '',
  category_id: 0,
  author_id: 0,
  publish_date: '' //
})

const formatPrice = (p: number) => new Intl.NumberFormat('vi-VN').format(p)

// Hàm lấy URL ảnh chuẩn
const getImageUrl = (path: string) => {
  if (!path) return '/placeholder-book.jpg'
  if (path.startsWith('http')) return path
  // Fix dấu gạch chéo ngược trên Windows để hiển thị trên web
  return `${import.meta.env.VITE_API_BASE_URL}/${path}`.replace('/api/', '/').replace(/\\/g, '/') 
}

const loadInitialData = async () => {
  try {
    const [bRes, cRes, aRes] = await Promise.all([
      bookApi.getAll(),
      categoryApi.getAll(),
      authorApi.getAll()
    ])
    books.value = bRes.data.data
    categories.value = cRes.data.data
    authors.value = aRes.data.data
  } catch (err) {
    console.error("Lỗi tải dữ liệu:", err)
  }
}

const openAddModal = () => {
  isEdit.value = false
  selectedFile.value = null
  Object.assign(form, { 
    id: 0, 
    title: '', 
    isbn: '', 
    price: 0, 
    stock_quantity: 0, 
    description: '', 
    category_id: categories.value[0]?.id || 0, 
    author_id: authors.value[0]?.id || 0,
    publish_date: ''
  })
  showModal.value = true
}

const editBook = (book: any) => {
  isEdit.value = true
  selectedFile.value = null
  // Cắt chuỗi ISO từ Backend để hiển thị đúng trong <input type="date"> (YYYY-MM-DD)
  const formattedDate = book.publish_date ? book.publish_date.substring(0, 10) : ''
  
  Object.assign(form, {
    id: book.id,
    title: book.title,
    isbn: book.isbn,
    price: book.price,
    stock_quantity: book.stock_quantity,
    description: book.description,
    category_id: book.category_id,
    author_id: book.author_id,
    publish_date: formattedDate
  })
  showModal.value = true
}

const onFileChange = (e: any) => {
  const file = e.target.files[0]
  if (file) {
    selectedFile.value = file
    console.log("Log file selected:", file.name) //
  } 
}

const saveBook = async () => {
  loading.value = true
  try {
    const fd = new FormData() //
    
    // Đóng gói dữ liệu
    fd.append('title', form.title)
    fd.append('isbn', form.isbn)
    fd.append('price', form.price.toString())
    fd.append('stock_quantity', form.stock_quantity.toString())
    fd.append('description', form.description)
    fd.append('category_id', form.category_id.toString())
    fd.append('author_id', form.author_id.toString())
    fd.append('publish_date', form.publish_date) // Định dạng YYYY-MM-DD
    
    // Gửi file với key là 'file' để khớp với Middleware Backend
    if (selectedFile.value) {
      fd.append('file', selectedFile.value)
    }

    if (isEdit.value) {
      await bookApi.update(form.id, fd as any)
    } else {
      await bookApi.create(fd as any)
    }
    
    showModal.value = false
    await loadInitialData()
    alert("Cập nhật dữ liệu thành công!")
  } catch (err: any) {
    const msg = err.response?.data?.error || "Lỗi lưu dữ liệu"
    alert(msg)
  } finally {
    loading.value = false
  }
}

const handleDelete = async (id: number) => {
  if (confirm("Bạn có chắc chắn muốn xóa cuốn sách này không?")) {
    try {
      await bookApi.delete(id)
      loadInitialData()
    } catch (err) {
      alert("Không thể xóa sách này.")
    }
  }
}

const handleImageError = (e: Event) => {
  const target = e.target as HTMLImageElement;
  target.src = '/placeholder-book.jpg';
};

onMounted(loadInitialData)
</script>