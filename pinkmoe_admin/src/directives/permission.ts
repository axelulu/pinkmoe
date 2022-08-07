/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-04-21 22:34:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:17:56
 * @FilePath: /pinkmoe_admin/src/directives/permission.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { ObjectDirective } from 'vue';
import { usePermission } from '@/hooks/web/usePermission';

export const permission: ObjectDirective = {
  mounted(el: HTMLButtonElement, binding) {
    if (binding.value == undefined) return;
    const { action, effect } = binding.value;
    const { hasPermission } = usePermission();
    if (!hasPermission(action)) {
      if (effect == 'disabled') {
        el.disabled = true;
        el.style['disabled'] = 'disabled';
        el.classList.add('n-button--disabled');
      } else {
        el.remove();
      }
    }
  },
};
