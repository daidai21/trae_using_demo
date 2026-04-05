export interface User {
  id: number
  username: string
  created_at: string
  updated_at: string
}

export interface Merchant {
  id: number
  user_id: number
  name: string
  description: string
  created_at: string
  updated_at: string
  user?: User
}

export interface Product {
  id: number
  merchant_id: number
  name: string
  description: string
  price: number
  stock: number
  created_at: string
  updated_at: string
  merchant?: Merchant
}

export interface CartItem {
  id: number
  user_id: number
  product_id: number
  quantity: number
  created_at: string
  updated_at: string
  user?: User
  product?: Product
}

export interface OrderItem {
  id: number
  order_id: number
  product_id: number
  quantity: number
  price: number
  created_at: string
  updated_at: string
  product?: Product
}

export interface Order {
  id: number
  user_id: number
  total_amount: number
  status: string
  created_at: string
  updated_at: string
  user?: User
  order_items?: OrderItem[]
}

export interface ApiResponse<T = any> {
  code: number
  message: string
  data: T
}

export interface LoginRequest {
  username: string
  password: string
}

export interface RegisterRequest {
  username: string
  password: string
}

export interface LoginResponse {
  token: string
  user: User
}

export interface MerchantRequest {
  name: string
  description: string
}

export interface ProductRequest {
  name: string
  description: string
  price: number
  stock: number
}

export interface CartRequest {
  product_id: number
  quantity: number
}

export interface OrderStatusRequest {
  status: string
}
