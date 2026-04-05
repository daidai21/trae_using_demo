import React from 'react'
import { Routes, Route, Navigate } from 'react-router-dom'
import { Layout, useAuth } from '@ecommerce/ui'
import { Login, Register } from '@ecommerce/user'
import { ProductList, ProductDetail, ProductForm, MerchantList, MerchantForm } from '@ecommerce/product'
import { Cart, OrderList, OrderDetail } from '@ecommerce/trade'
import { AuctionList, AuctionDetail, CreateAuction } from '@ecommerce/auction'

const ProtectedRoute: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const { isAuthenticated } = useAuth()
  return isAuthenticated ? <>{children}</> : <Navigate to="/login" replace />
}

const App: React.FC = () => {
  return (
    <Layout>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/" element={<Navigate to="/products" replace />} />
        <Route path="/products" element={
          <ProtectedRoute>
            <ProductList />
          </ProtectedRoute>
        } />
        <Route path="/products/:id" element={
          <ProtectedRoute>
            <ProductDetail />
          </ProtectedRoute>
        } />
        <Route path="/products/new" element={
          <ProtectedRoute>
            <ProductForm />
          </ProtectedRoute>
        } />
        <Route path="/merchants" element={
          <ProtectedRoute>
            <MerchantList />
          </ProtectedRoute>
        } />
        <Route path="/merchants/new" element={
          <ProtectedRoute>
            <MerchantForm />
          </ProtectedRoute>
        } />
        <Route path="/cart" element={
          <ProtectedRoute>
            <Cart />
          </ProtectedRoute>
        } />
        <Route path="/orders" element={
          <ProtectedRoute>
            <OrderList />
          </ProtectedRoute>
        } />
        <Route path="/orders/:id" element={
          <ProtectedRoute>
            <OrderDetail />
          </ProtectedRoute>
        } />
        <Route path="/auctions" element={
          <ProtectedRoute>
            <AuctionList />
          </ProtectedRoute>
        } />
        <Route path="/auctions/create" element={
          <ProtectedRoute>
            <CreateAuction />
          </ProtectedRoute>
        } />
        <Route path="/auctions/:id" element={
          <ProtectedRoute>
            <AuctionDetail />
          </ProtectedRoute>
        } />
      </Routes>
    </Layout>
  )
}

export default App
