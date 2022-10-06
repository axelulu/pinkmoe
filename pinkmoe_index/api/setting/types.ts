/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-20 20:59:50
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 17:41:46
 * @FilePath: /pinkmoe_vitesse/src/api/setting/types.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
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
