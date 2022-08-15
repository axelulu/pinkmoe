/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-18 21:44:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-16 00:14:23
 * @FilePath: /pinkmoe_index/src/store/modules/user/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import { defineStore } from 'pinia';
import {
  login as userLogin,
  logout as userLogout,
  getUserProfile,
  LoginData,
  RegData,
  reg,
  forget,
  checkInStatus,
} from '/@/api/user/index';
import { setToken, clearToken, isLogin } from '/@/utils/auth';
import { UserState } from '/@/store/modules/user/types';
import { ResAuthor } from '/@/api/home/types';
import { useSocketStore } from '../socket';

export const useUserStore = defineStore('user', {
  state: (): UserState => ({
    userInfo: <any>undefined,
    isLogin: isLogin(),
    checkInStatus: false,
  }),
  getters: {
    userProfile(state: UserState): UserState {
      return { ...state };
    },
  },
  actions: {
    async checkIn() {
      const { result } = await checkInStatus();
      this.checkInStatus = result;
    },
    // 设置用户的信息
    setInfo(userInfo: ResAuthor) {
      this.userInfo = userInfo;
    },
    // 设置用户的信息
    setLogin(login: boolean) {
      this.isLogin = login;
    },
    // 重置用户信息
    resetInfo() {
      this.$reset();
    },
    // 获取用户信息
    async info() {
      const result = await getUserProfile();
      this.setInfo(result?.userInfo);
    },
    // 异步登录并存储token
    async login(loginForm: LoginData) {
      const { code, result, message } = await userLogin(loginForm);
      const token = result?.token;
      if (code === 200 && result?.token) {
        setToken(token);
        this.setLogin(true);
      }
      await this.checkIn();
      return { result, code, message };
    },
    async reg(regForm: RegData) {
      const { code, result, message } = await reg(regForm);
      return { result, code, message };
    },
    async forget(regForm: RegData) {
      const { code, result, message } = await forget(regForm);
      return { result, code, message };
    },
    // Logout
    async logout() {
      this.resetInfo();
      this.setLogin(false);
      clearToken();
      await userLogout();
      const { closeSocket } = useSocketStore();
      closeSocket();
    },
  },
});
