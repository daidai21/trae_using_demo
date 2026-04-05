import React, { useState, useEffect } from 'react'
import { List, Card, Button, message, Image, Space } from 'antd'
import { ShoppingCartOutlined } from '@ant-design/icons'
import { useNavigate } from 'react-router-dom'
import { productApi, cartApi, type Product } from '@ecommerce/ui'

export const ProductList: React.FC = () => {
  const [products, setProducts] = useState<Product[]>([])
  const [loading, setLoading] = useState(false)
  const navigate = useNavigate()

  useEffect(() => {
    fetchProducts()
  }, [])

  const fetchProducts = async () => {
    setLoading(true)
    try {
      const res = await productApi.list()
      setProducts(res.data)
    } catch (error: any) {
      message.error(error.response?.data?.message || '获取商品列表失败')
    } finally {
      setLoading(false)
    }
  }

  const handleAddToCart = async (product: Product) => {
    try {
      await cartApi.add({ product_id: product.id, quantity: 1 })
      message.success('已添加到购物车')
    } catch (error: any) {
      message.error(error.response?.data?.message || '添加到购物车失败')
    }
  }

  return (
    <div>
      <h2>商品列表</h2>
      <List
        grid={{ gutter: 16, column: 4 }}
        dataSource={products}
        loading={loading}
        renderItem={(product) => (
          <List.Item>
            <Card
              hoverable
              style={{ cursor: 'pointer' }}
              cover={
                <div style={{ height: 200, display: 'flex', alignItems: 'center', justifyContent: 'center', background: '#f5f5f5' }}>
                  <Image
                    alt={product.name}
                    src="https://copilot-cn.bytedance.net/api/ide/v1/text_to_image?prompt=ecommerce%20product%20image&image_size=square"
                    style={{ width: '100%', height: '100%', objectFit: 'cover' }}
                  />
                </div>
              }
              onClick={() => navigate(`/products/${product.id}`)}
              actions={[
                <Button
                  type="primary"
                  icon={<ShoppingCartOutlined />}
                  onClick={(e) => {
                    e.stopPropagation()
                    handleAddToCart(product)
                  }}
                >
                  加入购物车
                </Button>
              ]}
            >
              <Card.Meta
                title={product.name}
                description={
                  <Space direction="vertical" style={{ width: '100%' }}>
                    <span>¥{product.price.toFixed(2)}</span>
                    <span style={{ color: '#666' }}>库存: {product.stock}</span>
                  </Space>
                }
              />
              <p style={{ color: '#666', fontSize: '12px', marginTop: 8 }}>
                {product.description}
              </p>
            </Card>
          </List.Item>
        )}
      />
    </div>
  )
}
