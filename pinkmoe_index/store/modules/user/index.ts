/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-18 21:44:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-09-10 15:21:05
 * @FilePath: /pinkmoe_index/store/modules/user/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
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
