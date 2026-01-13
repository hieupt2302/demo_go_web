import api from '../config/axios';
import type { ApiResponse } from '../interfaces/api.interface';
import type { Author } from '../interfaces/book.interface';

export const authorApi = {
    getAll: () => api.get<ApiResponse<Author[]>>(`/author/`),
    getById: (id: number) => api.get<ApiResponse<Author>>(`/author/${id}`),
    create: (data: Partial<Author>) => api.post<ApiResponse<Author>>(`/author/create`, data),
    update: (id: number, data: Partial<Author>) => api.put<ApiResponse<Author>>(`/author/update/${id}`, data),
    delete: (id: number) => api.delete<ApiResponse<null>>(`/author/delete/${id}`)
};
