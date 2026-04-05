import { Dropdown, MenuProps, Space } from 'antd';
import { GlobalOutlined } from '@ant-design/icons';
import { useTranslation } from 'react-i18next';

const LanguageSwitcher = () => {
  const { i18n, t } = useTranslation();

  const items: MenuProps['items'] = [
    {
      key: 'zh-CN',
      label: t('language.zhCN'),
      onClick: () => i18n.changeLanguage('zh-CN'),
    },
    {
      key: 'en-US',
      label: t('language.enUS'),
      onClick: () => i18n.changeLanguage('en-US'),
    },
    {
      key: 'id-ID',
      label: t('language.idID'),
      onClick: () => i18n.changeLanguage('id-ID'),
    },
  ];

  return (
    <Dropdown menu={{ items }} placement="bottomRight">
      <Space>
        <GlobalOutlined />
        {t('language.select')}
      </Space>
    </Dropdown>
  );
};

export default LanguageSwitcher;
