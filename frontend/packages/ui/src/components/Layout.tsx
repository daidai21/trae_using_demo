import React from 'react'
import { Layout as AntLayout, Menu } from 'antd'
import { Link, useNavigate, useLocation } from 'react-router-dom'
import {
  ShopOutlined,
  ShoppingCartOutlined,
  ShoppingOutlined,
  UserOutlined,
  LogoutOutlined
} from '@ant-design/icons'
import { useAuth } from '../store'

const { Header, Sider, Content } = AntLayout

interface LayoutProps {
  children: React.ReactNode
}

export function Layout({ children }: LayoutProps) {
  const { user, logout, isAuthenticated } = useAuth()
  const navigate = useNavigate()
  const location = useLocation()

  const menuItems = [
    {
      key: '/products',
      icon: <ShopOutlined />,
      label: <Link to="/products">商品列表</Link>
    },
    {
      key: '/cart',
      icon: <ShoppingCartOutlined />,
      label: <Link to="/cart">购物车</Link>
    },
    {
      key: '/orders',
      icon: <ShoppingOutlined />,
      label: <Link to="/orders">我的订单</Link>
    },
    {
      key: '/merchants',
      icon: <UserOutlined />,
      label: <Link to="/merchants">商家管理</Link>
    }
  ]

  const handleLogout = () => {
    logout()
    navigate('/login')
  }

  return (
    <AntLayout style={{ minHeight: '100vh' }}>
      {isAuthenticated && (
        <Sider theme="dark" width={200}>
          <div style={{ height: 64, display: 'flex', alignItems: 'center', justifyContent: 'center', color: 'white', fontSize: '18px', fontWeight: 'bold' }}>
            电商平台
          </div>
          <Menu
            theme="dark"
            mode="inline"
            selectedKeys={[location.pathname]}
            items={menuItems}
          />
        </Sider>
      )}
      <AntLayout>
        <Header style={{ padding: '0 24px', background: '#fff', display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
          <h2 style={{ margin: 0 }}>电商交易平台</h2>
          {isAuthenticated && (
            <div style={{ display: 'flex', alignItems: 'center', gap: 16 }}>
              <span>欢迎, {user?.username}</span>
              <a onClick={handleLogout} style={{ cursor: 'pointer' }}>
                <LogoutOutlined /> 退出
              </a>
            </div>
          )}
        </Header>
        <Content style={{ margin: '24px', padding: 24, background: '#fff', minHeight: 280 }}>
          {children}
        </Content>
      </AntLayout>
    </AntLayout>
  )
}
