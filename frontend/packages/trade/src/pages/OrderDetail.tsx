import React, { useState, useEffect } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { Card, Descriptions, List, Tag, Button, message, Select, Image } from 'antd'
import { ArrowLeftOutlined } from '@ant-design/icons'
import { orderApi, type Order } from '@ecommerce/ui'

const { Option } = Select

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

export const OrderDetail: React.FC = () => {
  const { id } = useParams<{ id: string }>()
  const navigate = useNavigate()
  const [order, setOrder] = useState<Order | null>(null)
  const [loading, setLoading] = useState(false)
  const [updating, setUpdating] = useState(false)

  useEffect(() => {
    if (id) {
      fetchOrder(Number(id))
    }
  }, [id])

  const fetchOrder = async (orderId: number) => {
    setLoading(true)
    try {
      const res = await orderApi.get(orderId)
      setOrder(res.data)
    } catch (error: any) {
      message.error(error.response?.data?.message || '获取订单详情失败')
    } finally {
      setLoading(false)
    }
  }

  const handleStatusChange = async (status: string) => {
    if (!order) return
    setUpdating(true)
    try {
      await orderApi.updateStatus(order.id, status)
      message.success('订单状态已更新')
      fetchOrder(order.id)
    } catch (error: any) {
      message.error(error.response?.data?.message || '更新订单状态失败')
    } finally {
      setUpdating(false)
    }
  }

  if (!order) return null

  return (
    <div>
      <Button icon={<ArrowLeftOutlined />} onClick={() => navigate(-1)} style={{ marginBottom: 16 }}>
        返回
      </Button>
      <Card loading={loading}>
        <div style={{ marginBottom: 24 }}>
          <h2>订单 #{order.id}</h2>
          <Tag color={statusColorMap[order.status] || 'default'} style={{ marginTop: 8 }}>
            {statusMap[order.status] || order.status}
          </Tag>
        </div>
        <Descriptions title="订单信息" bordered column={2} style={{ marginBottom: 24 }}>
          <Descriptions.Item label="订单ID">{order.id}</Descriptions.Item>
          <Descriptions.Item label="总金额">¥{order.total_amount.toFixed(2)}</Descriptions.Item>
          <Descriptions.Item label="创建时间">{order.created_at}</Descriptions.Item>
          <Descriptions.Item label="更新时间">{order.updated_at}</Descriptions.Item>
        </Descriptions>
        <div style={{ marginBottom: 24 }}>
          <h3 style={{ marginBottom: 16 }}>订单商品</h3>
          <List
            dataSource={order.order_items}
            renderItem={(item) => (
              <List.Item>
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
                    <>
                      <p>单价: ¥{item.price.toFixed(2)}</p>
                      <p>数量: {item.quantity}</p>
                      <p>小计: ¥{(item.price * item.quantity).toFixed(2)}</p>
                    </>
                  }
                />
              </List.Item>
            )}
          />
        </div>
        <div>
          <h3 style={{ marginBottom: 16 }}>更新订单状态</h3>
          <Select
            style={{ width: 200, marginRight: 16 }}
            value={order.status}
            onChange={handleStatusChange}
            loading={updating}
          >
            <Option value="pending">待支付</Option>
            <Option value="paid">已支付</Option>
            <Option value="shipped">已发货</Option>
            <Option value="delivered">已送达</Option>
            <Option value="cancelled">已取消</Option>
          </Select>
        </div>
      </Card>
    </div>
  )
}
