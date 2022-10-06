/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-30 16:26:56
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 17:41:07
 * @FilePath: /pinkmoe_vitesse/src/api/chat/types.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
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
