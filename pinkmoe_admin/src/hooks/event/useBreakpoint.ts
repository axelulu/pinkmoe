/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-04-21 22:34:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:18:17
 * @FilePath: /pinkmoe_admin/src/hooks/event/useBreakpoint.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { ref, computed, ComputedRef, unref } from 'vue';
import { useEventListener } from '@/hooks/event/useEventListener';
import { screenMap, sizeEnum, screenEnum } from '@/enums/breakpointEnum';

let globalScreenRef: ComputedRef<sizeEnum | undefined>;
let globalWidthRef: ComputedRef<number>;
let globalRealWidthRef: ComputedRef<number>;

export interface CreateCallbackParams {
  screen: ComputedRef<sizeEnum | undefined>;
  width: ComputedRef<number>;
  realWidth: ComputedRef<number>;
  screenEnum: typeof screenEnum;
  screenMap: Map<sizeEnum, number>;
  sizeEnum: typeof sizeEnum;
}

export function useBreakpoint() {
  return {
    screenRef: computed(() => unref(globalScreenRef)),
    widthRef: globalWidthRef,
    screenEnum,
    realWidthRef: globalRealWidthRef,
  };
}

// Just call it once
export function createBreakpointListen(fn?: (opt: CreateCallbackParams) => void) {
  const screenRef = ref<sizeEnum>(sizeEnum.XL);
  const realWidthRef = ref(window.innerWidth);

  function getWindowWidth() {
    const width = document.body.clientWidth;
    const xs = screenMap.get(sizeEnum.XS)!;
    const sm = screenMap.get(sizeEnum.SM)!;
    const md = screenMap.get(sizeEnum.MD)!;
    const lg = screenMap.get(sizeEnum.LG)!;
    const xl = screenMap.get(sizeEnum.XL)!;
    if (width < xs) {
      screenRef.value = sizeEnum.XS;
    } else if (width < sm) {
      screenRef.value = sizeEnum.SM;
    } else if (width < md) {
      screenRef.value = sizeEnum.MD;
    } else if (width < lg) {
      screenRef.value = sizeEnum.LG;
    } else if (width < xl) {
      screenRef.value = sizeEnum.XL;
    } else {
      screenRef.value = sizeEnum.XXL;
    }
    realWidthRef.value = width;
  }

  useEventListener({
    el: window,
    name: 'resize',

    listener: () => {
      getWindowWidth();
      resizeFn();
    },
    // wait: 100,
  });

  getWindowWidth();
  globalScreenRef = computed(() => unref(screenRef));
  globalWidthRef = computed((): number => screenMap.get(unref(screenRef)!)!);
  globalRealWidthRef = computed((): number => unref(realWidthRef));

  function resizeFn() {
    fn?.({
      screen: globalScreenRef,
      width: globalWidthRef,
      realWidth: globalRealWidthRef,
      screenEnum,
      screenMap,
      sizeEnum,
    });
  }

  resizeFn();
  return {
    screenRef: globalScreenRef,
    screenEnum,
    widthRef: globalWidthRef,
    realWidthRef: globalRealWidthRef,
  };
}
