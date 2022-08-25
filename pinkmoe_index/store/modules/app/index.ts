/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-18 21:44:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-25 19:14:52
 * @FilePath: /pinkmoe_index/store/modules/app/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import { defineStore } from 'pinia'
import piniaStore from '/@/store/index'
import type { AppState } from './types'
import { getSiteSetting } from '/@/api/setting'
import type { ResSiteSetting } from '/@/api/setting/types'

export const useAppStore = defineStore(
  // 唯一ID
  'app',
  {
    state: () => {
      return {
        siteBasic: <any>null,
        siteFooter: <any>null,
        userLevel: <any>null,
        userShop: <any>null,
        userSearch: <any>null,
        userReward: <any>null,
        theme: 'dark',
      }
    },
    getters: {},
    actions: {
      // 获取用户信息
      async settings() {
        const siteSetting = await getSiteSetting()
        this.setInfo(siteSetting)
      },
      // 设置用户的信息
      setInfo(siteSetting: ResSiteSetting) {
        this.siteBasic = JSON.parse(siteSetting.basic?.value)
        this.siteFooter = JSON.parse(siteSetting.footer?.value)
        this.userLevel = JSON.parse(siteSetting.user_level?.value)
        this.userShop = JSON.parse(siteSetting.user_shop?.value)
        this.userSearch = JSON.parse(siteSetting.user_search?.value)
        this.userReward = JSON.parse(siteSetting.user_reward?.value)
      },
      // Update app settings
      updateSettings(partial: Partial<AppState>) {
        this.$patch(partial)
      },

      // Change theme color
      toggleTheme(dark: boolean) {
        if (process.client) {
          if (dark) {
            this.theme = 'dark'
            document.documentElement.classList.add('dark')
            document.body.setAttribute('arco-theme', 'dark')
            if (process.client)
              localStorage.setItem('theme', this.theme)
          }
          else {
            this.theme = 'light'
            document.documentElement.classList.remove('dark')
            document.body.removeAttribute('arco-theme')
            if (process.client)
              localStorage.setItem('theme', this.theme)
          }
        }
      },
    },
  },
)

export function useAppOutsideStore() {
  return useAppStore(piniaStore)
}
