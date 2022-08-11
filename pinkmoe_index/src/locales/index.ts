/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-26 09:37:01
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-11 15:10:40
 * @FilePath: /pinkmoe_index/src/locales/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import { createI18n } from 'vue-i18n';
import zh from './cn';
import jp from './jp';

const i18n = createI18n({
  legacy: false,
  globalInjection: true,
  // 如果本地有语言标识就采用本地，没有就根据浏览器语言默认标识显示语言
  locale: localStorage.getItem('locale') || navigator.language.slice(0, 2),
  messages: {
    zh,
    jp,
  },
});
export default i18n;
