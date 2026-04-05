import React from 'react'
import { Layout, Menu, Button, Space } from 'antd'
import { useNavigate, useLocation } from 'react-router-dom'
import {
  HomeOutlined,
  ShoppingCartOutlined,
  ShoppingOutlined,
  ShopOutlined,
  UserOutlined,
  LogoutOutlined,
  LoginOutlined,
  FileTextOutlined,
} from '@ant-design/icons'
import { useAuth } from '../store/AuthContext'

const { Header, Content, Footer } = Layout

const AppLayout: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const navigate = useNavigate()
  const location = useLocation()
  const { user, logout, isAuthenticated } = useAuth()

  const menuItems = [
    {
      key: '/',
      icon: <HomeOutlined />,
      label: '商品列表',
    },
    {
      key: '/merchants',
      icon: <ShopOutlined />,
      label: '商家列表',
    },
  ]

  if (isAuthenticated) {
    menuItems.push(
      {
        key: '/cart',
        icon: <ShoppingCartOutlined />,
        label: '购物车',
      },
      {
        key: '/orders',
        icon: <FileTextOutlined />,
        label: '我的订单',
      },
      {
        key: '/merchant/form',
        icon: <ShoppingOutlined />,
        label: '商家管理',
      },
      {
        key: '/product/form',
        icon: <ShoppingOutlined />,
        label: '发布商品',
      }
    )
  }

  return (
    <Layout className="min-h-screen">
      <Header style={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between' }}>
        <div style={{ display: 'flex', alignItems: 'center' }}>
          <div style={{ color: 'white', fontSize: '20px', fontWeight: 'bold', marginRight: '24px' }}>
            电商平台
          </div>
          <Menu
            theme="dark"
            mode="horizontal"
            selectedKeys={[location.pathname]}
            items={menuItems}
            onClick={({ key }) => navigate(key)}
            style={{ flex: 1, minWidth: 0 }}
          />
        </div>
        <Space>
          {isAuthenticated ? (
            <>
              <span style={{ color: 'white' }}>欢迎，{user?.username}</span>
              <Button
                type="text"
                icon={<LogoutOutlined />}
                onClick={() => {
                  logout()
                  navigate('/')
                }}
                style={{ color: 'white' }}
              >
                退出
              </Button>
            </>
          ) : (
            <>
              <Button
                type="text"
                icon={<LoginOutlined />}
                onClick={() => navigate('/login')}
                style={{ color: 'white' }}
              >
                登录
              </Button>
              <Button
                type="text"
                icon={<UserOutlined />}
                onClick={() => navigate('/register')}
                style={{ color: 'white' }}
              >
                注册
              </Button>
            </>
          )}
        </Space>
      </Header>
      <Content style={{ padding: '24px 50px', minHeight: 'calc(100vh - 128px)' }}>
        {children}
      </Content>
      <Footer style={{ textAlign: 'center' }}>
        电商平台 ©{new Date().getFullYear()} Created with React + Ant Design
      </Footer>
    </Layout>
  )
}

export default AppLayout
