import React, { useState, useEffect } from 'react'
import { Form, Input, Button, Card, message } from 'antd'
import { useNavigate } from 'react-router-dom'
import { merchantApi } from '../services/api'
import { Merchant } from '../types'

const MerchantForm: React.FC = () => {
  const [form] = Form.useForm()
  const [loading, setLoading] = useState(false)
  const [merchant, setMerchant] = useState<Merchant | null>(null)
  const navigate = useNavigate()

  const fetchMyMerchant = async () => {
    setLoading(true)
    try {
      const res = await merchantApi.getMyMerchant()
      setMerchant(res.data)
      form.setFieldsValue(res.data)
    } catch (error) {
      console.error('获取商家信息失败:', error)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    fetchMyMerchant()
  }, [])

  const onFinish = async (values: { name: string; description: string }) => {
    setLoading(true)
    try {
      if (merchant) {
        await merchantApi.update(merchant.id, values)
        message.success('商家信息更新成功')
      } else {
        await merchantApi.create(values)
        message.success('商家入驻成功')
      }
      fetchMyMerchant()
    } catch (error) {
      console.error('提交失败:', error)
    } finally {
      setLoading(false)
    }
  }

  return (
    <div>
      <Card title={merchant ? '编辑商家信息' : '商家入驻'}>
        <Form
          form={form}
          layout="vertical"
          onFinish={onFinish}
          initialValues={{ description: '' }}
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
          >
            <Input.TextArea rows={4} placeholder="请输入商家描述" />
          </Form.Item>

          <Form.Item>
            <Button type="primary" htmlType="submit" loading={loading}>
              {merchant ? '更新' : '提交入驻申请'}
            </Button>
          </Form.Item>
        </Form>
      </Card>
    </div>
  )
}

export default MerchantForm
