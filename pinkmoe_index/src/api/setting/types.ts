/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-20 20:59:50
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 17:41:46
 * @FilePath: /pinkmoe_vitesse/src/api/setting/types.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import type { BaseResponseType } from '/@/api/common/types'

export interface ResSiteSetting {
  basic?: ResSetting<any>
  footer?: ResSetting<any>
  user_level?: ResSetting<any>
  user_shop?: ResSetting<any>
  user_search?: ResSetting<any>
  user_reward?: ResSetting<any>
}

export interface ResSetting<T> extends BaseResponseType {
  ID: number
  name?: string
  slug?: string
  value?: T
}

export interface ResBasic {
  title?: string
  keyword?: string
  desc?: string
  color?: string
  icon?: Array<string>
  background?: Array<string>
  logo?: Array<string>
  createTime?: string
}

export interface ResFooter {
  links?: Array<ResLinks>
  notice?: string
  contact?: string
  about?: string
  friendsLinks?: Array<ResLinks>
}

export interface ResLinks {
  name?: string
  url?: string
}
