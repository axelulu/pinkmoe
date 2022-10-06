/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-29 11:13:55
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-21 10:27:46
 * @FilePath: /pinkmoe_vitesse/src/components/Vipconfirm/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
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
