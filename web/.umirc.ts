import { IConfig } from 'umi-types';

// ref: https://umijs.org/config/
const config: IConfig = {
  treeShaking: true,
  routes: [
    {
      path: '/',
      component: '../pages/index',
    },
    {
      path: '/admin',
      component: '../layouts/admin',
      routes: [
        {
          path: '/admin/merchant-manage',
          component: '../pages/admin/merchant-manage'
        }
      ]
    }
  ],
  plugins: [
    // ref: https://umijs.org/plugin/umi-plugin-react.html
    ['umi-plugin-react', {
      antd: true,
      dva: true,
      dynamicImport: false,
      title: 'web',
      dll: false,

      routes: {
        exclude: [
          /models\//,
          /services\//,
          /model\.(t|j)sx?$/,
          /service\.(t|j)sx?$/,
          /components\//,
        ],
      },
    }],
  ],
  proxy: {
    "/api": {
      "target": "http://localhost:8080",
      "changeOrigin": true,
    }
  }
}

export default config;
