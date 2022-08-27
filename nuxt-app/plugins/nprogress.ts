/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-25 13:55:58
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-27 08:21:43
 * @FilePath: /nuxt-app/plugins/nprogress.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
export default defineNuxtPlugin((nuxtApp) => {
    const bar = ref(null)
  
    nuxtApp.hook('app:mounted', (e) => {
      console.log(1111)
    })
    nuxtApp.hook('page:start', (e) => {
        if (process.client)
            console.log(222)
    })
    nuxtApp.hook('page:finish', (e) => {
        if (process.client)
            console.log(333)
    })
    nuxtApp.hook('app:error', (e) => {
      // 判断是否在客户端
      if (process.client)
        console.log(444)
    })
  })
  