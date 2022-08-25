/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-25 13:55:58
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-25 14:54:05
 * @FilePath: /pinkmoe_index/plugins/nprogress.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
/*!
 * image lazyload
 */
import NProgress from 'nprogress'

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.hook('page:start', () => {
    NProgress.start()
  })

  nuxtApp.hook('page:finish', () => {
    NProgress.end()
  })
})
