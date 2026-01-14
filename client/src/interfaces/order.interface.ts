import type { Book } from "./book.interface";

export type OrderStatus = 'pending' | 'paid' | 'failed' | 'shipped' | 'delivered' | 'cancelled';
export type PaymentMethod = 'VNPAY' | 'MOMO' | 'BANK_TRANSFER' | 'COD';

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
  status: OrderStatus;
  shipping_address: string;
  payment_method: PaymentMethod;
  payment_status: 'unpaid' | 'paid' | 'refunded';
  txn_id: string;
  order_items: OrderItem[];
}


export interface CreateOrderInput {
  // user_id: number;
  // total_amount: number;
  // status?: string;
  shipping_address: string;
  payment_method: PaymentMethod;
  items: {
    book_id: number;
    quantity: number;
  }[];
}

export interface UpdateOrderInput {
  status?: OrderStatus;
  shipping_address?: string;
}