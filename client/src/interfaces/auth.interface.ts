export interface User {
  id: number;
  fullname: string;
  email: string;
  phone?: string;
  address?: string;
  role: 'customer' | 'admin';
  created_at?: string;
}

