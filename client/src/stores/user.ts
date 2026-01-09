import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import type { User } from '../interfaces/auth.interface';

export const useAuthStore = defineStore('auth', () => {
  // --- STATE (Dữ liệu) ---
  const accessToken = ref<string | null>(null);
  const refreshToken = ref<string | null>(null);
  const user = ref<User | null>(null);

  // --- GETTERS (Dữ liệu tính toán) ---
  const isAuthenticated = computed(() => !!accessToken.value);
  const isAdmin = computed(() => user.value?.role === 'admin');

  // --- ACTIONS (Hàm xử lý) ---
  
  // Lưu token và thông tin user sau khi login thành công
  function setTokens(at: string, rt: string) {
    accessToken.value = at;
    refreshToken.value = rt;
  }

  function setUser(userData: User) {
    user.value = userData;
  }

  function logout() {
    accessToken.value = null;
    refreshToken.value = null;
    user.value = null;
    
    localStorage.removeItem('auth_data');
    
    // Điều hướng về trang login (nếu dùng router)
    // window.location.href = '/login';
  }

  return { 
    accessToken, 
    refreshToken, 
    user, 
    isAuthenticated, 
    isAdmin,
    setTokens, 
    setUser,
    logout 
  };
}, {
  persist: {
    pick: ['refreshToken', 'user'],
  }
});