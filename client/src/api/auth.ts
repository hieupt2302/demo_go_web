import api from '../config/axios';
import type { ApiResponse } from '../interfaces/api.interface';
import type { User } from '../interfaces/auth.interface';


export interface LoginData {
  access_token: string;
  refresh_token: string;
  user?: User;
}

export const authApi = {
  login: (data: any) => api.post<ApiResponse<LoginData>>('/auth/login', data),
  
  register: (data: any) => api.post<ApiResponse<null>>('/auth/register', data),
  
  refresh: (rt: string) => api.post<ApiResponse<LoginData>>('/auth/refresh', { refresh_token: rt })

};