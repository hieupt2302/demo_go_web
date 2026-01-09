import api from '../config/axios';
import type { Book } from '../interfaces/book.interface';

export const bookApi = {
  getAll: () => api.get<Book[]>('/book/'),
  getById: (id: number) => api.get<Book>(`/book/${id}`),
  create: (data: Partial<Book>) => api.post('/book/create', data),
  update: (id: number, data: Partial<Book>) => api.put(`/book/update/${id}`, data),
  delete: (id: number) => api.delete(`/book/delete/${id}`),
};
