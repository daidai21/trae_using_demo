import React, { useState, useEffect } from 'react'
import { Table, Button, InputNumber, message, Space, Card, Checkbox } from 'antd'
import { DeleteOutlined, ShoppingOutlined } from '@ant-design/icons'
import { cartApi, orderApi } from '../services/api'
import { CartItem } from '../types'

const Cart: React.FC = () => {
  const [cartItems, setCartItems] = useState<CartItem[]>([])
  const [loading, setLoading] = useState(false)
  const [selectedRowKeys, setSelectedRowKeys] = useState<React.Key[]>([])

  const fetchCart = async () => {
    setLoading(true)
    try {
      const res = await cartApi.list()
      setCartItems(res.data)
    } catch (error) {
      console.error('获取购物车失败:', error)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    fetchCart()
  }, [])

  const updateQuantity = async (id: number, quantity: number) => {
    try {
      await cartApi.update(id, { product_id: 0, quantity })
      message.success('更新成功')
      fetchCart()
    } catch (error) {
      console.error('更新失败:', error)
    }
  }

  const deleteItem = async (id: number) => {
    try {
      await cartApi.delete(id)
      message.success('删除成功')
      fetchCart()
    } catch (error) {
      console.error('删除失败:', error)
    }
  }

  const checkout = async () => {
    if (selectedRowKeys.length === 0) {
      message.warning('请选择要结算的商品')
      return
    }
    try {
      await orderApi.create({ cart_ids: selectedRowKeys.map(Number) })
      message.success('下单成功')
      fetchCart()
      setSelectedRowKeys([])
    } catch (error) {
      console.error('下单失败:', error)
    }
  }

  const total = cartItems
    .filter((item) => selectedRowKeys.includes(item.id))
    .reduce((sum, item) => sum + (item.product?.price || 0) * item.quantity, 0)

  const columns = [
    {
      title: '商品',
      dataIndex: ['product', 'name'],
      key: 'name',
    },
    {
      title: '单价',
      dataIndex: ['product', 'price'],
      key: 'price',
      render: (price: number) => `¥${price.toFixed(2)}`,
    },
    {
      title: '数量',
      key: 'quantity',
      render: (_, record: CartItem) => (
        <InputNumber
          min={1}
          max={record.product?.stock}
          value={record.quantity}
          onChange={(value) => value && updateQuantity(record.id, value)}
        />
      ),
    },
    {
      title: '小计',
      key: 'subtotal',
      render: (_, record: CartItem) =>
        `¥${((record.product?.price || 0) * record.quantity).toFixed(2)}`,
    },
    {
      title: '操作',
      key: 'action',
      render: (_, record: CartItem) => (
        <Button
          type="text"
          danger
          icon={<DeleteOutlined />}
          onClick={() => deleteItem(record.id)}
        >
          删除
        </Button>
      ),
    },
  ]

  const rowSelection = {
    selectedRowKeys,
    onChange: (newSelectedRowKeys: React.Key[]) => {
      setSelectedRowKeys(newSelectedRowKeys)
    },
  }

  return (
    <div>
      <Card title="购物车">
        <Table
          rowSelection={rowSelection}
          columns={columns}
          dataSource={cartItems}
          rowKey="id"
          loading={loading}
          pagination={false}
          footer={() => (
            <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
              <Space>
                <span>已选择 {selectedRowKeys.length} 件商品</span>
                <span style={{ fontSize: '18px', fontWeight: 'bold' }}>
                  合计: <span style={{ color: '#f5222d' }}>¥{total.toFixed(2)}</span>
                </span>
              </Space>
              <Button
                type="primary"
                size="large"
                icon={<ShoppingOutlined />}
                onClick={checkout}
                disabled={selectedRowKeys.length === 0}
              >
                结算
              </Button>
            </div>
          )}
        />
      </Card>
    </div>
  )
}

export default Cart
