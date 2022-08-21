/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-08 16:34:38
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-21 10:14:01
 * @FilePath: /pinkmoe_vitesse/src/modules/aplayer.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import { type UserModule } from '/@/types'

export const install: UserModule = async ({ app, isClient }) => {
  if (isClient) {
    const VueAPlayer = await import('vue3-aplayer')
    app.component('Aplayer', VueAPlayer)
  }
}
