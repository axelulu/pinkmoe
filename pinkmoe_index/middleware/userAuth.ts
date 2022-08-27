/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-25 14:11:32
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-27 11:34:13
 * @FilePath: /pinkmoe_index/middleware/userAuth.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
/*!
 * Router Global middleware
 */

import { isLogin } from '../utils/auth'

export default defineNuxtRouteMiddleware(async (to, _from) => {
  if (!isLogin())
    return navigateTo('/')
})
