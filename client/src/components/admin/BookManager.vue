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
            <th class="p-4 font-semibold text-gray-600">Thể loại / Tác giả</th>
            <th class="p-4 font-semibold text-gray-600">Giá bán</th>
            <th class="p-4 font-semibold text-gray-600">Tồn kho</th>
            <th class="p-4 font-semibold text-gray-600 text-right">Thao tác</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-50">
          <tr v-for="book in books" :key="book.id" class="hover:bg-blue-50/30 transition-colors">
            <td class="p-4">
              <img :src="getImageUrl(book.cover_image_url)" class="w-12 h-16 object-cover rounded-lg shadow-sm bg-gray-100" @error="handleImageError" />
            </td>
            <td class="p-4">
              <div class="font-bold text-gray-900">{{ book.title }}</div>
              <div class="text-xs text-gray-400 mt-1">ISBN: {{ book.isbn || 'N/A' }}</div>
            </td>
            <td class="p-4 max-w-xs">
              <div class="flex flex-wrap gap-1 mb-1">
                <span v-for="c in book.categories" :key="c.id" class="px-1.5 py-0.5 bg-emerald-50 text-emerald-600 rounded text-[10px] border border-emerald-100">
                  {{ c.name }}
                </span>
              </div>
              <div class="flex flex-wrap gap-1">
                <span v-for="a in book.authors" :key="a.id" class="px-1.5 py-0.5 bg-indigo-50 text-indigo-600 rounded text-[10px] border border-indigo-100">
                  {{ a.name }}
                </span>
              </div>
            </td>
            <td class="p-4 text-blue-600 font-bold">{{ formatPrice(book.price) }}đ</td>
            <td class="p-4 text-sm">{{ book.stock_quantity }} cuốn</td>
            <td class="p-4 text-right space-x-2">
              <button @click="editBook(book)" class="text-amber-600 hover:underline">Sửa</button>
              <button @click="handleDelete(book.id)" class="text-red-600 hover:underline">Xóa</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-if="showModal" class="fixed inset-0 bg-gray-900/60 backdrop-blur-md flex items-center justify-center z-[100] p-4 sm:p-6">
      <div class="bg-white rounded-[2rem] w-full max-w-4xl max-h-[92vh] flex flex-col shadow-2xl overflow-hidden animate-in fade-in zoom-in duration-300">
        
        <div class="px-8 py-6 border-b border-gray-100 flex justify-between items-center bg-gray-50/50">
          <div>
            <h2 class="text-2xl font-bold text-gray-800">{{ isEdit ? 'Cập nhật thông tin sách' : 'Phát hành sách mới' }}</h2>
            <p class="text-sm text-gray-500">Vui lòng điền đầy đủ các thông tin bắt buộc (*)</p>
          </div>
          <button @click="showModal = false" class="p-2 hover:bg-white hover:shadow-sm rounded-full transition-all text-gray-400 hover:text-gray-600">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/></svg>
          </button>
        </div>

        <form @submit.prevent="saveBook" class="flex-1 overflow-y-auto p-8 space-y-8">
          
          <div class="space-y-4">
            <h3 class="text-sm font-bold text-blue-600 uppercase tracking-widest flex items-center">
              <span class="w-8 h-px bg-blue-200 mr-3"></span> Thông tin cơ bản
            </h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div class="group">
                <label class="block text-sm font-bold text-gray-700 mb-2 transition-colors group-focus-within:text-blue-600">Tiêu đề sách *</label>
                <input v-model="form.title" required placeholder="Ví dụ: Harry Potter và Hòn đá Phù thủy" 
                       class="w-full bg-gray-50 border-2 border-gray-100 rounded-2xl px-5 py-3 outline-none focus:border-blue-500 focus:bg-white transition-all shadow-sm" />
              </div>
              <div>
                <label class="block text-sm font-bold text-gray-700 mb-2">Mã ISBN</label>
                <input v-model="form.isbn" placeholder="978-3-16-148410-0" 
                       class="w-full bg-gray-50 border-2 border-gray-100 rounded-2xl px-5 py-3 outline-none focus:border-blue-500 focus:bg-white transition-all shadow-sm" />
              </div>
            </div>
          </div>

          <div class="space-y-4">
            <h3 class="text-sm font-bold text-blue-600 uppercase tracking-widest flex items-center">
              <span class="w-8 h-px bg-blue-200 mr-3"></span> Giá & Kho hàng
            </h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div class="relative">
                <label class="block text-sm font-bold text-gray-700 mb-2">Giá bán (VNĐ) *</label>
                <div class="relative">
                  <input type="number" v-model="form.price" required class="w-full bg-gray-50 border-2 border-gray-100 rounded-2xl pl-5 pr-12 py-3 outline-none focus:border-blue-500 focus:bg-white transition-all shadow-sm" />
                  <span class="absolute right-4 top-1/2 -translate-y-1/2 text-gray-400 font-medium">₫</span>
                </div>
              </div>
              <div>
                <label class="block text-sm font-bold text-gray-700 mb-2">Số lượng tồn kho *</label>
                <input type="number" v-model="form.stock_quantity" required class="w-full bg-gray-50 border-2 border-gray-100 rounded-2xl px-5 py-3 outline-none focus:border-blue-500 focus:bg-white transition-all shadow-sm" />
              </div>
            </div>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
            <div class="space-y-4">
            <h3 class="text-sm font-bold text-blue-600 uppercase tracking-widest flex items-center justify-between">
                <span class="flex items-center"><span class="w-8 h-px bg-blue-200 mr-3"></span> Thể loại</span>
                <span class="text-[10px] bg-blue-100 px-2 py-0.5 rounded text-blue-600">Đã chọn: {{ form.category_ids.length }}</span>
            </h3>
            
            <div class="relative group">
                <input 
                v-model="categorySearch" 
                placeholder="Tra cứu tên thể loại..." 
                class="w-full bg-white border-2 border-gray-100 rounded-xl pl-10 pr-4 py-2 text-sm outline-none focus:border-blue-400 transition-all shadow-sm"
                />
                <svg class="w-4 h-4 text-gray-400 absolute left-3 top-1/2 -translate-y-1/2" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/></svg>
            </div>

            <div class="bg-gray-50 border-2 border-gray-100 rounded-2xl p-2 max-h-48 overflow-y-auto space-y-1">
                <label v-for="c in filteredCategories()" :key="c.id" 
                    :class="['flex items-center p-2.5 rounded-xl cursor-pointer transition-all border-2', 
                                form.category_ids.includes(c.id) ? 'bg-white border-blue-500 shadow-sm' : 'bg-transparent border-transparent hover:bg-white/50']">
                <input type="checkbox" :value="c.id" v-model="form.category_ids" class="hidden" />
                <div :class="['w-4 h-4 rounded-md border-2 mr-3 flex items-center justify-center transition-all', 
                                form.category_ids.includes(c.id) ? 'bg-blue-600 border-blue-600' : 'bg-white border-gray-300']">
                    <svg v-if="form.category_ids.includes(c.id)" class="w-3 h-3 text-white" fill="currentColor" viewBox="0 0 20 20"><path d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"/></svg>
                </div>
                <span class="text-sm font-semibold" :class="form.category_ids.includes(c.id) ? 'text-blue-700' : 'text-gray-600'">{{ c.name }}</span>
                </label>
                <p v-if="filteredCategories().length === 0" class="text-center py-4 text-xs text-gray-400 italic">Không tìm thấy thể loại này</p>
            </div>
            </div>

            <div class="space-y-4">
            <h3 class="text-sm font-bold text-blue-600 uppercase tracking-widest flex items-center justify-between">
                <span class="flex items-center"><span class="w-8 h-px bg-blue-200 mr-3"></span> Tác giả</span>
                <span class="text-[10px] bg-blue-100 px-2 py-0.5 rounded text-blue-600">Đã chọn: {{ form.author_ids.length }}</span>
            </h3>

            <div class="relative group">
                <input  
                v-model="authorSearch" 
                placeholder="Tra cứu tên tác giả..." 
                class="w-full bg-white border-2 border-gray-100 rounded-xl pl-10 pr-4 py-2 text-sm outline-none focus:border-blue-400 transition-all shadow-sm"
                />
                <svg class="w-4 h-4 text-gray-400 absolute left-3 top-1/2 -translate-y-1/2" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/></svg>
            </div>

            <div class="bg-gray-50 border-2 border-gray-100 rounded-2xl p-2 max-h-48 overflow-y-auto space-y-1">
                <label v-for="a in filteredAuthors()" :key="a.id" 
                    :class="['flex items-center p-2.5 rounded-xl cursor-pointer transition-all border-2', 
                                form.author_ids.includes(a.id) ? 'bg-white border-blue-500 shadow-sm' : 'bg-transparent border-transparent hover:bg-white/50']">
                <input type="checkbox" :value="a.id" v-model="form.author_ids" class="hidden" />
                <div :class="['w-4 h-4 rounded-md border-2 mr-3 flex items-center justify-center transition-all', 
                                form.author_ids.includes(a.id) ? 'bg-blue-600 border-blue-600' : 'bg-white border-gray-300']">
                    <svg v-if="form.author_ids.includes(a.id)" class="w-3 h-3 text-white" fill="currentColor" viewBox="0 0 20 20"><path d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"/></svg>
                </div>
                <span class="text-sm font-semibold" :class="form.author_ids.includes(a.id) ? 'text-blue-700' : 'text-gray-600'">{{ a.name }}</span>
                </label>
                <p v-if="filteredAuthors().length === 0" class="text-center py-4 text-xs text-gray-400 italic">Không tìm thấy tác giả này</p>
            </div>
            </div>
        </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6 pt-4">
            <div class="space-y-6">
              <div>
                <label class="block text-sm font-bold text-gray-700 mb-2">Ngày xuất bản</label>
                <input type="date" v-model="form.publish_date" class="w-full bg-gray-50 border-2 border-gray-100 rounded-2xl px-5 py-3 outline-none focus:border-blue-500 shadow-sm" />
              </div>
              <div>
                <label class="block text-sm font-bold text-gray-700 mb-2">Nhà xuất bản</label>
                <input v-model="form.publisher" placeholder="Ví dụ: NXB Trẻ" class="w-full bg-gray-50 border-2 border-gray-100 rounded-2xl px-5 py-3 outline-none focus:border-blue-500 shadow-sm" />
              </div>
            </div>
            
            <div>
              <label class="block text-sm font-bold text-gray-700 mb-2">Ảnh bìa tác phẩm</label>
              <div class="relative group h-[148px]">
                <input type="file" @change="onFileChange" accept="image/*" 
                       class="absolute inset-0 w-full h-full opacity-0 cursor-pointer z-10" />
                <div class="h-full border-2 border-dashed border-gray-200 rounded-2xl flex flex-col items-center justify-center bg-gray-50 group-hover:bg-blue-50 group-hover:border-blue-300 transition-all">
                  <svg class="w-8 h-8 text-gray-400 group-hover:text-blue-500 mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"/></svg>
                  <span class="text-xs font-semibold text-gray-500">{{ selectedFile ? selectedFile.name : 'Kéo thả hoặc click để chọn ảnh' }}</span>
                </div>
              </div>
            </div>
          </div>

          <div class="space-y-2 pt-4">
            <label class="block text-sm font-bold text-gray-700">Tóm tắt nội dung</label>
            <textarea v-model="form.description" rows="4" 
                      placeholder="Viết một vài dòng giới thiệu về cuốn sách..."
                      class="w-full bg-gray-50 border-2 border-gray-100 rounded-2xl px-5 py-4 outline-none focus:border-blue-500 focus:bg-white transition-all shadow-sm resize-none"></textarea>
          </div>
        </form>

        <div class="px-8 py-6 bg-gray-50/50 border-t border-gray-100 flex justify-end gap-4">
          <button @click="showModal = false" class="px-8 py-3 text-gray-600 font-bold hover:bg-gray-200/50 rounded-2xl transition-all">Hủy bỏ</button>
          <button @click="saveBook" :disabled="loading" class="px-10 py-3 bg-blue-600 text-white font-bold rounded-2xl hover:bg-blue-700 shadow-lg shadow-blue-200 disabled:bg-blue-300 transition-all flex items-center">
            <svg v-if="loading" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path></svg>
            {{ loading ? 'Đang xử lý...' : (isEdit ? 'Lưu thay đổi' : 'Xác nhận thêm') }}
          </button>
        </div>
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
const categorySearch = ref('')
const authorSearch = ref('')

