import React, { useState } from 'react'
import { Form, Input, Button, Card, message } from 'antd'
import { useNavigate } from 'react-router-dom'
import { merchantApi, type MerchantRequest } from '@ecommerce/ui'

export const MerchantForm: React.FC = () => {
  const [form] = Form.useForm()
  const [loading, setLoading] = useState(false)
  const navigate = useNavigate()

  const onFinish = async (values: MerchantRequest) => {
    setLoading(true)
    try {
      await merchantApi.create(values)
      message.success('商家入驻成功')
      navigate('/merchants')
    } catch (error: any) {
      message.error(error.response?.data?.message || '商家入驻失败')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div>
      <Card title="商家入驻">
        <Form
          form={form}
          layout="vertical"
          onFinish={onFinish}
        >
          <Form.Item
            name="name"
            label="商家名称"
            rules={[{ required: true, message: '请输入商家名称' }]}
          >
            <Input placeholder="请输入商家名称" />
          </Form.Item>

          <Form.Item
            name="description"
            label="商家描述"
            rules={[{ required: true, message: '请输入商家描述' }]}
          >
            <Input.TextArea rows={4} placeholder="请输入商家描述" />
          </Form.Item>

          <Form.Item>
            <Button type="primary" htmlType="submit" loading={loading}>
              入驻
            </Button>
          </Form.Item>
        </Form>
      </Card>
    </div>
  )
}
