/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-18 21:44:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 17:42:03
 * @FilePath: /pinkmoe_vitesse/src/api/user/types.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import type { ResAuthor } from '/@/api/home/types'

export interface ReqParams {
  username: string
  password: string
}

export interface ReqAuth {
  auths: string[]
  modules: string[]
  is_admin?: 0 | 1
}

export interface ReqUserDetail {
  nickName?: string
  desc?: string
}

export interface ReqUserPwd {
  oldPassword?: string
  password?: string
  reNewPwd?: string
}

export interface ReqUserEmail {
  emailCaptcha?: string
  emailCode?: string
}

export interface ReqUserEmailCaptcha {
  emailType?: string
  emailCode?: string
}

export interface ReqUserAvatar {
  avatar?: string
}

export interface ResResult {
  data?: ResResultData
  status: string | ''
  headers: object
}

export interface ResCaptcha {
  captchaImg?: string
  captchaId?: string
}

export interface ResLogin<T> {
  token: string
  userInfo: T
}

export interface ResCheckIn {
  credit?: string
}

export interface ResUserProfile {
  userInfo: ResAuthor
  userMenu: any
}

export interface ResResultData {
  code?: number
  result?: any
  message: string
  status: string
}

export interface ReqUserBuyVip {
  authorityId?: string
  authorityParamId?: string
  authorityParamKey?: string
  authorityParamValue?: number
  authorityParamDay?: number
  authorityNum?: number
}

export interface ReqUserBuyShop {
  shopType?: string
  priceType?: string
  shopCredit?: number
  shopValue?: number
  shopNum?: number
  shopKey?: string
}
