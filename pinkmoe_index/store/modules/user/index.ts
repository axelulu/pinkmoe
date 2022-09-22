/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-18 21:44:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-09-10 15:21:05
 * @FilePath: /pinkmoe_index/store/modules/user/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import { defineStore } from 'pinia'
import type {
  LoginData,
  RegData,
} from '/@/api/user/index'
import {
  checkInStatus,
  forget,
  reg,
  login as userLogin,
  logout as userLogout,
} from '/@/api/user/index'
import { clearToken, setToken } from '/@/utils/auth'
import type { UserState } from '/@/store/modules/user/types'
import type { ResAuthor } from '/@/api/home/types'
import { useSocketStore } from '../socket'

export const useUserStore = defineStore('user', {
  state: (): UserState => ({
    userInfo: <any>undefined,
    isLogin: undefined,
    checkInStatus: false,
  }),
  getters: {
    userProfile(state: UserState): UserState {
      return { ...state }
    },
  },
  actions: {
    setCheckIn(checkInStatus: any) {
      this.checkInStatus = checkInStatus
    },
    // 设置用户的信息
    setInfo(userInfo: ResAuthor) {
      this.userInfo = userInfo
    },
    // 设置用户的信息
    setLogin(login: boolean) {
      this.isLogin = login
    },
    // 重置用户信息
    resetInfo() {
      this.$reset()
    },
    // 异步登录并存储token
    async login(loginForm: LoginData) {
      const { code, result, message } = await userLogin(loginForm)
      const token = result?.token
      if (code === 200 && result?.token) {
        setToken(token)
        this.setLogin(true)
      }
      const { result: checkStatus } = await checkInStatus()
      this.setCheckIn(checkStatus)
      return { result, code, message }
    },
    async reg(regForm: RegData) {
      const { code, result, message } = await reg(regForm)
      return { result, code, message }
    },
    async forget(regForm: RegData) {
      const { code, result, message } = await forget(regForm)
      return { result, code, message }
    },
    // Logout
    async logout() {
      this.resetInfo()
      this.setLogin(false)
      clearToken()
      await userLogout()
      const { closeSocket } = useSocketStore()
      closeSocket()
    },
  },
})
