import React, { useState, useEffect } from 'react'
import { List, Card, Button, message } from 'antd'
import { PlusOutlined } from '@ant-design/icons'
import { useNavigate } from 'react-router-dom'
import { merchantApi, type Merchant } from '@ecommerce/ui'

export const MerchantList: React.FC = () => {
  const [merchants, setMerchants] = useState<Merchant[]>([])
  const [loading, setLoading] = useState(false)
  const navigate = useNavigate()

  useEffect(() => {
    fetchMerchants()
  }, [])

  const fetchMerchants = async () => {
    setLoading(true)
    try {
      const res = await merchantApi.list()
      setMerchants(res.data)
    } catch (error: any) {
      message.error(error.response?.data?.message || '获取商家列表失败')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 16 }}>
        <h2>商家列表</h2>
        <Button type="primary" icon={<PlusOutlined />} onClick={() => navigate('/merchants/new')}>
          商家入驻
        </Button>
      </div>
      <List
        grid={{ gutter: 16, column: 3 }}
        dataSource={merchants}
        loading={loading}
        renderItem={(merchant) => (
          <List.Item>
            <Card
              hoverable
              onClick={() => navigate(`/merchants/${merchant.id}`)}
            >
              <Card.Meta
                title={merchant.name}
                description={merchant.description}
              />
            </Card>
          </List.Item>
        )}
      />
    </div>
  )
}
