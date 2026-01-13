import api from '../config/axios';
import type { ApiResponse } from '../interfaces/api.interface';
import type { Category } from '../interfaces/book.interface';

export const categoryApi = {
    getAll: () => api.get<ApiResponse<Category[]>>(`/category/`),
    getById: (id: number) => api.get<ApiResponse<Category>>(`/category/${id}`),
    create: (data: Partial<Category>) => api.post<ApiResponse<Category>>(`/category/create`, data),
    update: (id: number, data: Partial<Category>) => api.put<ApiResponse<Category>>(`/category/update/${id}`, data),
    delete: (id: number) => api.delete<ApiResponse<null>>(`/category/delete/${id}`)
};