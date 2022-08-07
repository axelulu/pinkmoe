/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-04-21 22:34:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:20:52
 * @FilePath: /pinkmoe_admin/src/store/modules/lockscreen.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { defineStore } from 'pinia';
import { IS_LOCKSCREEN } from '@/store/mutation-types';
import { storage } from '@/utils/Storage';

// 长时间不操作默认锁屏时间
const initTime = 60 * 60;

const isLock = storage.get(IS_LOCKSCREEN, false);

export type ILockscreenState = {
  isLock: boolean; // 是否锁屏
  lockTime: number;
};

export const useLockscreenStore = defineStore({
  id: 'app-lockscreen',
  state: (): ILockscreenState => ({
    isLock: isLock === true, // 是否锁屏
    lockTime: isLock == 'true' ? initTime : 0,
  }),
  getters: {},
  actions: {
    setLock(payload) {
      this.isLock = payload;
      storage.set(IS_LOCKSCREEN, this.isLock);
    },
    setLockTime(payload = initTime) {
      this.lockTime = payload;
    },
  },
});
