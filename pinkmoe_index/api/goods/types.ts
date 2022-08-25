/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 13:57:54
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 17:41:23
 * @FilePath: /pinkmoe_vitesse/src/api/goods/types.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import type { ResAuthor, ResPost } from '/@/api/home/types'
import type { BaseResponseType } from '/@/api/common/types'
import type { ResGoods } from '../shop/types'

export interface ReqGoodsItem {
  goodsId?: string | string[]
}

export interface ResGoodsItem {
  authorPosts?: Array<ResPost>
  comments?: Array<ResComment>
  goods?: ResGoods
  users?: Array<ResAuthor>
  commentCount?: number
  fansCount?: number
  followCount?: number
  followStatus?: boolean
  postCount?: number
}

export interface ResComment extends BaseResponseType {
  ID?: number
  children?: Array<ResComment>
  content?: string
  formUid?: string
  parentId?: number
  postId?: string
  status?: string
  toUid?: string
  type?: string
  postRelation?: ResPost
  ToUidRelation?: ResAuthor
  FormUidRelation?: ResAuthor
}
