/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-25 13:53:10
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-09-10 16:07:37
 * @FilePath: /pinkmoe_index/nuxt.config.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
import path from 'path'
import { defineNuxtConfig } from 'nuxt'

function pathResolve(dir: string) {
  return path.resolve(process.cwd(), '.', dir)
}

export default defineNuxtConfig({
  vite: {
    define: {
      __VUE_I18N_FULL_INSTALL__: true,
      __VUE_I18N_LEGACY_API__: false,
      __INTLIFY_PROD_DEVTOOLS__: false,
    },
    resolve: {
      extensions: ['.js', '.ts', '.mjs'],
    },
    server: {
      proxy: {
        '/uploads': {
          target: 'http://localhost:9527/uploads',
          changeOrigin: true,
          secure: false,
          rewrite: path => path.replace('/uploads', ''),
        },
      },
    },
  },
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
