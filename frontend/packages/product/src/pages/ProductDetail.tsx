import React, { useState, useEffect } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { Card, Button, Descriptions, message, Image, Space } from 'antd'
import { ShoppingCartOutlined, ArrowLeftOutlined } from '@ant-design/icons'
import { productApi, cartApi, type Product } from '@ecommerce/ui'

export const ProductDetail: React.FC = () => {
  const { id } = useParams<{ id: string }>()
  const navigate = useNavigate()
  const [product, setProduct] = useState<Product | null>(null)
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    if (id) {
      fetchProduct(Number(id))
    }
  }, [id])

  const fetchProduct = async (productId: number) => {
    setLoading(true)
    try {
      const res = await productApi.get(productId)
      setProduct(res.data)
    } catch (error: any) {
      message.error(error.response?.data?.message || '获取商品详情失败')
    } finally {
      setLoading(false)
    }
  }

  const handleAddToCart = async () => {
    if (!product) return
    try {
      await cartApi.add({ product_id: product.id, quantity: 1 })
      message.success('已添加到购物车')
    } catch (error: any) {
      message.error(error.response?.data?.message || '添加到购物车失败')
    }
  }

  if (!product) return null

  return (
    <div>
      <Button icon={<ArrowLeftOutlined />} onClick={() => navigate(-1)} style={{ marginBottom: 16 }}>
        返回
      </Button>
      <Card loading={loading}>
        <Space direction="vertical" size="large" style={{ width: '100%' }}>
          <div style={{ display: 'flex', gap: 32 }}>
            <div style={{ flex: 1 }}>
              <Image
                alt={product.name}
                src="https://copilot-cn.bytedance.net/api/ide/v1/text_to_image?prompt=ecommerce%20product%20detail%20image&image_size=square"
                style={{ width: '100%', height: 400, objectFit: 'cover' }}
              />
            </div>
            <div style={{ flex: 1 }}>
              <h1 style={{ fontSize: 28, marginBottom: 16 }}>{product.name}</h1>
              <p style={{ fontSize: 32, color: '#ff4d4f', marginBottom: 16 }}>
                ¥{product.price.toFixed(2)}
              </p>
              <p style={{ color: '#666', marginBottom: 24 }}>{product.description}</p>
              <p style={{ marginBottom: 24 }}>库存: {product.stock}</p>
              <Button
                type="primary"
                size="large"
                icon={<ShoppingCartOutlined />}
                onClick={handleAddToCart}
                disabled={product.stock <= 0}
              >
                加入购物车
              </Button>
            </div>
          </div>
          <Descriptions title="商品详情" bordered column={1}>
            <Descriptions.Item label="商品ID">{product.id}</Descriptions.Item>
            <Descriptions.Item label="价格">¥{product.price.toFixed(2)}</Descriptions.Item>
            <Descriptions.Item label="库存">{product.stock}</Descriptions.Item>
            <Descriptions.Item label="创建时间">{product.created_at}</Descriptions.Item>
          </Descriptions>
        </Space>
      </Card>
    </div>
  )
}
