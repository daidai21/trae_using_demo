import React from 'react'
import ReactDOM from 'react-dom/client'
import { BrowserRouter } from 'react-router-dom'
import { ConfigProvider } from 'antd'
import zhCN from 'antd/locale/zh_CN'
import enUS from 'antd/locale/en_US'
import idID from 'antd/locale/id_ID'
import { useTranslation } from 'react-i18next'
import { AuthProvider } from '@ecommerce/ui'
import '@ecommerce/ui/i18n'
import App from './App'
import './index.css'

const AntdLocaleProvider = ({ children }: { children: React.ReactNode }) => {
  const { i18n } = useTranslation()
  
  const getAntdLocale = () => {
    switch (i18n.language) {
      case 'en-US':
        return enUS
      case 'id-ID':
        return idID
      default:
        return zhCN
    }
  }
  
  return <ConfigProvider locale={getAntdLocale()}>{children}</ConfigProvider>
}

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <BrowserRouter>
      <AntdLocaleProvider>
        <AuthProvider>
          <App />
        </AuthProvider>
      </AntdLocaleProvider>
    </BrowserRouter>
  </React.StrictMode>,
)
