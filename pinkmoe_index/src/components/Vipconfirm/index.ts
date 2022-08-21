/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-29 11:13:55
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-21 10:27:46
 * @FilePath: /pinkmoe_vitesse/src/components/Vipconfirm/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import VipConfirm from './index.vue'
import { createModal } from '/@/utils/createModal'
import Message from '/@/components/Message/index'

function VipConfirmDia(options = {}) {
  const res = createModal(VipConfirm, options)
  res.app.use(Message)
  const $inst = res.app.mount(res.container)
  return $inst
}
VipConfirmDia.install = (app) => {
  app.component('VipConfirmDia', VipConfirm)
  app.config.globalProperties.$vipConfirm = VipConfirmDia
}
export default VipConfirmDia
