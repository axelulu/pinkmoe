/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-20 20:59:50
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 17:41:51
 * @FilePath: /pinkmoe_vitesse/src/api/shop/types.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import type { BaseResponseType } from '/@/api/common/types'
import type { ResCategory } from '/@/api/category/types'
import type { ResFile } from '../upload/types'

export interface ResShopList {
  content?: Array<ResContent>
  popular?: Array<ResGoods>
  shopCategory?: Array<ResCategory>
}

export interface ResContent {
  icon?: string
  iconType?: string
  name?: string
  slug?: string
  sort?: number
  style?: number
  goods?: Array<ResGoods>
}

export interface ResGoods extends BaseResponseType {
  author?: string
  category?: string
  content?: string
  cover?: string
  from?: string
  goodsId?: string
  goodsImg?: string
  status?: string
  title?: string
  type?: string
  view?: number
  AuthorRelation?: ResAuthor
  CategoryRelation?: ResCategory
  sizeValRelation?: Array<any>
  sizeRelation?: Array<any>
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
