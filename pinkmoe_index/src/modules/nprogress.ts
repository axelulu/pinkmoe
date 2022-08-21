/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-08 16:34:38
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 16:02:37
 * @FilePath: /pinkmoe_vitesse/src/modules/nprogress.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import NProgress from 'nprogress'
import { type UserModule } from '/@/types'
import { useAppStore, useUserStore } from '/@/store'

export const install: UserModule = ({ isClient, router }) => {
  if (isClient) {
    router.beforeEach(async (_to, _from) => {
      if (_to.path !== _from.path)
        NProgress.start()
      const { settings, siteBasic, siteFooter, userLevel } = useAppStore()
      if (!siteBasic || !siteFooter || !userLevel)
        await settings()

      const { info, userInfo, isLogin } = useUserStore()
      // if (_to.meta.auth && !isLogin)
      //   next('/')

      if (!userInfo && isLogin)
        await info()
    })
    router.afterEach(() => { NProgress.done() })
  }
}
