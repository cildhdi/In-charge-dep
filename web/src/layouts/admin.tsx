import React from 'react';
import Link from 'umi/link';

import { Layout, Menu, Result } from 'antd';
const { Header, Footer, Content } = Layout;

const menus: string[] = ['merchant-manage', 'customer-manage'];

const AdminLayout: React.FC = props => {
  let index = menus.findIndex((value) => location.pathname.endsWith(value));
  if (index == -1) {
    return (<Result
      status="404"
      title="404"
      subTitle="页面不存在"
      extra={<Link to={menus[0]}>回到首页</Link>}
    />);
  }
  return (
    <Layout style={{ height: '100%' }}>
      <Header style={{ position: 'fixed', zIndex: 1, width: '100%' }}>
        <Menu
          theme='dark'
          mode='horizontal'
          defaultSelectedKeys={[menus[0]]}
          selectedKeys={[menus[index]]}
          style={{ lineHeight: '64px' }}
        >
          <Menu.Item key={menus[0]}><Link to={menus[0]}>商户管理</Link></Menu.Item>
          <Menu.Item key={menus[1]}><Link to={menus[1]}>用户管理</Link></Menu.Item>
        </Menu >
      </Header >
      <Content style={{ padding: '50px', marginTop: 64, minHeight: '500px' }}>
        <div style={{ background: '#fff', padding: 24, height: '100%' }}>{props.children}</div>
      </Content>
      <Footer style={{ textAlign: 'center' }}></Footer>
    </Layout >
  );
};

export default AdminLayout;
