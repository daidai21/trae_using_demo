import React, { useState } from 'react'
import { Form, Input, InputNumber, Button, Card, message } from 'antd'
import { useNavigate } from 'react-router-dom'
import { productApi, type ProductRequest } from '@ecommerce/ui'

export const ProductForm: React.FC = () => {
  const [form] = Form.useForm()
  const [loading, setLoading] = useState(false)
  const navigate = useNavigate()

  const onFinish = async (values: ProductRequest) => {
    setLoading(true)
    try {
      await productApi.create(values)
      message.success('商品发布成功')
      navigate('/products')
    } catch (error: any) {
      message.error(error.response?.data?.message || '商品发布失败')
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
          initialValues={{ price: 0, stock: 0 }}
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
            rules={[{ required: true, message: '请输入商品描述' }]}
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
              min={0.01}
              precision={2}
              placeholder="请输入价格"
            />
          </Form.Item>

          <Form.Item
            name="stock"
            label="库存"
            rules={[{ required: true, message: '请输入库存' }]}
          >
            <InputNumber
              style={{ width: '100%' }}
              min={0}
              placeholder="请输入库存"
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
