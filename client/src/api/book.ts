import api from '../config/axios';
import type { Book } from '../interfaces/book.interface';
import type { ApiResponse } from '../interfaces/api.interface';

export const bookApi = {
  getAll: () => api.get<ApiResponse<Book[]>>(`/book/`),
  getById: (id: number) => api.get<ApiResponse<Book>>(`/book/${id}`),
  create: (data: Partial<Book>) => api.post<ApiResponse<Book>>(`/book/create`, data),
  update: (id: number, data: Partial<Book>) => api.put<ApiResponse<Book>>(`/book/update/${id}`, data),
  delete: (id: number) => api.delete<ApiResponse<null>>(`/book/delete/${id}`),
  search: (params: { type: string, q: string }) => api.get<ApiResponse<Book[]>>('/book/advanced-search', { params }),
};
