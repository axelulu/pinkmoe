/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 13:57:54
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 17:40:54
 * @FilePath: /pinkmoe_vitesse/src/api/author/types.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import type { BaseResponseType, ReqPage } from '/@/api/common/types'
import type { ResAuthor } from '/@/api/home/types'

export interface ReqAuthorPost extends ReqPage {
  author?: string | string[]
}

export interface ReqAuthorFans extends ReqPage {
  toUid?: string | string[]
}

export interface ReqAuthorFollow extends ReqPage {
  formUid?: string | string[]
}

export interface ResFollow extends BaseResponseType {
  ID?: number
  formUid?: string
  toUid?: string
  formUidRelation?: ResAuthor
  toUidRelation?: ResAuthor
}
