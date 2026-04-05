import React, { useState, useEffect } from 'react'
import { List, Card, Tag, Button, Space, Typography } from 'antd'
import { useNavigate } from 'react-router-dom'
import { PlusOutlined, ClockCircleOutlined, UserOutlined } from '@ant-design/icons'
import { useTranslation } from 'react-i18next'
import { auctionApi } from '@ecommerce/ui'

const { Title, Text } = Typography

export const AuctionList: React.FC = () => {
  const { t } = useTranslation()
  const navigate = useNavigate()
  const [auctions, setAuctions] = useState<any[]>([])
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    loadAuctions()
  }, [])

  const loadAuctions = async () => {
    setLoading(true)
    try {
      const response = await auctionApi.getLiveAuctions()
      setAuctions(response.data.data || [])
    } catch (error) {
      console.error('Failed to load auctions:', error)
    } finally {
      setLoading(false)
    }
  }

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'live':
        return 'green'
      case 'pending':
        return 'blue'
      case 'sold':
        return 'purple'
      case 'unsold':
        return 'default'
      default:
        return 'default'
    }
  }

  const getStatusText = (status: string) => {
    switch (status) {
      case 'live':
        return '拍卖中'
      case 'pending':
        return '未开始'
      case 'sold':
        return '已成交'
      case 'unsold':
        return '流拍'
      default:
        return status
    }
  }

  return (
    <div>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 24 }}>
        <Title level={2}>{t('product.auction.title')}</Title>
        <Button type="primary" icon={<PlusOutlined />} onClick={() => navigate('/auctions/create')}>
          发起拍卖
        </Button>
      </div>

      <List
        grid={{ gutter: 16, xs: 1, sm: 2, md: 3, lg: 3, xl: 4 }}
        dataSource={auctions}
        loading={loading}
        renderItem={(auction) => (
          <List.Item>
            <Card
              hoverable
              onClick={() => navigate(`/auctions/${auction.id}`)}
              cover={
                <div style={{ height: 200, background: '#f5f5f5', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
                  <img 
                    src={`https://copilot-cn.bytedance.net/api/ide/v1/text_to_image?prompt=${encodeURIComponent('auction product ' + auction.title)}&image_size=square`}
                    alt={auction.title}
                    style={{ maxWidth: '100%', maxHeight: '100%', objectFit: 'cover' }}
                  />
                </div>
              }
            >
              <Card.Meta
                title={
                  <Space>
                    {auction.title}
                    <Tag color={getStatusColor(auction.status)}>
                      {getStatusText(auction.status)}
                    </Tag>
                  </Space>
                }
                description={
                  <Space direction="vertical" style={{ width: '100%' }}>
                    <Text strong>当前价格: {auction.current_price}</Text>
                    <Text type="secondary">
                      <UserOutlined /> {auction.bid_count} 次出价
                    </Text>
                    {auction.end_time && (
                      <Text type="secondary">
                        <ClockCircleOutlined /> 结束时间: {new Date(auction.end_time).toLocaleString()}
                      </Text>
                    )}
                  </Space>
                }
              />
            </Card>
          </List.Item>
        )}
      />
    </div>
  )
}
