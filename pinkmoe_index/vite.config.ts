/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-25 15:53:12
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-25 19:24:10
 * @FilePath: /pinkmoe_index/vite.config.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import { defineConfig } from 'vite'
import { API_BASE_URL, API_TARGET_URL } from './constant'

export default defineConfig({
  define: {
    __VUE_I18N_FULL_INSTALL__: true,
    __VUE_I18N_LEGACY_API__: false,
    __INTLIFY_PROD_DEVTOOLS__: false,
  },
  server: {
    proxy: {
      [API_BASE_URL]: {
        target: API_TARGET_URL,
        changeOrigin: true,
        rewrite: path => path.replace(API_BASE_URL, ''),
      },
    },
  },
})
