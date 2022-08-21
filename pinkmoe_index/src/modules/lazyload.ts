/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-19 15:50:43
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 15:52:22
 * @FilePath: /pinkmoe_vitesse/src/modules/lazyload.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import VueLazyLoad from 'vue3-lazyload'
import type { UserModule } from '../types'

export const install: UserModule = ({ app }) => {
  app.use(VueLazyLoad, {
    preLoad: 1.3, // 距离当前dom距离页面底部的高度
    error: 'data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7', // 加载失败显示的图片
    loading: 'data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7', // 加载中显示的图片
    attempt: 1, // 图片加载失败，最多重试的次数
  })
}
