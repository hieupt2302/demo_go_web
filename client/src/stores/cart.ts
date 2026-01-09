import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { Book } from '../interfaces/book.interface'
import type { CartItem } from '../interfaces/cartItem.interfaces'



export const useCartStore = defineStore('cart', () => {
  // --- STATE ---
  const items = ref<CartItem[]>([])

  // --- GETTERS ---
  // Tính tổng số lượng sản phẩm (để hiển thị trên icon giỏ hàng ở Navbar)
  const totalItems = computed(() =>
    items.value.reduce((total, item) => total + item.quantity, 0)
  )

  // Tính tổng tiền của cả giỏ hàng
  const totalPrice = computed(() =>
    items.value.reduce((total, item) => total + (item.book.price * item.quantity), 0)
  )

  // --- ACTIONS ---
  // Hàm thêm sách vào giỏ
  function addToCart(book: Book) {
    const existingItem = items.value.find(item => item.book.id === book.id)

    if (existingItem) {
      // Nếu sách đã có, chỉ tăng số lượng
      existingItem.quantity++
    } else {
      // Nếu chưa có, thêm mới vào mảng
      items.value.push({ book, quantity: 1 })
    }
  }

  // Hàm xóa sản phẩm khỏi giỏ
  function removeFromCart(bookId: number) {
    items.value = items.value.filter(item => item.book.id !== bookId)
  }

  // Hàm cập nhật số lượng sản phẩm
  function updateQuantity(bookId: number, quantity: number) {
    const item = items.value.find(item => item.book.id === bookId)
    if (item && quantity > 0) {
      item.quantity = quantity
    } else if (item && quantity <= 0) {
      removeFromCart(bookId)
    }
  }

  // Hàm tăng số lượng
  function incrementQuantity(bookId: number) {
    const item = items.value.find(item => item.book.id === bookId)
    if (item) {
      item.quantity++
    }
  }

  // Hàm giảm số lượng
  function decrementQuantity(bookId: number) {
    const item = items.value.find(item => item.book.id === bookId)
    if (item) {
      if (item.quantity > 1) {
        item.quantity--
      } else {
        removeFromCart(bookId)
      }
    }
  }

  // Hàm xóa sạch giỏ hàng (dùng sau khi thanh toán thành công)
  function clearCart() {
    items.value = []
  }

  return {
    items,
    totalItems,
    totalPrice,
    addToCart,
    removeFromCart,
    updateQuantity,
    incrementQuantity,
    decrementQuantity,
    clearCart
  }
}, {
  // Lưu giỏ hàng vào LocalStorage để F5 không bị mất
  persist: true
})