/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-20 20:59:50
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-21 14:55:29
 * @FilePath: /pinkmoe_index/src/api/home/types.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
import type { BaseResponseType } from '/@/api/common/types'
import type { ResCategory } from '/@/api/category/types'
import type { ResFile } from '../upload/types'
import type { ReqPostDownloadPost, ReqPostMusicPost, ReqPostVideoPost } from '../post/types'

export interface ResHomeList {
  content?: Array<ResContent>
  popular?: Array<ResPost>
  recommend?: Array<ResPost>
}

export interface ResContent {
  icon?: string
  name?: string
  slug?: string
  sort?: number
  style?: number
  topic?: Array<any>
  post?: Array<ResPost>
}

export interface ResPost extends BaseResponseType {
  author?: string
  category?: string
  commentStatus?: string
  content?: string
  cover?: string
  exerpt?: string
  from?: string
  postId?: string
  postImg?: string
  status?: string
  title?: string
  type?: string
  view?: number
  topic?: Array<any>
  AuthorRelation?: ResAuthor
  CategoryRelation?: ResCategory
  collectRelation?: Array<any>
  starRelation?: Array<any>
  musicRelation?: Array<ReqPostMusicPost>
  videoRelation?: Array<ReqPostVideoPost>
  downloadRelation?: Array<ReqPostDownloadPost>
  fileRelation?: Array<ResFile>
}

export interface ResAuthor extends BaseResponseType {
  authorities?: ResAuthority
  authority?: string
  authorityId?: string
  avatar?: string
  cash?: number
  credit?: number
  desc?: string
  email?: string
  exp?: number
  headerImg?: string
  nickName?: string
  phone?: string
  sex?: string
  userName?: string
  uuid?: string
}

export interface ResAuthority extends BaseResponseType {
  authorityId?: string
  authorityName?: string
  dataAuthorityId?: string
  menus?: string
}
