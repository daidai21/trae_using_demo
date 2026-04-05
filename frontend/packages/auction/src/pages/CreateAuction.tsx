import React, { useState } from 'react'
import { Form, Input, InputNumber, Select, Button, Card, Typography, message } from 'antd'
import { useNavigate } from 'react-router-dom'
import { ArrowLeftOutlined } from '@ant-design/icons'
import { useTranslation } from 'react-i18next'
import { auctionApi } from '@ecommerce/ui'

const { Title } = Typography
const { TextArea } = Input
const { Option } = Select

export const CreateAuction: React.FC = () => {
  const { t } = useTranslation()
  const navigate = useNavigate()
  const [form] = Form.useForm()
  const [loading, setLoading] = useState(false)

  const handleSubmit = async (values: any) => {
    setLoading(true)
    try {
      await auctionApi.createAuction({
        ...values,
        currency: values.currency || 'CNY',
      })
      message.success('拍卖创建成功')
      navigate('/auctions')
    } catch (error: any) {
      message.error(error.response?.data?.message || '创建拍卖失败')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div>
      <Button 
        icon={<ArrowLeftOutlined />} 
        onClick={() => navigate('/auctions')}
        style={{ marginBottom: 16 }}
      >
        返回
      </Button>

      <Card>
        <Title level={2}>发起拍卖</Title>
        <Form
          form={form}
          layout="vertical"
          onFinish={handleSubmit}
          initialValues={{
            currency: 'CNY',
            bid_increment: 10,
            duration: 3600,
          }}
        >
          <Form.Item
            name="title"
            label="拍卖标题"
            rules={[{ required: true, message: '请输入拍卖标题' }]}
          >
            <Input placeholder="请输入拍卖标题" />
          </Form.Item>

          <Form.Item
            name="description"
            label="拍卖描述"
          >
            <TextArea rows={4} placeholder="请输入拍卖描述" />
          </Form.Item>

          <Form.Item
            name="product_id"
            label="关联商品ID"
            rules={[{ required: true, message: '请输入关联商品ID' }]}
          >
            <InputNumber style={{ width: '100%' }} placeholder="请输入关联商品ID" />
          </Form.Item>

          <Form.Item
            name="currency"
            label="货币"
            rules={[{ required: true, message: '请选择货币' }]}
          >
            <Select placeholder="请选择货币">
              <Option value="CNY">CNY (¥)</Option>
              <Option value="USD">USD ($)</Option>
              <Option value="IDR">IDR (Rp)</Option>
            </Select>
          </Form.Item>

          <Form.Item
            name="start_price"
            label="起拍价"
            rules={[{ required: true, message: '请输入起拍价' }]}
          >
            <InputNumber style={{ width: '100%' }} placeholder="请输入起拍价" min={0} />
          </Form.Item>

          <Form.Item
            name="reserve_price"
            label="保留价 (可选)"
          >
            <InputNumber style={{ width: '100%' }} placeholder="请输入保留价" min={0} />
          </Form.Item>

          <Form.Item
            name="bid_increment"
            label="加价幅度"
            rules={[{ required: true, message: '请输入加价幅度' }]}
          >
            <InputNumber style={{ width: '100%' }} placeholder="请输入加价幅度" min={1} />
          </Form.Item>

          <Form.Item
            name="duration"
            label="拍卖时长 (秒)"
            rules={[{ required: true, message: '请输入拍卖时长' }]}
          >
            <InputNumber style={{ width: '100%' }} placeholder="请输入拍卖时长" min={60} />
          </Form.Item>

          <Form.Item>
            <Button type="primary" htmlType="submit" loading={loading} block size="large">
              创建拍卖
            </Button>
          </Form.Item>
        </Form>
      </Card>
    </div>
  )
}
