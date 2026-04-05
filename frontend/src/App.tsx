import React from 'react'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import { ConfigProvider } from 'antd'
import zhCN from 'antd/locale/zh_CN'
import { AuthProvider } from './store/AuthContext'
import AppLayout from './components/Layout'
import Login from './pages/Login'
import Register from './pages/Register'
import ProductList from './pages/ProductList'
import ProductDetail from './pages/ProductDetail'
import Cart from './pages/Cart'
import OrderList from './pages/OrderList'
import OrderDetail from './pages/OrderDetail'
import MerchantList from './pages/MerchantList'
import MerchantForm from './pages/MerchantForm'
import ProductForm from './pages/ProductForm'

const App: React.FC = () => {
  return (
    <ConfigProvider locale={zhCN}>
      <AuthProvider>
        <Router>
          <AppLayout>
            <Routes>
              <Route path="/" element={<ProductList />} />
              <Route path="/login" element={<Login />} />
              <Route path="/register" element={<Register />} />
              <Route path="/products/:id" element={<ProductDetail />} />
              <Route path="/cart" element={<Cart />} />
              <Route path="/orders" element={<OrderList />} />
              <Route path="/orders/:id" element={<OrderDetail />} />
              <Route path="/merchants" element={<MerchantList />} />
              <Route path="/merchant/form" element={<MerchantForm />} />
              <Route path="/product/form" element={<ProductForm />} />
            </Routes>
          </AppLayout>
        </Router>
      </AuthProvider>
    </ConfigProvider>
  )
}

export default App
