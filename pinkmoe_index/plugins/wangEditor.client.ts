/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-25 13:55:58
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-09-10 23:29:36
 * @FilePath: /pinkmoe_index/plugins/wangEditor.client.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
/*!
 * image lazyload
 */
import { Editor, Toolbar } from '@wangeditor/editor-for-vue'
import '@wangeditor/editor/dist/css/style.css'
export default defineNuxtPlugin((nuxtApp) => {
    nuxtApp.vueApp.component('Editor', Editor)
    nuxtApp.vueApp.component('Toolbar', Toolbar)
})
