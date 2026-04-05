import React, { useState, useEffect } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { Card, Descriptions, Button, message, Space, InputNumber } from 'antd'
import { ShoppingCartOutlined, ArrowLeftOutlined } from '@ant-design/icons'
import { productApi, cartApi } from '../services/api'
import { Product } from '../types'
import { useAuth } from '../store/AuthContext'

const ProductDetail: React.FC = () => {
  const { id } = useParams<{ id: string }>()
  const [product, setProduct] = useState<Product | null>(null)
  const [loading, setLoading] = useState(false)
  const [quantity, setQuantity] = useState(1)
  const navigate = useNavigate()
  const { isAuthenticated } = useAuth()

  const fetchProduct = async () => {
    if (!id) return
    setLoading(true)
    try {
      const res = await productApi.get(parseInt(id))
      setProduct(res.data)
    } catch (error) {
      console.error('获取商品详情失败:', error)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    fetchProduct()
  }, [id])

  const addToCart = async () => {
    if (!product) return
    if (!isAuthenticated) {
      message.warning('请先登录')
      navigate('/login')
      return
    }
    try {
      await cartApi.add({ product_id: product.id, quantity })
      message.success('已加入购物车')
    } catch (error) {
      console.error('添加购物车失败:', error)
    }
  }

  if (!product && !loading) {
    return <div>商品不存在</div>
  }

  return (
    <div>
      <Button
        icon={<ArrowLeftOutlined />}
        onClick={() => navigate('/')}
        style={{ marginBottom: 16 }}
      >
        返回商品列表
      </Button>
      <Card loading={loading}>
        <Card.Grid
          style={{ width: '40%', textAlign: 'center', padding: '48px' }}
          hoverable={false}
        >
          <div
            style={{
              height: 300,
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'center',
              background: '#f5f5f5',
            }}
          >
            <span style={{ color: '#999' }}>商品图片</span>
          </div>
        </Card.Grid>
        <Card.Grid style={{ width: '60%', padding: '24px' }} hoverable={false}>
          <h1 style={{ marginBottom: 16 }}>{product?.name}</h1>
          <p style={{ color: '#f5222d', fontSize: '24px', fontWeight: 'bold', marginBottom: 16 }}>
            ¥{product?.price.toFixed(2)}
          </p>
          <Descriptions column={1} style={{ marginBottom: 24 }}>
            <Descriptions.Item label="库存">{product?.stock}</Descriptions.Item>
            <Descriptions.Item label="商品描述">{product?.description}</Descriptions.Item>
            {product?.merchant && (
              <Descriptions.Item label="商家">{product.merchant.name}</Descriptions.Item>
            )}
          </Descriptions>
          <Space>
            <span>数量:</span>
            <InputNumber
              min={1}
              max={product?.stock}
              value={quantity}
              onChange={setQuantity}
            />
            <Button
              type="primary"
              size="large"
              icon={<ShoppingCartOutlined />}
              onClick={addToCart}
            >
              加入购物车
            </Button>
          </Space>
        </Card.Grid>
      </Card>
    </div>
  )
}

export default ProductDetail
