/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-04-21 22:34:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:18:29
 * @FilePath: /pinkmoe_admin/src/hooks/setting/useDesignSetting.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { computed } from 'vue';
import { useDesignSettingStore } from '@/store/modules/designSetting';

export function useDesignSetting() {
  const designStore = useDesignSettingStore();

  const getDarkTheme = computed(() => designStore.darkTheme);

  const getAppTheme = computed(() => designStore.appTheme);

  const getAppThemeList = computed(() => designStore.appThemeList);

  return {
    getDarkTheme,
    getAppTheme,
    getAppThemeList,
  };
}
