import React, { useState, useEffect } from 'react'
import { Card, List, Tag, Button } from 'antd'
import { ShopOutlined, EyeOutlined } from '@ant-design/icons'
import { merchantApi } from '../services/api'
import { Merchant } from '../types'

const MerchantList: React.FC = () => {
  const [merchants, setMerchants] = useState<Merchant[]>([])
  const [loading, setLoading] = useState(false)

  const fetchMerchants = async () => {
    setLoading(true)
    try {
      const res = await merchantApi.list()
      setMerchants(res.data)
    } catch (error) {
      console.error('获取商家列表失败:', error)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    fetchMerchants()
  }, [])

  return (
    <div>
      <Card title="商家列表">
        <List
          grid={{ gutter: 16, column: 1, md: 2, lg: 3 }}
          dataSource={merchants}
          loading={loading}
          renderItem={(merchant) => (
            <List.Item>
              <Card
                hoverable
                cover={
                  <div
                    style={{
                      height: 120,
                      display: 'flex',
                      alignItems: 'center',
                      justifyContent: 'center',
                      background: '#f5f5f5',
                    }}
                  >
                    <ShopOutlined style={{ fontSize: '48px', color: '#1890ff' }} />
                  </div>
                }
              >
                <Card.Meta
                  title={merchant.name}
                  description={
                    <div>
                      <p style={{ color: 'rgba(0, 0, 0, 0.45)', marginBottom: 8 }}>
                        {merchant.description}
                      </p>
                      <p style={{ fontSize: '12px', color: '#666' }}>
                        入驻时间: {new Date(merchant.created_at).toLocaleDateString()}
                      </p>
                    </div>
                  }
                />
              </Card>
            </List.Item>
          )}
        />
      </Card>
    </div>
  )
}

export default MerchantList
