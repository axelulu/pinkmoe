/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-18 21:44:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:04:25
 * @FilePath: /pinkmoe_index/src/router/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { createRouter, createWebHistory } from 'vue-router';
import routes from 'virtual:generated-pages';
import NProgress from 'nprogress';
import 'nprogress/nprogress.css';
import { useAppStore, useUserStore } from '/@/store';

routes.push({
  path: '/',
  redirect: '/login',
});
//导入生成的路由数据
const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach(async (_to, _from, next) => {
  NProgress.start();
  const { settings, siteBasic, siteFooter } = useAppStore();
  if (!siteBasic || !siteFooter) {
    await settings();
  }
  const { info, userInfo } = useUserStore();
  if (!userInfo) {
    await info();
  }
  next();
});

router.afterEach((_to) => {
  NProgress.done();
});

export default router;
