/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-29 19:46:40
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 18:09:50
 * @FilePath: /pinkmoe_vitesse/src/store/modules/socket/types.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
export interface AppState {
  theme: string
  colorWeek: boolean
  navbar: boolean
  menu: boolean
  menuCollapse: boolean
  footer: boolean
  themeColor: string
  menuWidth: number
  globalSettings: boolean
  [key: string]: unknown
}
