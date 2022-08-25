/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-25 13:55:58
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-25 14:01:15
 * @FilePath: /pinkmoe_index/plugins/dialog.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
/*!
 * image lazyload
 */
import Search from '/@/components/Search/index'
import Login from '/@/components/Login/index'
import VipConfirm from '/@/components/Vipconfirm/index'
import ShopConfirm from '/@/components/Shopconfirm/index'
import GlobalNotice from '/@/components/GlobalNotice/index'
import Message from '/@/components/Message/index'

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp
    .use(VipConfirm)
    .use(ShopConfirm)
    .use(GlobalNotice)
    .use(Search)
    .use(Login)
    .use(Message)
})
