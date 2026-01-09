import type { Book } from "./book.interface"

export interface CartItem {
  book: Book
  quantity: number
}