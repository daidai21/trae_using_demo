import React, { useState } from 'react'
import { Form, Input, InputNumber, Button, Card, message } from 'antd'
import { useNavigate } from 'react-router-dom'
import { productApi } from '../services/api'

const ProductForm: React.FC = () => {
  const [form] = Form.useForm()
  const [loading, setLoading] = useState(false)
  const navigate = useNavigate()

  const onFinish = async (values: {
    name: string
    description: string
    price: number
    stock: number
  }) => {
    setLoading(true)
    try {
      await productApi.create(values)
      message.success('商品发布成功')
      navigate('/')
    } catch (error) {
      console.error('发布商品失败:', error)
    } finally {
      setLoading(false)
    }
  }

  return (
    <div>
      <Card title="发布商品">
        <Form
          form={form}
          layout="vertical"
          onFinish={onFinish}
          initialValues={{ description: '', stock: 0 }}
        >
          <Form.Item
            name="name"
            label="商品名称"
            rules={[{ required: true, message: '请输入商品名称' }]}
          >
            <Input placeholder="请输入商品名称" />
          </Form.Item>

          <Form.Item
            name="description"
            label="商品描述"
          >
            <Input.TextArea rows={4} placeholder="请输入商品描述" />
          </Form.Item>

          <Form.Item
            name="price"
            label="价格"
            rules={[{ required: true, message: '请输入价格' }]}
          >
            <InputNumber
              style={{ width: '100%' }}
              placeholder="请输入价格"
              min={0.01}
              precision={2}
            />
          </Form.Item>

          <Form.Item
            name="stock"
            label="库存"
            rules={[{ required: true, message: '请输入库存' }]}
          >
            <InputNumber
              style={{ width: '100%' }}
              placeholder="请输入库存"
              min={0}
            />
          </Form.Item>

          <Form.Item>
            <Button type="primary" htmlType="submit" loading={loading}>
              发布商品
            </Button>
          </Form.Item>
        </Form>
      </Card>
    </div>
  )
}

export default ProductForm
