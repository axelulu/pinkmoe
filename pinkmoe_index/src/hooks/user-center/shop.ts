/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-04 07:27:27
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 08:11:53
 * @FilePath: /pinkmoe_index/src/hooks/user-center/shop.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import { useAppStore, useUserStore } from '/@/store';

export const useUserCenterShop = () => {
  const loading = ref(true);
  const currentIndex = ref(0);
  const auth = useUserStore();
  const { proxy } = getCurrentInstance();
  const { userShop } = useAppStore();

  function buyShop() {
    proxy.$shopConfirm({
      authorShop: userShop[currentIndex.value],
    });
  }

  onMounted(() => {
    loading.value = true;
    setTimeout(() => {
      loading.value = false;
    }, 300);
  });

  return {
    loading,
    currentIndex,
    auth,
    userShop,
    buyShop,
  };
};
