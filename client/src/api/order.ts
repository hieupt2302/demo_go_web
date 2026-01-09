import api from '../config/axios';
import type { Order, CreateOrderInput } from '../interfaces/order.interface';

export const orderApi = {
  create: (data: CreateOrderInput) => api.post('/order/create', data),
  getAll: () => api.get<Order[]>('/order/'),
  getById: (id: number) => api.get<Order>(`/order/${id}`),
};