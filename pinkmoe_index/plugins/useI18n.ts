/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-25 13:55:58
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-25 16:45:27
 * @FilePath: /pinkmoe_index/plugins/useI18n.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
/*!
 * image lazyload
 */
import { createI18n } from 'vue-i18n'
import jp from '../locales/jp'
import cn from '../locales/cn'

const messages = {
  cn,
  jp,
}

export default defineNuxtPlugin((nuxtApp) => {
  // 获取cookie
  const langages = useCookie('lang')
  const i18n = createI18n({
    legacy: false,
    globalInjection: true,
    fallbackLocale: 'cn',
    warnHtmlMessage: false,
    locale: langages.value || 'cn',
    messages,
  })

  nuxtApp.vueApp.use(i18n)
})
