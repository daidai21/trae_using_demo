import React, { useState, useEffect } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { Card, Row, Col, Statistic, Button, InputNumber, Space, Typography, Tag, List, message } from 'antd'
import { ClockCircleOutlined, UserOutlined, GavelOutlined, ArrowLeftOutlined } from '@ant-design/icons'
import { useTranslation } from 'react-i18next'
import { auctionApi, useAuth } from '@ecommerce/ui'
import { useAuctionSocket } from '../hooks/useAuctionSocket'

const { Title, Text } = Typography

export const AuctionDetail: React.FC = () => {
  const { t } = useTranslation()
  const { id } = useParams<{ id: string }>()
  const navigate = useNavigate()
  const { user } = useAuth()
  const [auction, setAuction] = useState<any>(null)
  const [bidAmount, setBidAmount] = useState<number>(0)
  const [loading, setLoading] = useState(false)
  const [bidLoading, setBidLoading] = useState(false)

  const { isConnected, messages, onlineCount } = useAuctionSocket(
    Number(id),
    user?.id
  )

  useEffect(() => {
    if (id) {
      loadAuction(Number(id))
    }
  }, [id])

  useEffect(() => {
    if (messages.length > 0) {
      const lastMessage = messages[messages.length - 1]
      if (lastMessage.type === 'update' && lastMessage.data) {
        setAuction((prev: any) => ({
          ...prev,
          ...lastMessage.data,
        }))
      }
    }
  }, [messages])

  const loadAuction = async (auctionId: number) => {
    setLoading(true)
    try {
      const response = await auctionApi.getAuction(auctionId)
      const data = response.data.data
      setAuction(data)
      setBidAmount(data.current_price + data.bid_increment)
    } catch (error) {
      console.error('Failed to load auction:', error)
    } finally {
      setLoading(false)
    }
  }

  const handlePlaceBid = async () => {
    if (!id || !bidAmount) return

    setBidLoading(true)
    try {
      await auctionApi.placeBid(Number(id), bidAmount)
      message.success(t('product.auction.bidSuccess'))
    } catch (error: any) {
      message.error(error.response?.data?.message || t('product.auction.bidFailed'))
    } finally {
      setBidLoading(false)
    }
  }

  const handleStartAuction = async () => {
    if (!id) return
    try {
      await auctionApi.startAuction(Number(id))
      message.success('拍卖已开始')
      loadAuction(Number(id))
    } catch (error: any) {
      message.error(error.response?.data?.message || '启动拍卖失败')
    }
  }

  const handleEndAuction = async () => {
    if (!id) return
    try {
      await auctionApi.endAuction(Number(id))
      message.success('拍卖已结束')
      loadAuction(Number(id))
    } catch (error: any) {
      message.error(error.response?.data?.message || '结束拍卖失败')
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

  if (loading || !auction) {
    return <div>加载中...</div>
  }

  return (
    <div>
      <Button 
        icon={<ArrowLeftOutlined />} 
        onClick={() => navigate('/auctions')}
        style={{ marginBottom: 16 }}
      >
        返回列表
      </Button>

      <Row gutter={[16, 16]}>
        <Col xs={24} lg={14}>
          <Card>
            <div style={{ height: 400, background: '#f5f5f5', display: 'flex', alignItems: 'center', justifyContent: 'center', marginBottom: 16 }}>
              <img 
                src={`https://copilot-cn.bytedance.net/api/ide/v1/text_to_image?prompt=${encodeURIComponent('auction product ' + auction.title)}&image_size=landscape_16_9`}
                alt={auction.title}
                style={{ maxWidth: '100%', maxHeight: '100%', objectFit: 'contain' }}
              />
            </div>

            <Title level={2}>
              <Space>
                {auction.title}
                <Tag color={getStatusColor(auction.status)}>
                  {getStatusText(auction.status)}
                </Tag>
              </Space>
            </Title>
            <Text>{auction.description}</Text>

            {auction.bids && auction.bids.length > 0 && (
              <div style={{ marginTop: 24 }}>
                <Title level={4}>出价历史</Title>
                <List
                  dataSource={auction.bids}
                  renderItem={(bid: any) => (
                    <List.Item>
                      <Space>
                        <UserOutlined />
                        <Text>用户 {bid.user_id}</Text>
                        <Text strong>出价: {bid.amount}</Text>
                        <Text type="secondary">{new Date(bid.created_at).toLocaleString()}</Text>
                      </Space>
                    </List.Item>
                  )}
                />
              </div>
            )}
          </Card>
        </Col>

        <Col xs={24} lg={10}>
          <Card>
            <Space direction="vertical" style={{ width: '100%' }} size="large">
              <Statistic
                title={t('product.auction.currentPrice')}
                value={auction.current_price}
                precision={2}
                valueStyle={{ color: '#cf1322' }}
              />

              <Statistic
                title={t('product.auction.bidCount')}
                value={auction.bid_count}
                prefix={<GavelOutlined />}
              />

              <Statistic
                title="在线人数"
                value={onlineCount}
                valueStyle={{ color: '#3f8600' }}
              />

              {auction.status === 'live' && (
                <Space direction="vertical" style={{ width: '100%' }}>
                  <Space.Compact style={{ width: '100%' }}>
                    <InputNumber
                      style={{ width: '100%' }}
                      size="large"
                      min={auction.current_price + auction.bid_increment}
                      step={auction.bid_increment}
                      value={bidAmount}
                      onChange={setBidAmount}
                      addonBefore={t('product.auction.yourBid')}
                    />
                  </Space.Compact>
                  <Button
                    type="primary"
                    size="large"
                    block
                    loading={bidLoading}
                    onClick={handlePlaceBid}
                  >
                    {t('product.auction.placeBid')}
                  </Button>
                </Space>
              )}

              {auction.status === 'pending' && (
                <Button
                  type="primary"
                  size="large"
                  block
                  onClick={handleStartAuction}
                >
                  开始拍卖
                </Button>
              )}

              {auction.status === 'live' && (
                <Button
                  danger
                  size="large"
                  block
                  onClick={handleEndAuction}
                >
                  结束拍卖
                </Button>
              )}

              <div>
                <Text type="secondary">起拍价: {auction.start_price}</Text>
                <br />
                <Text type="secondary">加价幅度: {auction.bid_increment}</Text>
                {auction.reserve_price > 0 && (
                  <>
                    <br />
                    <Text type="secondary">保留价: {auction.reserve_price}</Text>
                  </>
                )}
              </div>

              <Tag color={isConnected ? 'green' : 'red'}>
                {isConnected ? '实时连接已建立' : '连接断开'}
              </Tag>
            </Space>
          </Card>
        </Col>
      </Row>
    </div>
  )
}
