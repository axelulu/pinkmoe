/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-19 15:50:43
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 16:13:13
 * @FilePath: /pinkmoe_vitesse/src/modules/dialog.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import Search from '/@/components/Search/index'
import Login from '/@/components/Login/index'
import VipConfirm from '/@/components/Vipconfirm/index'
import ShopConfirm from '/@/components/Shopconfirm/index'
import GlobalNotice from '/@/components/GlobalNotice/index'
import Message from '/@/components/Message/index'
import type { UserModule } from '../types'

export const install: UserModule = ({ app }) => {
  app
    .use(VipConfirm)
    .use(ShopConfirm)
    .use(GlobalNotice)
    .use(Search)
    .use(Login)
    .use(Message)
}
