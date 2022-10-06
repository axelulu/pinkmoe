/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 13:57:54
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 17:41:23
 * @FilePath: /pinkmoe_vitesse/src/api/goods/types.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
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
