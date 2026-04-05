import React, { useState, useEffect } from 'react'
import { List, Card, Button, InputNumber, message, Space, Image } from 'antd'
import { DeleteOutlined, ShoppingOutlined } from '@ant-design/icons'
import { useNavigate } from 'react-router-dom'
import { cartApi, buyApi, type CartItem } from '@ecommerce/ui'

export const Cart: React.FC = () => {
  const [cartItems, setCartItems] = useState<CartItem[]>([])
  const [loading, setLoading] = useState(false)
  const navigate = useNavigate()

  useEffect(() => {
    fetchCart()
  }, [])

  const fetchCart = async () => {
    setLoading(true)
    try {
      const res = await cartApi.list()
      setCartItems(res.data)
    } catch (error: any) {
      message.error(error.response?.data?.message || '获取购物车失败')
    } finally {
      setLoading(false)
    }
  }

  const handleUpdateQuantity = async (item: CartItem, quantity: number) => {
    try {
      await cartApi.update(item.id, { quantity })
      fetchCart()
    } catch (error: any) {
      message.error(error.response?.data?.message || '更新失败')
    }
  }

  const handleDeleteItem = async (item: CartItem) => {
    try {
      await cartApi.delete(item.id)
      fetchCart()
      message.success('已删除')
    } catch (error: any) {
      message.error(error.response?.data?.message || '删除失败')
    }
  }

  const handleCheckout = async () => {
    try {
      await buyApi.create()
      message.success('订单创建成功')
      navigate('/orders')
    } catch (error: any) {
      message.error(error.response?.data?.message || '创建订单失败')
    }
  }

  const total = cartItems.reduce((sum, item) => {
    const price = item.product?.price || 0
    return sum + price * item.quantity
  }, 0)

  return (
    <div>
      <h2>购物车</h2>
      {cartItems.length === 0 ? (
        <div style={{ textAlign: 'center', padding: 64 }}>
          <ShoppingOutlined style={{ fontSize: 64, color: '#d9d9d9' }} />
          <p style={{ color: '#999', marginTop: 16 }}>购物车是空的</p>
        </div>
      ) : (
        <>
          <List
            dataSource={cartItems}
            loading={loading}
            renderItem={(item) => (
              <List.Item
                actions={[
                  <Button
                    type="text"
                    danger
                    icon={<DeleteOutlined />}
                    onClick={() => handleDeleteItem(item)}
                  >
                    删除
                  </Button>
                ]}
              >
                <List.Item.Meta
                  avatar={
                    <Image
                      width={64}
                      height={64}
                      src="https://copilot-cn.bytedance.net/api/ide/v1/text_to_image?prompt=product%20thumbnail&image_size=square"
                    />
                  }
                  title={item.product?.name || '商品'}
                  description={
                    <Space direction="vertical" style={{ width: '100%' }}>
                      <span>¥{(item.product?.price || 0).toFixed(2)}</span>
                      <Space>
                        <span>数量:</span>
                        <InputNumber
                          min={1}
                          value={item.quantity}
                          onChange={(val) => val && handleUpdateQuantity(item, val)}
                        />
                      </Space>
                    </Space>
                  }
                />
              </List.Item>
            )}
          />
          <Card style={{ marginTop: 16 }}>
            <Space direction="vertical" style={{ width: '100%', alignItems: 'flex-end' }}>
              <span style={{ fontSize: 18 }}>
                总计: <strong style={{ fontSize: 24, color: '#ff4d4f' }}>¥{total.toFixed(2)}</strong>
              </span>
              <Button type="primary" size="large" onClick={handleCheckout}>
                结算
              </Button>
            </Space>
          </Card>
        </>
      )}
    </div>
  )
}
