import axios from 'axios'
import type {
  ApiResponse,
  LoginRequest,
  LoginResponse,
  RegisterRequest,
  User,
  Merchant,
  MerchantRequest,
  Product,
  ProductRequest,
  CartItem,
  CartRequest,
  Order
} from '../types'

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080'

const api = axios.create({
  baseURL: API_BASE_URL,
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json'
  }
})

api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('token')
      localStorage.removeItem('user')
      window.location.href = '/login'
    }
    return Promise.reject(error)
  }
)

export const authApi = {
  async register(data: RegisterRequest): Promise<ApiResponse<User>> {
    const res = await api.post('/api/auth/register', data)
    return res.data
  },

  async login(data: LoginRequest): Promise<ApiResponse<LoginResponse>> {
    const res = await api.post('/api/auth/login', data)
    return res.data
  },

  async getProfile(): Promise<ApiResponse<User>> {
    const res = await api.get('/api/users/profile')
    return res.data
  }
}

export const merchantApi = {
  async create(data: MerchantRequest): Promise<ApiResponse<Merchant>> {
    const res = await api.post('/api/merchants', data)
    return res.data
  },

  async list(): Promise<ApiResponse<Merchant[]>> {
    const res = await api.get('/api/merchants')
    return res.data
  },

  async get(id: number): Promise<ApiResponse<Merchant>> {
    const res = await api.get(`/api/merchants/${id}`)
    return res.data
  },

  async update(id: number, data: MerchantRequest): Promise<ApiResponse<Merchant>> {
    const res = await api.put(`/api/merchants/${id}`, data)
    return res.data
  }
}

export const productApi = {
  async create(data: ProductRequest): Promise<ApiResponse<Product>> {
    const res = await api.post('/api/products', data)
    return res.data
  },

  async list(): Promise<ApiResponse<Product[]>> {
    const res = await api.get('/api/products')
    return res.data
  },

  async get(id: number): Promise<ApiResponse<Product>> {
    const res = await api.get(`/api/products/${id}`)
    return res.data
  },

  async update(id: number, data: ProductRequest): Promise<ApiResponse<Product>> {
    const res = await api.put(`/api/products/${id}`, data)
    return res.data
  },

  async delete(id: number): Promise<ApiResponse<void>> {
    const res = await api.delete(`/api/products/${id}`)
    return res.data
  }
}

export const cartApi = {
  async add(data: CartRequest): Promise<ApiResponse<CartItem>> {
    const res = await api.post('/api/cart', data)
    return res.data
  },

  async list(): Promise<ApiResponse<CartItem[]>> {
    const res = await api.get('/api/cart')
    return res.data
  },

  async update(id: number, data: { quantity: number }): Promise<ApiResponse<CartItem>> {
    const res = await api.put(`/api/cart/${id}`, data)
    return res.data
  },

  async delete(id: number): Promise<ApiResponse<void>> {
    const res = await api.delete(`/api/cart/${id}`)
    return res.data
  }
}

export const buyApi = {
  async create(): Promise<ApiResponse<Order>> {
    const res = await api.post('/api/buy')
    return res.data
  }
}

export const orderApi = {
  async list(): Promise<ApiResponse<Order[]>> {
    const res = await api.get('/api/orders')
    return res.data
  },

  async get(id: number): Promise<ApiResponse<Order>> {
    const res = await api.get(`/api/orders/${id}`)
    return res.data
  },

  async updateStatus(id: number, status: string): Promise<ApiResponse<Order>> {
    const res = await api.put(`/api/orders/${id}/status`, { status })
    return res.data
  }
}

export const auctionApi = {
  async createAuction(data: any): Promise<ApiResponse<any>> {
    const res = await api.post('/api/auctions', data)
    return res.data
  },

  async getAuction(id: number): Promise<ApiResponse<any>> {
    const res = await api.get(`/api/auctions/${id}`)
    return res.data
  },

  async getAllAuctions(): Promise<ApiResponse<any[]>> {
    const res = await api.get('/api/auctions')
    return res.data
  },

  async getLiveAuctions(): Promise<ApiResponse<any[]>> {
    const res = await api.get('/api/auctions/live')
    return res.data
  },

  async startAuction(id: number): Promise<ApiResponse<void>> {
    const res = await api.post(`/api/auctions/${id}/start`)
    return res.data
  },

  async placeBid(id: number, amount: number): Promise<ApiResponse<any>> {
    const res = await api.post(`/api/auctions/${id}/bid`, { amount })
    return res.data
  },

  async endAuction(id: number): Promise<ApiResponse<void>> {
    const res = await api.post(`/api/auctions/${id}/end`)
    return res.data
  }
}
