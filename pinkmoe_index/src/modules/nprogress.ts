/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-08 16:34:38
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-22 22:57:39
 * @FilePath: /pinkmoe_index/src/modules/nprogress.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import NProgress from 'nprogress'
import { type UserModule } from '/@/types'
import { useAppStore, useUserStore } from '/@/store'
import { isLogin } from '../utils/auth'

export const install: UserModule = ({ isClient, router }) => {
  if (isClient) {
    router.beforeEach(async (_to, _from) => {
      if (_to.path !== _from.path)
        NProgress.start()
      const { settings, siteBasic, siteFooter, userLevel } = useAppStore()
      if (!siteBasic || !siteFooter || !userLevel)
        await settings()

      const { info, userInfo, isLogin: login, setLogin } = useUserStore()
      if (login === undefined)
        setLogin(isLogin() as boolean)

      if (_to.meta.auth && !login)
        router.push('/')

      if (!userInfo && login)
        await info()
    })
    router.afterEach(() => { NProgress.done() })
  }
}
