/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-04 07:40:06
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:11:27
 * @FilePath: /pinkmoe_index/src/components/Shopconfirm/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import ShopConfirm from './index.vue';
import { createModal } from '/@/utils/createModal';
import Message from '/@/components/Message/index';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

function ShopConfirmDia(options = {}) {
  const res = createModal(ShopConfirm, options);
  res.app.use(Message);
  res.app.component('FontAwesomeIcon', FontAwesomeIcon); // 渲染到创建的div节点上
  const $inst = res.app.mount(res.container);
  return $inst;
}
ShopConfirmDia.install = (app) => {
  app.component('ShopConfirmDia', ShopConfirm);
  app.config.globalProperties.$shopConfirm = ShopConfirmDia;
};
export default ShopConfirmDia;