// Sửa lại form dùng category_ids và author_ids (Mảng)
const form = reactive({
  id: 0,
  title: '',
  isbn: '',
  price: 0,
  stock_quantity: 0,
  description: '',
  publisher: '',
  category_ids: [] as number[],
  author_ids: [] as number[],
  publish_date: '' 
})

const formatPrice = (p: number) => new Intl.NumberFormat('vi-VN').format(p)

const getImageUrl = (path: string) => {
  if (!path) return '/placeholder-book.jpg'
  if (path.startsWith('http')) return path
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
    id: 0, title: '', isbn: '', price: 0, stock_quantity: 0, 
    description: '', publisher: '', category_ids: [], author_ids: [], publish_date: ''
  })
  showModal.value = true
}

const editBook = (book: any) => {
  isEdit.value = true
  selectedFile.value = null
  const formattedDate = book.publish_date ? book.publish_date.substring(0, 10) : ''
  
  // Ánh xạ các mảng Object từ API về mảng ID cho form
  Object.assign(form, {
    id: book.id,
    title: book.title,
    isbn: book.isbn,
    price: book.price,
    stock_quantity: book.stock_quantity,
    description: book.description,
    publisher: book.publisher,
    category_ids: book.categories ? book.categories.map((c: any) => c.id) : [],
    author_ids: book.authors ? book.authors.map((a: any) => a.id) : [],
    publish_date: formattedDate
  })
  showModal.value = true
}

