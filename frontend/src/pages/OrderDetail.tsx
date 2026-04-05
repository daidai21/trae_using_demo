import React, { useState, useEffect } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { Card, Descriptions, Table, Tag, Button } from 'antd'
import { ArrowLeftOutlined } from '@ant-design/icons'
import { orderApi } from '../services/api'
import { Order } from '../types'

const OrderDetail: React.FC = () => {
  const { id } = useParams<{ id: string }>()
  const [order, setOrder] = useState<Order | null>(null)
  const [loading, setLoading] = useState(false)
  const navigate = useNavigate()

  const fetchOrder = async () => {
    if (!id) return
    setLoading(true)
    try {
      const res = await orderApi.get(parseInt(id))
      setOrder(res.data)
    } catch (error) {
      console.error('获取订单详情失败:', error)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    fetchOrder()
  }, [id])

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'pending':
        return 'orange'
      case 'paid':
        return 'blue'
      case 'shipped':
        return 'cyan'
      case 'completed':
        return 'green'
      case 'cancelled':
        return 'red'
      default:
        return 'default'
    }
  }

  const getStatusText = (status: string) => {
    switch (status) {
      case 'pending':
        return '待支付'
      case 'paid':
        return '已支付'
      case 'shipped':
        return '已发货'
      case 'completed':
        return '已完成'
      case 'cancelled':
        return '已取消'
      default:
        return status
    }
  }

  const columns = [
    {
      title: '商品',
      dataIndex: ['product', 'name'],
      key: 'name',
    },
    {
      title: '单价',
      dataIndex: 'price',
      key: 'price',
      render: (price: number) => `¥${price.toFixed(2)}`,
    },
    {
      title: '数量',
      dataIndex: 'quantity',
      key: 'quantity',
    },
    {
      title: '小计',
      key: 'subtotal',
      render: (_, record: any) => `¥${(record.price * record.quantity).toFixed(2)}`,
    },
  ]

  if (!order && !loading) {
    return <div>订单不存在</div>
  }

  return (
    <div>
      <Button
        icon={<ArrowLeftOutlined />}
        onClick={() => navigate('/orders')}
        style={{ marginBottom: 16 }}
      >
        返回订单列表
      </Button>
      <Card title={`订单 #${order?.id}`} loading={loading}>
        <Descriptions column={2} style={{ marginBottom: 24 }}>
          <Descriptions.Item label="订单状态">
            <Tag color={getStatusColor(order?.status || '')}>
              {getStatusText(order?.status || '')}
            </Tag>
          </Descriptions.Item>
          <Descriptions.Item label="总金额">
            <span style={{ fontSize: '18px', fontWeight: 'bold', color: '#f5222d' }}>
              ¥{order?.total_amount.toFixed(2)}
            </span>
          </Descriptions.Item>
          <Descriptions.Item label="下单时间">
            {order?.created_at ? new Date(order.created_at).toLocaleString() : ''}
          </Descriptions.Item>
          <Descriptions.Item label="更新时间">
            {order?.updated_at ? new Date(order.updated_at).toLocaleString() : ''}
          </Descriptions.Item>
        </Descriptions>
        <h3 style={{ marginBottom: 16 }}>商品明细</h3>
        <Table
          columns={columns}
          dataSource={order?.order_items}
          rowKey="id"
          pagination={false}
        />
      </Card>
    </div>
  )
}

export default OrderDetail
