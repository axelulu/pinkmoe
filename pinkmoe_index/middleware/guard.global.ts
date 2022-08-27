/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-25 14:11:32
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-27 18:21:11
 * @FilePath: /pinkmoe_index/middleware/guard.global.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
/*!
 * Router Global middleware
 */
import { useAppStore } from '../store/modules/app'
import { useUserStore } from '../store/modules/user'
import { isLogin } from '../utils/auth'

export default defineNuxtRouteMiddleware(async (_to, _from) => {
  const { settings, siteBasic, siteFooter, userLevel } = useAppStore()
  if (!siteBasic || !siteFooter || !userLevel)
    await settings()

  const { info, userInfo, isLogin: login, setLogin } = useUserStore()
  if (login === undefined)
    setLogin(isLogin() as boolean)

  if (_to.meta.auth && !login)
    return navigateTo('/')

  if (!userInfo && login)
    await info()
})
