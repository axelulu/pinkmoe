/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-18 21:44:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:04:27
 * @FilePath: /pinkmoe_index/src/router/root.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
export default [
  {
    path: '/',
    component: () => import('@/pages/login/index.vue'),
  },
  {
    path: 'login',
    component: () => import('@/pages/login/index.vue'),
  },
];
