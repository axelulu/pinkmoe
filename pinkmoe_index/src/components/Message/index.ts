/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-24 12:36:06
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:10:39
 * @FilePath: /pinkmoe_index/src/components/Message/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import Message from './index.vue';
import { createModal } from '/@/utils/createModal';

function MessageDia(options = {}) {
  const res = createModal(Message, options);
  res.app.component('FontAwesomeIcon', FontAwesomeIcon); // 渲染到创建的div节点上
  const $inst = res.app.mount(res.container);
  return $inst;
}
MessageDia.install = (app) => {
  app.component('MessageDia', Message);
  app.config.globalProperties.$message = MessageDia;
  // app.provide('subDialog', V3Popup)
};
export default MessageDia;
