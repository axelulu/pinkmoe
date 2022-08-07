/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-04-21 22:34:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:18:49
 * @FilePath: /pinkmoe_admin/src/hooks/useDomWidth.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { ref, onMounted, onUnmounted } from 'vue';
import { debounce } from 'lodash';

/**
 * description: 获取页面宽度
 */

export function useDomWidth() {
  const domWidth = ref(window.innerWidth);

  function resize() {
    domWidth.value = document.body.clientWidth;
  }

  onMounted(() => {
    window.addEventListener('resize', debounce(resize, 80));
  });
  onUnmounted(() => {
    window.removeEventListener('resize', resize);
  });

  return domWidth;
}
