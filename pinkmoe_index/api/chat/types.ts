/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-30 16:26:56
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 17:41:07
 * @FilePath: /pinkmoe_vitesse/src/api/chat/types.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import type { ResAuthor } from '../home/types'
import type { BaseResponseType, ReqPage } from '/@/api/common/types'

export interface ReqChat extends ReqPage {
  sendId?: string
}

export interface ResChat extends BaseResponseType {
  ID?: number
  amount?: number
  cmd?: number
  content?: string
  media?: number
  memo?: string
  pic?: string
  sendId?: string
  url?: string
  userId?: string
  sendIdRelation?: ResAuthor
  userIdRelation?: ResAuthor
}

export interface ResChatRelation extends BaseResponseType {
  ID?: number
  amount?: number
  cmd?: number
  content?: string
  media?: number
  memo?: string
  pic?: string
  sendId?: string
  url?: string
  userId?: string
  sendIdRelation?: ResAuthor
  userIdRelation?: ResAuthor
}
