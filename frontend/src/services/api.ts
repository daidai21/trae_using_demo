import axios, { AxiosInstance, InternalAxiosRequestConfig } from 'axios'
import { message } from 'antd'
import type {
  ApiResponse,
  User,
  Merchant,
  Product,
  CartItem,
  Order,
  LoginRequest,
  RegisterRequest,
  LoginResponse,
  MerchantRequest,
  ProductRequest,
  CartRequest,
  OrderRequest,
} from '../types'

const BASE_URL = 'http://localhost:8080/api'

const api: AxiosInstance = axios.create({
  baseURL: BASE_URL,
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
})

api.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    const token = localStorage.getItem('token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

api.interceptors.response.use(
  (response) => {
    const res = response.data as ApiResponse
    if (res.code !== 0) {
      message.error(res.message || '请求失败')
      if (res.code === 401) {
        localStorage.removeItem('token')
        localStorage.removeItem('user')
        window.location.href = '/login'
      }
      return Promise.reject(new Error(res.message || '请求失败'))
    }
    return res
  },
  (error) => {
    message.error(error.message || '网络错误')
    return Promise.reject(error)
  }
)

export const authApi = {
  login: (data: LoginRequest) => api.post<LoginResponse>('/auth/login', data),
  register: (data: RegisterRequest) => api.post<void>('/auth/register', data),
}

export const productApi = {
  list: () => api.get<Product[]>('/products'),
  get: (id: number) => api.get<Product>(`/products/${id}`),
  create: (data: ProductRequest) => api.post<Product>('/products', data),
  update: (id: number, data: ProductRequest) => api.put<Product>(`/products/${id}`, data),
  delete: (id: number) => api.delete<void>(`/products/${id}`),
}

export const cartApi = {
  list: () => api.get<CartItem[]>('/cart'),
  add: (data: CartRequest) => api.post<CartItem>('/cart', data),
  update: (id: number, data: CartRequest) => api.put<CartItem>(`/cart/${id}`, data),
  delete: (id: number) => api.delete<void>(`/cart/${id}`),
}

export const orderApi = {
  list: () => api.get<Order[]>('/orders'),
  get: (id: number) => api.get<Order>(`/orders/${id}`),
  create: (data: OrderRequest) => api.post<Order>('/orders', data),
}

export const merchantApi = {
  list: () => api.get<Merchant[]>('/merchants'),
  get: (id: number) => api.get<Merchant>(`/merchants/${id}`),
  create: (data: MerchantRequest) => api.post<Merchant>('/merchants', data),
  update: (id: number, data: MerchantRequest) => api.put<Merchant>(`/merchants/${id}`, data),
  getMyMerchant: () => api.get<Merchant>('/merchants/me'),
}

export default api
