/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-31 12:03:31
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-21 10:27:28
 * @FilePath: /pinkmoe_vitesse/src/components/GlobalNotice/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
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