const onFileChange = (e: any) => {
  const file = e.target.files[0]
  if (file) selectedFile.value = file
}

const saveBook = async () => {
  loading.value = true
  try {
    const fd = new FormData()
    fd.append('title', form.title)
    fd.append('isbn', form.isbn)
    fd.append('price', form.price.toString())
    fd.append('stock_quantity', form.stock_quantity.toString())
    fd.append('description', form.description)
    fd.append('publisher', form.publisher)
    fd.append('publish_date', form.publish_date)
    
    // GỬI MẢNG ID: Lặp qua từng phần tử và append cùng một key
    form.category_ids.forEach(id => fd.append('category_ids', id.toString()))
    form.author_ids.forEach(id => fd.append('author_ids', id.toString()))

    if (selectedFile.value) fd.append('file', selectedFile.value)

    if (isEdit.value) {
      await bookApi.update(form.id, fd as any)
    } else {
      await bookApi.create(fd as any)
    }
    
    showModal.value = false
    await loadInitialData()
    alert("Thành công!")
  } catch (err: any) {
    alert(err.response?.data?.error || "Lỗi lưu dữ liệu")
  } finally {
    loading.value = false
  }
}

const handleDelete = async (id: number) => {
  if (confirm("Xóa cuốn sách này?")) {
    try {
      await bookApi.delete(id)
      loadInitialData()
    } catch (err) {
      alert("Lỗi xóa")
    }
  }
}

const handleImageError = (e: Event) => {
  (e.target as HTMLImageElement).src = '/placeholder-book.jpg';
};
const filteredCategories = () => {
  return categories.value.filter(c => 
    c.name.toLowerCase().includes(categorySearch.value.toLowerCase())
  )
}

const filteredAuthors = () => {
  return authors.value.filter(a => 
    a.name.toLowerCase().includes(authorSearch.value.toLowerCase())
  )
}
onMounted(loadInitialData)
</script>