import api from '../config/axios';
import type { Order, CreateOrderInput, UpdateOrderInput } from '../interfaces/order.interface';
import type { ApiResponse } from '../interfaces/api.interface';

export const orderApi = {
  create: (data: CreateOrderInput) => api.post<ApiResponse<Order>>('/order/create', data),
  getAll: () => api.get<ApiResponse<Order[]>>('/order/'),
  getById: (id: number) => api.get<ApiResponse<Order>>(`/order/${id}`),
  getByUser: (id: number) => api.get<ApiResponse<Order[]>>(`/order/user/${id}`),
  update: (id: number, data: UpdateOrderInput) => api.put<ApiResponse<Order>>(`/order/${id}`, data),
  delete: (id: number) => api.delete<ApiResponse<null>>(`/order/${id}`),
};
