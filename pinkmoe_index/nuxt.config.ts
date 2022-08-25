/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-25 13:53:10
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-25 20:13:18
 * @FilePath: /pinkmoe_index/nuxt.config.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import path from 'path'
import { defineNuxtConfig } from 'nuxt'

function pathResolve(dir: string) {
  return path.resolve(process.cwd(), '.', dir)
}

export default defineNuxtConfig({
  css: [
    'nprogress/nprogress.css',
    'uno.css',
    '/@/assets/styles/main.css',
    '/@/assets/styles/client.css',
  ],
  alias: {
    '/@/': `${pathResolve('.')}/`,
  },
  modules: [
    '@vueuse/nuxt',
    '@unocss/nuxt',
    '@pinia/nuxt',
    '@nuxtjs/color-mode',
  ],
  experimental: {
    reactivityTransform: true,
    viteNode: false,
  },
  unocss: {
    preflight: true,
  },
  colorMode: {
    classSuffix: '',
  },
  // https://github.com/nuxt/framework/issues/6204#issuecomment-1201398080
  hooks: {
    'vite:extendConfig': function (config: any, { isServer }: any) {
      if (isServer) {
        // Workaround for netlify issue
        // https://github.com/nuxt/framework/issues/6204
        config.build.rollupOptions.output.inlineDynamicImports = true
      }
    },
  },
})
