/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-04-21 22:34:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:18:31
 * @FilePath: /pinkmoe_admin/src/hooks/setting/useProjectSetting.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { computed } from 'vue';
import { useProjectSettingStore } from '@/store/modules/projectSetting';

export function useProjectSetting() {
  const projectStore = useProjectSettingStore();

  const getNavMode = computed(() => projectStore.navMode);

  const getNavTheme = computed(() => projectStore.navTheme);

  const getIsMobile = computed(() => projectStore.isMobile);

  const getHeaderSetting = computed(() => projectStore.headerSetting);

  const getMultiTabsSetting = computed(() => projectStore.multiTabsSetting);

  const getMenuSetting = computed(() => projectStore.menuSetting);

  const getCrumbsSetting = computed(() => projectStore.crumbsSetting);

  const getPermissionMode = computed(() => projectStore.permissionMode);

  const getShowFooter = computed(() => projectStore.showFooter);

  const getIsPageAnimate = computed(() => projectStore.isPageAnimate);

  const getPageAnimateType = computed(() => projectStore.pageAnimateType);

  return {
    getNavMode,
    getNavTheme,
    getIsMobile,
    getHeaderSetting,
    getMultiTabsSetting,
    getMenuSetting,
    getCrumbsSetting,
    getPermissionMode,
    getShowFooter,
    getIsPageAnimate,
    getPageAnimateType,
  };
}
