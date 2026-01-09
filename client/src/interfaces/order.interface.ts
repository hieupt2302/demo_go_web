import type { Book } from "./book.interface";

export interface OrderItem {
  id?: number;
  order_id?: number;
  book_id: number;
  quantity: number;
  price_at_purchase: number;
  book?: Book;
}

export interface Order {
  id: number;
  user_id: number;
  order_date: string;
  total_amount: number;
  status: 'pending' | 'shipped' | 'delivered' | 'cancelled';
  shipping_address: string;
  order_items: OrderItem[];
}


export interface CreateOrderInput {
  user_id: number;
  shipping_address: string;
  order_items: {
    book_id: number;
    quantity: number;
  }[];
}