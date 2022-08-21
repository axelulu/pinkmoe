/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-31 12:03:31
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-21 10:27:28
 * @FilePath: /pinkmoe_vitesse/src/components/GlobalNotice/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */

/**
 * 通知插件
 * 全局注册组件
 * vue3 已经去掉 整体导出 Vue，这样无法将方法直接挂载在Vue的原型上了，而是通过config注册
 * 使用 组件获取instance.proxy，const { proxy } = getCurrentInstance();
 * proxy.$notify({ content: 'test' });
 * 不要使用 ctx，生产环境不支持
 */
import type { App, Plugin } from 'vue'
import GlobalNotice from './index.vue'
import notify from './func'

// 挂载组件方法
const install = (app: App): App => {
  app.config.globalProperties.$notify = notify
  app.component(GlobalNotice.name, GlobalNotice)
  return app
}

export default install as Plugin
