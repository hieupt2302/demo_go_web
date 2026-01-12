import api from '../config/axios';
import type { ApiResponse } from '../interfaces/api.interface';
import type { User } from '../interfaces/auth.interface';

export const userApi = {
    getUserById: (id: number) => api.get<ApiResponse<User>>(`/user/${id}`),

    updateUser: (id: number, data: Partial<User>) => api.put<ApiResponse<User>>(`/user/update/${id}`, data),

    deleteUser: (id: number) => api.delete<ApiResponse<null>>(`/user/delete/${id}`)
};
