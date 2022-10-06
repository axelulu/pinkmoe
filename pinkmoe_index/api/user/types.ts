/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-18 21:44:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 17:42:03
 * @FilePath: /pinkmoe_vitesse/src/api/user/types.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
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
