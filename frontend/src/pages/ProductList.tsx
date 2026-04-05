import React, { useState, useEffect } from 'react'
import { Card, Row, Col, Button, message } from 'antd'
import { ShoppingCartOutlined, EyeOutlined } from '@ant-design/icons'
import { useNavigate } from 'react-router-dom'
import { productApi, cartApi } from '../services/api'
import { Product } from '../types'
import { useAuth } from '../store/AuthContext'

const ProductList: React.FC = () => {
  const [products, setProducts] = useState<Product[]>([])
  const [loading, setLoading] = useState(false)
  const navigate = useNavigate()
  const { isAuthenticated } = useAuth()

  const fetchProducts = async () => {
    setLoading(true)
    try {
      const res = await productApi.list()
      setProducts(res.data)
    } catch (error) {
      console.error('获取商品列表失败:', error)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    fetchProducts()
  }, [])

  const addToCart = async (product: Product) => {
    if (!isAuthenticated) {
      message.warning('请先登录')
      navigate('/login')
      return
    }
    try {
      await cartApi.add({ product_id: product.id, quantity: 1 })
      message.success('已加入购物车')
    } catch (error) {
      console.error('添加购物车失败:', error)
    }
  }

  return (
    <div>
      <h1 style={{ marginBottom: '24px' }}>商品列表</h1>
      <Row gutter={[16, 16]}>
        {products.map((product) => (
          <Col xs={24} sm={12} md={8} lg={6} key={product.id}>
            <Card
              hoverable
              cover={
                <div
                  style={{
                    height: 200,
                    display: 'flex',
                    alignItems: 'center',
                    justifyContent: 'center',
                    background: '#f5f5f5',
                  }}
                >
                  <span style={{ color: '#999' }}>商品图片</span>
                </div>
              }
              actions={[
                <EyeOutlined key="view" onClick={() => navigate(`/products/${product.id}`)} />,
                <ShoppingCartOutlined key="cart" onClick={() => addToCart(product)} />,
              ]}
            >
              <Card.Meta
                title={product.name}
                description={
                  <div>
                    <p style={{ color: 'rgba(0, 0, 0, 0.45)', marginBottom: 8 }}>
                      {product.description}
                    </p>
                    <p style={{ color: '#f5222d', fontSize: '18px', fontWeight: 'bold' }}>
                      ¥{product.price.toFixed(2)}
                    </p>
                    <p style={{ color: '#666' }}>库存: {product.stock}</p>
                  </div>
                }
              />
            </Card>
          </Col>
        ))}
      </Row>
      {products.length === 0 && !loading && (
        <div style={{ textAlign: 'center', padding: '48px', color: '#999' }}>
          暂无商品
        </div>
      )}
    </div>
  )
}

export default ProductList
