/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-25 16:42:37
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-25 16:42:40
 * @FilePath: /pinkmoe_index/composables/languages.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import { useCookie, useState } from 'nuxt/app'
// import { useState, useCookie } from '#app' 这两种方式都可以导入useState,方法
// 示例：（使用命名导出）
export const userLangeages = () => {
  const langages = useCookie('lang')
  return useState('userLang', () => langages.value ?? 'cn')
}
