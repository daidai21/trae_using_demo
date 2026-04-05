import React, { useState, useEffect } from 'react'
import { List, Card, Tag, Button, message } from 'antd'
import { useNavigate } from 'react-router-dom'
import { orderApi, type Order } from '@ecommerce/ui'

const statusMap: Record<string, string> = {
  pending: '待支付',
  paid: '已支付',
  shipped: '已发货',
  delivered: '已送达',
  cancelled: '已取消'
}

const statusColorMap: Record<string, string> = {
  pending: 'orange',
  paid: 'blue',
  shipped: 'cyan',
  delivered: 'green',
  cancelled: 'red'
}

export const OrderList: React.FC = () => {
  const [orders, setOrders] = useState<Order[]>([])
  const [loading, setLoading] = useState(false)
  const navigate = useNavigate()

  useEffect(() => {
    fetchOrders()
  }, [])

  const fetchOrders = async () => {
    setLoading(true)
    try {
      const res = await orderApi.list()
      setOrders(res.data)
    } catch (error: any) {
      message.error(error.response?.data?.message || '获取订单列表失败')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div>
      <h2>我的订单</h2>
      <List
        grid={{ gutter: 16, column: 1 }}
        dataSource={orders}
        loading={loading}
        renderItem={(order) => (
          <List.Item>
            <Card
              hoverable
              onClick={() => navigate(`/orders/${order.id}`)}
            >
              <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                <div>
                  <h3 style={{ marginBottom: 8 }}>订单 #{order.id}</h3>
                  <p style={{ margin: 4, color: '#666' }}>
                    下单时间: {order.created_at}
                  </p>
                  <p style={{ margin: 4, fontSize: 18 }}>
                    总金额: <strong style={{ color: '#ff4d4f' }}>¥{order.total_amount.toFixed(2)}</strong>
                  </p>
                </div>
                <Tag color={statusColorMap[order.status] || 'default'}>
                  {statusMap[order.status] || order.status}
                </Tag>
              </div>
            </Card>
          </List.Item>
        )}
      />
    </div>
  )
}
